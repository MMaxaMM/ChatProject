from concurrent import futures
import grpc
from rag_pb2 import (
    RAGRequest,
    RAGResponse
)
import rag_pb2_grpc as rag_grpc

import time

MAX_WORKERS = 5
ADDRESS = "0.0.0.0:50400"


class RAGService(rag_grpc.RAGServiceServicer):
    def Generate(self, request: RAGRequest, context: grpc.ServicerContext) -> RAGResponse:
        # Просмотр пришедшего запроса
        print(f"Запрос пользователя:")
        print(f"Content: {request.content}")

        # Заглушка для модели
        time.sleep(5)

        # Ответ сервиса
        context.set_code(grpc.StatusCode.OK)
        return RAGResponse(content="Найдены следующие документы...")
    
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_WORKERS))
    rag_grpc.add_RAGServiceServicer_to_server(RAGService(), server)
    server.add_insecure_port(ADDRESS)
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    print("### Starting RAG service ###")
    print(f"Address: {ADDRESS}")
    serve()