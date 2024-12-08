from concurrent import futures
import grpc
from rag_pb2 import (
    RAGRequest,
    RAGResponse
)
import rag_pb2_grpc as rag_grpc

from pymilvus import MilvusClient
import sqlite3

import torch
from sentence_transformers import SentenceTransformer

from llama_index.retrievers.bm25 import BM25Retriever
from llama_index.core.node_parser import SentenceSplitter
from llama_index.core.storage.docstore import SimpleDocumentStore
from llama_index.core import Document
import Stemmer

from transformers import pipeline

import os
import re
from tqdm import tqdm
import yaml

import warnings
warnings.filterwarnings("ignore")

DEFAULT_CONFIG = {
    "MAX_WORKERS": 5,
    "ADDRESS": "0.0.0.0:50400",
    "VECTORDB_PATH": "rag/index/milvus.db",
    "COLLECTION_NAME": "documents",
    "DOCUMENTDB_PATH": "rag/index/sqlite.db",
    "BM25_PATH": "rag/index/dokstore.bm25",
    "DOCUMENTS_PATH": "rag/documents",
    "UPDATE": False,
    "LLM_MODEL": "rag/models/Qwen2.5-7B-Instruct",
    "EMBEDDIND_MODEL": "rag/models/multilingual-e5-large-instruct",
    "RERANKER_MODEL": "rag/models/bge-reranker-v2-m3",
    "CHUNK_SIZE": 300,
    "CHUNK_OVERLAP": 25,
    "MAX_NEW_TOKENS": 4096,
}

PROMPTS = {
    "EMBEDDIND_MODEL": 'Instruct: Given a web search query, retrieve relevant passages that answer the query\nQuery: ',
    "RAG_1": '''You are an intelligent assistant in the RAG system. The user asks you his question, and you have to find relevant documents in the knowledge base in order to build your answer based on them.
First, you need to make a query to the knowledge base. Your request should not be a question, but affirmative suggestions, for example: instead of "What scholarships are available at Skoltech?", the query should be "Scholarships at Skoltech". The response should consist only of a query to the knowledge base.
The knowledge base consists of documents in Russian. The request to the knowledge base must be in Russian.
''',
    "RAG_2": '''You are an intelligent assistant in the RAG system. The user asks you his question, and you have to find relevant documents in the knowledge base in order to build your answer based on them.
You have already received the documents from the knowledge base, they are listed below. For each document, the degree of similarity of the document to the knowledge base query is also indicated.
Based on these documents, give a detailed answer to the user's question. Answer only the question you have asked, the answer should be brief and relevant to the question. Do not distort the information contained in the knowledge base!
If you are sure that the documents found do not correspond to the user's question, do not try to answer the question yourself!!! In this case, tell the user that the information they are interested in is not contained in the knowledge base.
The knowledge base consists of documents in Russian. The answer must be in Russian.

Documents found:
{}
'''
}

class DocumentDB:
    def __init__(self, uri):
        self.documentdb = sqlite3.connect(uri)
        self.cursor = self.documentdb.cursor()

    def init(self):
        self.cursor.execute('''
            CREATE TABLE IF NOT EXISTS Documents (
            id INTEGER PRIMARY KEY,
            document TEXT NOT NULL)
            '''
            )
        self.cursor.execute("CREATE INDEX IF NOT EXISTS id_index ON Documents(id)")
        self.documentdb.commit()

        return self

    def insert(self, document) -> int:
        result = self.cursor.execute("INSERT INTO Documents (document) VALUES (?) RETURNING id", (document, ))
        id = result.fetchone() or (0, )
        self.documentdb.commit()

        return id[0]

    def select(self, id) -> str:
        result = self.cursor.execute("SELECT document FROM Documents WHERE id = ?", (id, ))
        document = result.fetchone() or ("", )

        return document[0]

class VectorDB:
    def __init__(self, uri, collection_name, embedding_model, limit=5):
        self.vectordb = MilvusClient(uri)
        self.collection_name = collection_name
        self.embedding_model = embedding_model
        self.limit = limit

    def init(self):
        if not self.vectordb.has_collection(self.collection_name):
            with torch.no_grad():
                dimension = len(self.embedding_model.encode("Hello world"))
            self.vectordb.create_collection(
                collection_name=self.collection_name,
                dimension=dimension,
                auto_id = True,
                metric_type="IP",
                consistency_level="Strong",
            )

        return self

    def insert(self, text, document_id):
        with torch.no_grad():
            embedding = self.embedding_model.encode(PROMPTS["EMBEDDIND_MODEL"] + text, normalize_embeddings=True)
        data = [{"vector": embedding, "text": text, "document_id": document_id}]
        self.vectordb.insert(collection_name=self.collection_name, data=data)

    def search(self, query) -> int:
        with torch.no_grad():
            embedding = self.embedding_model.encode(PROMPTS["EMBEDDIND_MODEL"] + query, normalize_embeddings=True)

        result = self.vectordb.search(
            collection_name = self.collection_name,
            data=[embedding],
            limit=self.limit,
            search_params={"metric_type": "IP", "params": {}},
            output_fields=["text", "document_id"],
        )
        retrieved_documents = [(res["entity"]["text"], res["entity"]["document_id"]) for res in result[0]]

        return retrieved_documents
    
class BM25:
    def __init__(self, docstore_path, similarity_top_k=5):
        self.docstore_path = docstore_path
        self.docstore = SimpleDocumentStore()
        self.bm25 = None
        self.similarity_top_k = similarity_top_k

    def init(self):
        try:
            self.docstore.from_persist_path(self.docstore_path)
            self.bm25 = BM25Retriever.from_defaults(
                docstore = self.docstore,
                similarity_top_k = self.similarity_top_k,
                stemmer = Stemmer.Stemmer("russian"),
                language = "russian"
            )
        except FileNotFoundError:
            self.docstore.persist(self.docstore_path)

        return self

    def update_index(self):
        self.docstore.persist(self.docstore_path)
        self.bm25 = BM25Retriever.from_defaults(
            docstore = self.docstore,
            similarity_top_k = self.similarity_top_k,
            stemmer = Stemmer.Stemmer("russian"),
            language = "russian"
        )

    def insert(self, text, document_id):
        document = Document(text=text, metadata={"document_id": document_id})
        self.docstore.add_documents([document])

    def search(self, query):
        result = self.bm25.retrieve(query)
        retrieved_documents = [(res.text, res.metadata["document_id"]) for res in result]

        return retrieved_documents
    
class Parser:
    @staticmethod
    def parse_txt(filepath) -> str:
        with open(filepath, "r", encoding="utf-8") as file:
            document = re.sub(r'[\n\t\s]+', ' ', file.read())
        return document
    
class RetrieverTool:
    def __init__(self, documentdb, vectordb, bm25, parser, splitter, reranker, k=3):
        self.documentdb = documentdb
        self.vectordb = vectordb
        self.bm25 = bm25
        self.parser = parser
        self.splitter = splitter
        self.reranker = reranker
        self.k = k

    def update(self, documents_path):
        for filename in tqdm(os.listdir(documents_path)):
            filepath = os.path.join(documents_path, filename)
            document = self.parser.parse_txt(filepath)

            document_id = self.documentdb.insert(document)

            for chunk in self.splitter.split_text(document):
                self.vectordb.insert(chunk, document_id)
                self.bm25.insert(chunk, document_id)

        self.bm25.update_index()

    def forward(self, query) -> str:
        search_result = self.vectordb.search(query) + self.bm25.search(query)

        document_ids = [result[1] for result in search_result]
        documents = [self.documentdb.select(id) for id in set(document_ids)] #На уровне чанков?

        with torch.no_grad():
            query_embedding = self.reranker.encode(query)
            ranked_search_result = [
                (self.reranker.similarity(query_embedding, self.reranker.encode(document)).item(), document)
                for document in documents
            ]
        ranked_search_result = sorted(ranked_search_result, key=lambda x: x[0], reverse=True)[:self.k]

        return "Retrieved documents:\n" + "".join(
            [
                f"===== Document {i}, similarity to query: {score_and_doc[0]} =====\n{score_and_doc[1]}"
                for i, score_and_doc in enumerate(ranked_search_result)
            ]
        )
    
class HFLLM:
    def __init__(self, model_name, device, max_new_tokens=4096):
        self.model_name = model_name
        self.device = device
        self.max_new_tokens = max_new_tokens
        self.pipeline = pipeline("text-generation", model=self.model_name, device=self.device)

    def generate(self, messages):
        with torch.no_grad():
            result = self.pipeline(messages, max_new_tokens=self.max_new_tokens)

        return result[0]["generated_text"]
    
class RAG:
    def __init__(self, llm_model, retriever_tool, verbose=True):
        self.llm_model = llm_model
        self.retriever_tool = retriever_tool
        self.verbose = verbose

    def question(self, question):
        messages = [{"role": "system", "content": PROMPTS["RAG_1"]},
                    {"role": "user", "content": question}]
        retrieve_query = self.llm_model.generate(messages)

        retrieved_documents = self.retriever_tool.forward(retrieve_query)

        if self.verbose:
            print(retrieved_documents)

        messages = [{"role": "system", "content": PROMPTS["RAG_2"].format(retrieved_documents)},
                    {"role": "user", "content": question}]

        result = self.llm_model.generate(messages)

        return result

class RAGService(rag_grpc.RAGServiceServicer):
    def __init__(self, config):
        self.config = config
        self.device = "cuda" if torch.cuda.is_available() else "cpu"

        self.embedding_model = SentenceTransformer(self.config["EMBEDDIND_MODEL"]).to(self.device)
        self.reranker_model = SentenceTransformer(self.config["RERANKER_MODEL"]).to(self.device)

        self.documentdb = DocumentDB(self.config["DOCUMENTDB_PATH"]).init()
        self.vectordb = VectorDB(self.config["VECTORDB_PATH"], self.config["COLLECTION_NAME"], self.embedding_model).init()
        self.bm25 = BM25(self.config["BM25_PATH"]).init()

        self.parser = Parser()
        self.splitter = SentenceSplitter(chunk_size=self.config["CHUNK_SIZE"], chunk_overlap=self.config["CHUNK_OVERLAP"])

        self.retriever_tool = RetrieverTool(self.documentdb, self.vectordb, self.bm25, self.parser, self.splitter, self.reranker_model)
        if self.config["UPDATE"]:
            self.retriever_tool.update(self.config["DOCUMENTS_PATH"])

        self.llm_model = HFLLM(self.config["LLM_MODEL"], self.device, self.config["MAX_NEW_TOKENS"])
        self.rag = RAG(self.llm_model, self.retriever_tool, verbose=False)

    def Generate(self, request:RAGRequest, context:grpc.ServicerContext) -> RAGResponse:
        print("### A request came from the user ###")
        
        question = request.content
        answer = self.rag.question(question)

        context.set_code(grpc.StatusCode.OK)
        return RAGResponse(content=answer)

def Serve(config):
    print("### Starting RAG service ###")
    print(f'### Address: {config["ADDRESS"]} ###')

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=config["MAX_WORKERS"]))

    service = RAGService(config)
    print(f"### Device: {service.device} ###")

    rag_grpc.add_RAGServiceServicer_to_server(service, server)
    server.add_insecure_port(config["ADDRESS"])

    print("### RAG service is running ###")

    server.start()
    server.wait_for_termination()

def LoadConfig(config_path):
    try:
        with open(config_path) as file:
            config = yaml.load(file, Loader=yaml.FullLoader)
    except:
        print(f"### Failed to load configuration file: {config_path} ###")
        print("### Load default configurations ###")
        return DEFAULT_CONFIG
    
    config = dict()
    config["MAX_WORKERS"] = config.get('max_workers', DEFAULT_CONFIG["MAX_WORKERS"])
    config["ADDRESS"] = config.get('address', DEFAULT_CONFIG["ADDRESS"])
    config["VECTORDB_PATH"] = config.get('vectordb_path', DEFAULT_CONFIG["VECTORDB_PATH"])
    config["COLLECTION_NAME"] = config.get('collection_name', DEFAULT_CONFIG["COLLECTION_NAME"])
    config["DOCUMENTDB_PATH"] = config.get('documentdb_path', DEFAULT_CONFIG["DOCUMENTDB_PATH"])
    config["BM25_PATH"] = config.get('bm25_path', DEFAULT_CONFIG["BM25_PATH"])
    config["DOCUMENTS_PATH"] = config.get('documents_path', DEFAULT_CONFIG["DOCUMENTS_PATH"])
    config["UPDATE"] = config.get('update', DEFAULT_CONFIG["UPDATE"])
    config["LLM_MODEL"] = config.get('llm_model', DEFAULT_CONFIG["LLM_MODEL"])
    config["EMBEDDIND_MODEL"] = config.get('embeddind_model', DEFAULT_CONFIG["EMBEDDIND_MODEL"])
    config["RERANKER_MODEL"] = config.get('reranker_model', DEFAULT_CONFIG["RERANKER_MODEL"])
    config["CHUNK_SIZE"] = config.get('chunk_size', DEFAULT_CONFIG["CHUNK_SIZE"])
    config["CHUNK_OVERLAP"] = config.get('chunk_overlap', DEFAULT_CONFIG["CHUNK_OVERLAP"])
    config["MAX_NEW_TOKENS"] = config.get('max_new_tokens', DEFAULT_CONFIG["MAX_NEW_TOKENS"])

    return config

    
if __name__ == "__main__":
    config_path = os.getenv("CONFIG_PATH")
    config = None

    if config_path is None:
        print("### The environment variable CONFIG_PATH is not set ###")
        print("### Load default configurations ###")
        config = DEFAULT_CONFIG
    else:
        config = LoadConfig(config_path)

    Serve(config)