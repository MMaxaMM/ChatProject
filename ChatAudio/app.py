from concurrent import futures
import grpc
from audio_pb2 import (
    Error,
    AudioRequest,
    AudioResponse
)
import audio_pb2_grpc as audio_grpc

MAX_WORKERS = 5
ADDRESS = "[::]:50100"


class AudioService(audio_grpc.AudioServiceServicer):
    def Recognize(self, request, context):
        raise NotImplementedError()
    
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_WORKERS))
    audio_grpc.add_AudioServiceServicer_to_server(AudioService(), server)
    server.add_insecure_port(ADDRESS)
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()