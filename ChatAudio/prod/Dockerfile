FROM nvidia/cuda:11.4.3-cudnn8-runtime-ubuntu20.04

RUN apt-get update
RUN apt-get install unzip
RUN apt-get -y install python3
RUN apt-get -y install python3-pip
RUN apt-get -y install python-concurrent.futures
RUN apt-get -y install nano
RUN export DEBIAN_FRONTEND=noninteractive && apt-get -y install ffmpeg

WORKDIR /audio
COPY . .
RUN pip3 install --no-cache-dir -r requirements.txt

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8 

CMD ["python3", "app.py"]