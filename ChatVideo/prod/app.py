from concurrent import futures
import grpc
from video_pb2 import (
    VideoRequest,
    VideoResponse
)
import video_pb2_grpc as video_grpc

from minio import Minio

import cv2
from ultralytics import YOLO
from facenet_pytorch import MTCNN, InceptionResnetV1
import numpy as np
import torch
import faiss

from collections import defaultdict
from tqdm import tqdm
import pickle
import re
import os
import uuid
import yaml

import warnings
warnings.filterwarnings("ignore")

os.environ['KMP_DUPLICATE_LIB_OK'] = 'True'
os.environ["TORCH_HOME"] = '/video/video/models/'

DEFAULT_CONFIG = {
    "MAX_WORKERS": 5,
    "ADDRESS": "0.0.0.0:50300",
    "YOLO_PATH": "/video/video/models/yolov11m-face.pt",
    "TRACKER_PATH": "/video/video/models/tracker.yaml",
    "MINIO_ADDRESS": "minio:9000",
    "BUCKET_NAME": "video",
    "PHOTO_PATH": "/video/video/photo",
    "INDEX_PATH": "/video/video/index/faiss.index",
    "NAMES_PATH": "/video/video/index/names.pkl",
    "THRESHOLD": 0.65,
    "STEP": 5,
}

BLUE = (255, 0, 0)
RED = (0, 0, 255)
GREEN = (0, 255, 0)

class VectorStore():
    def __init__(self, mtcnn, resnet, device, threshold=0.7, d=512):        
        self.index = None
        self.names = None
        self.mtcnn = mtcnn
        self.resnet = resnet
        self.device = device
        
        self.threshold = threshold
        self.d = d
    
    def load(self, index_path, names_path):
        self.index = faiss.read_index(index_path)
        with open(names_path, "rb") as file:
                self.names = pickle.load(file)
        
    def new(self, photo_path, index_path, names_path):
        self.index = faiss.IndexFlatIP(self.d)
        self.names = dict()
        
        for idx, file in enumerate(os.listdir(photo_path)):
            filename = os.path.basename(file)
            matches = re.findall(r"\$([\w\d_\-]+)\$", filename, re.IGNORECASE)      
            if len(matches) == 0:
                continue
            name = matches[0]
            
            image = cv2.imread(os.path.join(photo_path, file))
            with torch.no_grad():
                image_cropped = self.mtcnn(image).to(self.device)
                embedding = self.resnet(image_cropped.unsqueeze(0)).cpu()
            
            self.index.add(embedding)
            self.names[idx] = name
            print(f"Indexed {idx}: {name}")
            
        self.save(index_path, names_path)
            
    def save(self, index_path, names_path):
        faiss.write_index(self.index, index_path)
        with open(names_path, "wb") as file:
            pickle.dump(self.names, file)
         
    def search(self, embedding):
        D, I = self.index.search(embedding.cpu(), 1)
        D = D[D>self.threshold]
        I = I[D>self.threshold]
        
        if D.size > 0:
            score = np.max(D)
            idx = I[np.argmax(D)][0]     
            return (score, self.names[idx])
        else:
            return (0, "")


class VideoService(video_grpc.VideoServiceServicer):
    def __init__(self, config):
        self.config = config
        self.device = 'cuda' if torch.cuda.is_available() else 'cpu'
        self.yolo = YOLO(self.config["YOLO_PATH"]).to(self.device)
        self.mtcnn = MTCNN(image_size=160, margin=40, device=self.device).eval()
        self.resnet = InceptionResnetV1(pretrained='vggface2', device=self.device).eval()
        self.store = VectorStore(self.mtcnn, self.resnet, self.device, self.config["THRESHOLD"])
        self.store.new(self.config["PHOTO_PATH"], self.config["INDEX_PATH"], self.config["NAMES_PATH"])
        self.client = Minio(self.config["MINIO_ADDRESS"], os.environ['MINIO_USER'], os.environ['MINIO_PASSWORD'], secure=False)


    def Detect(self, request:VideoRequest, context:grpc.ServicerContext) -> VideoResponse:
        print("### A request came from the user ###")

        TRACKS = defaultdict(list)
        FACES = defaultdict(lambda: (0, ""))

        video = cv2.VideoCapture(request.URI)
        total_frames = int(video.get(cv2.CAP_PROP_FRAME_COUNT))

        for frame_num in tqdm(range(total_frames)):
            ret, frame = video.read()
            if not ret: break
            
            frame = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)   
            
            with torch.no_grad(): 
                result = self.yolo.track(source=frame, persist=True, tracker=self.config["TRACKER_PATH"], verbose=False)[0]
            
            boxes = result.boxes.xyxy.cpu()
            if result.boxes.id is None:
                continue 
            tarck_ids = result.boxes.id.int().cpu().tolist()
            
            for box, tarck_id in zip(boxes, tarck_ids): 
                TRACKS[tarck_id].append((frame_num, box))
                        
                if (len(TRACKS[tarck_id])) % self.config["STEP"] == 1:
                    with torch.no_grad():
                        face = self.mtcnn.extract(frame, box.unsqueeze(0), None).to(self.device)               
                        search_info = self.store.search(self.resnet(face.unsqueeze(0)).cpu())
                        if FACES[tarck_id][0] < search_info[0]:
                            FACES[tarck_id] = search_info


        frame_num2box = defaultdict(list)

        find_idx = [idx for idx in FACES if FACES[idx][0] > 0]
        for idx in find_idx:
            for frame_num, box in TRACKS[idx]:
                frame_num2box[frame_num].append((FACES[idx][1], box))

        video = cv2.VideoCapture(request.URI)

        filename = str(uuid.uuid4()) + ".mp4"
        fourcc = cv2.VideoWriter_fourcc(*'avc1')
        frame_width = int(video.get(cv2.CAP_PROP_FRAME_WIDTH))
        frame_height = int(video.get(cv2.CAP_PROP_FRAME_HEIGHT))
        out = cv2.VideoWriter(filename, fourcc, 30, (frame_width, frame_height))

        for frame_num in tqdm(range(total_frames)):
            ret, frame = video.read()
            if not ret: break
                
            for name, box in frame_num2box[frame_num]:
                x1, y1, x2, y2 = box.int().tolist()
                cv2.rectangle(frame, (x1, y1), (x2, y2), RED, 3)
                cv2.putText(frame, name, (x1, y1-10), cv2.FONT_HERSHEY_SIMPLEX, 1, RED, 3)
                    
            out.write(frame)
            
        video.release()
        out.release()
        cv2.destroyAllWindows()

        self.client.fput_object(self.config["BUCKET_NAME"], filename, filename, "video/mp4")
        os.remove(filename)

        print("### Done ###")

        context.set_code(grpc.StatusCode.OK)
        return VideoResponse(objectName=filename)


def Serve(config):
    print("### Starting Video service ###")
    print(f'### Address: {config["ADDRESS"]} ###')

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=config["MAX_WORKERS"]))

    service = VideoService(config)
    print(f"### Device: {service.device} ###")

    video_grpc.add_VideoServiceServicer_to_server(service, server)
    server.add_insecure_port(config["ADDRESS"])

    print("### Video service is running ###")

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
    config["YOLO_PATH"] = config.get('yolo_path', DEFAULT_CONFIG["YOLO_PATH"])
    config["TRACKER_PATH"] = config.get('tracker_path', DEFAULT_CONFIG["TRACKER_PATH"])
    config["MINIO_ADDRESS"] = config.get('minio_address', DEFAULT_CONFIG["MINIO_ADDRESS"])
    config["BUCKET_NAME"] = config.get('bucket_name', DEFAULT_CONFIG["BUCKET_NAME"])
    config["PHOTO_PATH"] = config.get('photo_path', DEFAULT_CONFIG["PHOTO_PATH"])
    config["INDEX_PATH"] = config.get('index_path', DEFAULT_CONFIG["INDEX_PATH"])
    config["NAMES_PATH"] = config.get('names_path', DEFAULT_CONFIG["NAMES_PATH"])
    config["THRESHOLD"] = config.get('threshold', DEFAULT_CONFIG["THRESHOLD"])
    config["STEP"] = config.get('step', DEFAULT_CONFIG["STEP"])

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