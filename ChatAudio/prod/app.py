from concurrent import futures
import grpc
from audio_pb2 import (
    AudioRequest,
    AudioResponse
)
import audio_pb2_grpc as audio_grpc
import torch
from transformers import (
    AutoModelForSpeechSeq2Seq, 
    AutoProcessor, 
    pipeline,
)
from pyannote.audio import Pipeline
import io
import os
import torchaudio
from torchaudio.functional import resample
import requests

import yaml

import warnings
warnings.filterwarnings("ignore")

MAX_WORKERS = 5 # default
ADDRESS = "[::]:50200" # default
FILESTORAGE = "minio:9000" # default

WHISPER_PATH = "models/whisper-large-v3-turbo"
DIARIZATION_PATH = "models/config.yaml"

class AudioService(audio_grpc.AudioServiceServicer):
    def __init__(self):
        self.device = "cuda" if torch.cuda.is_available() else "cpu"
        self.dtype = torch.float16 if torch.cuda.is_available() else torch.float32

        self.whisper_pipeline = InitWhisper(self.device, self.dtype)
        self.diarization_pipeline = InitDiarization(self.device)

    def Recognize(self, request:AudioRequest, context:grpc.ServicerContext) -> AudioResponse:
        print("### A request came from the user ###")
        audio_io = io.BytesIO(requests.get(f"http://{FILESTORAGE}/{request.filepath}").content)

        waveform, sample_rate = torchaudio.load(audio_io)

        audio = resample(waveform, sample_rate, 22050)
        audio = audio.squeeze(0).numpy()
        sr = 22050

        diarization = self.diarization_pipeline({"waveform": waveform, "sample_rate": sample_rate})
        
        dialog = []
        for speech_turn, _, speaker in diarization.itertracks(yield_label=True):
            segment = audio[int(speech_turn.start * sr) : int(speech_turn.end * sr)]
            result = self.whisper_pipeline(segment)
            dialog.append(f"`{speaker}`: {result['text']}")

        print("### Done ###")

        result = '\n'.join(dialog)
        context.set_code(grpc.StatusCode.OK)
        return AudioResponse(result=result)    

def InitWhisper(device, dtype):
    print("### Whisper initialization ###")

    model = AutoModelForSpeechSeq2Seq.from_pretrained(
        WHISPER_PATH, 
        torch_dtype=dtype, 
        use_safetensors=True,
    ).to(device)

    processor = AutoProcessor.from_pretrained(WHISPER_PATH)

    whisper_pipeline = pipeline(
        "automatic-speech-recognition",
        model=model,
        tokenizer=processor.tokenizer,
        feature_extractor=processor.feature_extractor,
        torch_dtype=dtype,
        device=device,
    )

    return whisper_pipeline

def InitDiarization(device):
    print("### Diarization initialization ###")

    diarization_pipeline = Pipeline.from_pretrained(DIARIZATION_PATH, use_auth_token="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").to(torch.device(device))
    return diarization_pipeline

def serve():
    print("### Starting Audio service ###")
    print(f"### Address: {ADDRESS} ###")

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_WORKERS))

    service = AudioService()
    print(f"### Device: {service.device} ###")

    audio_grpc.add_AudioServiceServicer_to_server(service, server)
    server.add_insecure_port(ADDRESS)

    print("### Audio service is running ###")

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
    FILESTORAGE = config.get('filestorage', FILESTORAGE)
    
if __name__ == "__main__":
    config_path = os.getenv("CONFIG_PATH")
    if config_path is None:
        print("### The environment variable CONFIG_PATH is not set ###")
        print("### Load default configurations ###")

    LoadConfig(config_path)
    serve()