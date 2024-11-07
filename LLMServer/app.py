from concurrent import futures
import grpc
from llm_pb2 import (
    Message,
    LLMRequest,
    LLMResponse
)
import llm_pb2_grpc as llm_grpc

import time

MAX_WORKERS = 5
ADDRESS = "[::]:50100"


class LLMService(llm_grpc.LLMServiceServicer):
    def Generate(self, request: LLMRequest, context: grpc.ServicerContext) -> LLMResponse:
        # Просмотр пришедшего запроса
        messages = [{message.role: message.content} for message in request.messages]
        print(f"Запрос пользователя:")
        print(f"Messages: {messages}")
        print(f"Max tokens: {request.max_tokens}")

        # Заглушка для модели
        time.sleep(5)

        # Ответ сервиса
        context.set_code(grpc.StatusCode.OK)
        message = Message(role="assistant", content="Ответ ассистента...")
        return LLMResponse(message=message)
    
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_WORKERS))
    llm_grpc.add_LLMServiceServicer_to_server(LLMService(), server)
    server.add_insecure_port(ADDRESS)
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    print("### Starting LLM service ###")
    print(f"Address: {ADDRESS}")
    serve()