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

DEFAULT_CONFIG = {
    "MAX_WORKERS": 5,
    "ADDRESS": "0.0.0.0:50200",
    "WHISPER_PATH": "/audio/audio/models/whisper-large-v3-turbo",
    "DIARIZATION_PATH": "/audio/audio/models/config.yaml",
}

class AudioService(audio_grpc.AudioServiceServicer):
    def __init__(self, config):
        self.config = config
        self.device = "cuda" if torch.cuda.is_available() else "cpu"
        self.dtype = torch.float16 if torch.cuda.is_available() else torch.float32

        self.whisper_pipeline = InitWhisper(self.device, self.dtype, self.config)
        self.diarization_pipeline = InitDiarization(self.device, self.config)

    def Recognize(self, request:AudioRequest, context:grpc.ServicerContext) -> AudioResponse:
        print("### A request came from the user ###")
        audio_io = io.BytesIO(requests.get(request.URI).content)

        waveform, sample_rate = torchaudio.load(audio_io)

        audio = resample(waveform, sample_rate, 22050)
        audio = audio.squeeze(0).numpy()
        sr = 22050

        diarization = self.diarization_pipeline({"waveform": waveform, "sample_rate": sample_rate})
        
        dialog = []
        for speech_turn, _, speaker in diarization.itertracks(yield_label=True):
            segment = audio[int(speech_turn.start * sr) : int(speech_turn.end * sr)]
            result = self.whisper_pipeline(segment)
            dialog.append(f"**{speaker}**: {result['text']}")

        print("### Done ###")

        result = '\n'.join(dialog)
        context.set_code(grpc.StatusCode.OK)
        return AudioResponse(content=result)    

def InitWhisper(device, dtype, config):
    print("### Whisper initialization ###")

    model = AutoModelForSpeechSeq2Seq.from_pretrained(
        config["WHISPER_PATH"], 
        torch_dtype=dtype, 
        use_safetensors=True,
    ).to(device)

    processor = AutoProcessor.from_pretrained(config["WHISPER_PATH"])

    whisper_pipeline = pipeline(
        "automatic-speech-recognition",
        model=model,
        tokenizer=processor.tokenizer,
        feature_extractor=processor.feature_extractor,
        torch_dtype=dtype,
        device=device,
    )

    return whisper_pipeline

def InitDiarization(device, config):
    print("### Diarization initialization ###")

    diarization_pipeline = Pipeline.from_pretrained(config["DIARIZATION_PATH"]).to(torch.device(device))
    return diarization_pipeline

def Serve(config):
    print("### Starting Audio service ###")
    print(f'### Address: {config["ADDRESS"]} ###')

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=config["MAX_WORKERS"]))

    service = AudioService(config)
    print(f"### Device: {service.device} ###")

    audio_grpc.add_AudioServiceServicer_to_server(service, server)
    server.add_insecure_port(config["ADDRESS"])

    print("### Audio service is running ###")

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
    config["WHISPER_PATH"] = config.get('whisper_path', DEFAULT_CONFIG["WHISPER_PATH"])
    config["DIARIZATION_PATH"] = config.get('diarization_path', DEFAULT_CONFIG["DIARIZATION_PATH"])

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