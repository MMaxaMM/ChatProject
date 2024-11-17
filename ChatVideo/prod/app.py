from concurrent import futures
import grpc

from video_pb2 import (
    VideoRequest,
    VideoResponse
)

import video_pb2_grpc as video_grpc
import torch
from ultralytics import YOLO
import uuid

from minio import Minio

import cv2
import os

import yaml

import warnings
warnings.filterwarnings("ignore")

MAX_WORKERS = 5 # default
ADDRESS = "[::]:50300" # default
YOLO_PATH = "models/YOLOv8/yolov8l.pt" # default

FILESTORAGE = "minio:9000" # default
USER = "minio" # default
PASSWORD = "11111111" # default
BUCKET_NAME = "video" # default


class VideoService(video_grpc.VideoServiceServicer):
    def __init__(self):
        self.device = "cuda" if torch.cuda.is_available() else "cpu"
        self.yolo = InitYOLO(self.device)
        self.client = Minio(FILESTORAGE, USER, PASSWORD, secure=False)


    def Detect(self, request:VideoRequest, context:grpc.ServicerContext) -> VideoResponse:
        print("### A request came from the user ###")

        cap = cv2.VideoCapture(f"http://{FILESTORAGE}/{request.filepath}")

        filename = str(uuid.uuid4()) + ".mp4"
        fourcc = cv2.VideoWriter_fourcc(*'mp4v')
        frame_width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
        frame_height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
        out = cv2.VideoWriter(filename, fourcc, 30, (frame_width, frame_height))
        while cap.isOpened():
            ret, frame = cap.read()
            if not ret: break
                
            # Выполняем обнаружение объектов
            results = self.yolo(frame)

            # Извлекаем аннотации
            boxes = results[0].boxes

            # Отображаем кадр с аннотациями
            for box in boxes:
                box = box.cpu()  # Перемещение тензора на CPU
                x1, y1, x2, y2 = box.xyxy[0]  # получение координат
                conf = box.conf[0]  # или box[4] в зависимости от структуры box
                cls = box.cls[0]
                cv2.rectangle(frame, (int(x1), int(y1)), (int(x2), int(y2)), (255, 0, 0), 2)
                cv2.putText(frame, f'{self.yolo.names[int(cls)]} {conf:.2f}', (int(x1), int(y1)-10), 
                            cv2.FONT_HERSHEY_SIMPLEX, 0.5, (255, 0, 0), 2)

            # Сохраняем кадр в выходное видео
            out.write(frame)

        cap.release()
        out.release()
        cv2.destroyAllWindows()

        self.client.fput_object(BUCKET_NAME, filename, filename, "video/mp4")
        os.remove(filename)

        print("### Done ###")

        context.set_code(grpc.StatusCode.OK)
        return VideoResponse(filepath=f"{BUCKET_NAME}/{filename}")


def InitYOLO(device):
    print("### YOLO initialization ###")

    model = YOLO(YOLO_PATH).to(device)
    return model 


def serve():
    print("### Starting Video service ###")
    print(f"### Address: {ADDRESS} ###")

    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=MAX_WORKERS),
        options=[("grpc.max_message_length", 200 * 1024 * 1024)])

    service = VideoService()
    print(f"### Device: {service.device} ###")

    video_grpc.add_VideoServiceServicer_to_server(service, server)
    server.add_insecure_port(ADDRESS)

    print("### Video service is running ###")

    server.start()
    server.wait_for_termination()


def LoadConfig(config_path):
    try:
        with open(config_path) as file:
            config = yaml.load(file, Loader=yaml.FullLoader)
    except:
        print(f"### Failed to load configuration file: {config_path} ###")
        return
    
    MAX_WORKERS = config.get('max_workers', MAX_WORKERS)
    ADDRESS = config.get('address', ADDRESS)
    YOLO_PATH = config.get('yolo_path', YOLO_PATH)

    MINIO_ADDRESS = config.get('minio_address', MINIO_ADDRESS)
    MINIO_USER = config.get('minio_user', MINIO_USER)
    MINIO_PASSWORD = config.get('minio_password', MINIO_PASSWORD)
    BUCKET_NAME = config.get('bucket_name', BUCKET_NAME)


if __name__ == "__main__":
    config_path = os.getenv("CONFIG_PATH")
    if config_path is None:
        print("### The environment variable CONFIG_PATH is not set ###")
        print("### Load default configurations ###")

    LoadConfig(config_path)
    serve()