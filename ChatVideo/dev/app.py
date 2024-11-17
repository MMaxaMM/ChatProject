from concurrent import futures
import grpc
from video_pb2 import (
    VideoRequest,
    VideoResponse
)
import video_pb2_grpc as video_grpc

import time

MAX_WORKERS = 5
ADDRESS = "[::]:50300"


class VideoService(video_grpc.VideoServiceServicer):
    def Detect(self, request:VideoRequest, context:grpc.ServicerContext) -> VideoResponse:
        # Запрос пользователя
        print("Пришел запрос от пользователя")
        print(request.filepath)

        # Заглушка для модели
        time.sleep(5)

        # Ответ сервиса
        context.set_code(grpc.StatusCode.OK)
        filepath = "video/not-exists.mp4"
        return VideoResponse(filepath=filepath)
    
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_WORKERS))
    video_grpc.add_VideoServiceServicer_to_server(VideoService(), server)
    server.add_insecure_port(ADDRESS)
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    print("### Starting Video service ###")
    print(f"Address: {ADDRESS}")
    serve()