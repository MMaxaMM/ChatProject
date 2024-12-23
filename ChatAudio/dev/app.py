from concurrent import futures
import grpc
from audio_pb2 import (
    AudioRequest,
    AudioResponse
)
import audio_pb2_grpc as audio_grpc

import time

MAX_WORKERS = 5
ADDRESS = "[::]:50200"


class AudioService(audio_grpc.AudioServiceServicer):
    def Recognize(self, request:AudioRequest, context:grpc.ServicerContext) -> AudioResponse:
        # Запрос пользователя
        print("Пришел запрос от пользователя")
        print(request.URI)

        # Заглушка для модели
        time.sleep(5)

        # Ответ сервиса
        context.set_code(grpc.StatusCode.OK)
        content = '''**SPEAKER_00**: Привет, как дела?
        **SPEAKER_01**: Привет, хорошо!'''
        return AudioResponse(content=content)
    
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_WORKERS))
    audio_grpc.add_AudioServiceServicer_to_server(AudioService(), server)
    server.add_insecure_port(ADDRESS)
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    print("### Starting Audio service ###")
    print(f"Address: {ADDRESS}")
    serve()