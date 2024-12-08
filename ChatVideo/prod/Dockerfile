FROM nvidia/cuda:11.4.3-cudnn8-runtime-ubuntu20.04

RUN apt update
RUN apt -y install unzip
RUN apt -y install python3
RUN apt -y install python3-pip
RUN apt -y install wget
RUN apt -y install python-concurrent.futures
RUN apt -y install nano
RUN export DEBIAN_FRONTEND=noninteractive && apt -y install ffmpeg

WORKDIR /video
COPY . .

RUN wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh
RUN mkdir /root/.conda
RUN bash Miniconda3-latest-Linux-x86_64.sh -b
RUN rm -f Miniconda3-latest-Linux-x86_64.sh
ENV PATH="/root/miniconda3/bin:/root/miniconda3/envs:{$PATH}"

RUN conda config --append channels conda-forge
RUN conda env create -f environment.yaml
SHELL ["conda", "run", "--name", "video", "/bin/bash", "-c"]

RUN pip3 install --no-cache-dir -r requirements.txt

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8 

CMD ["bash", "-c", "source activate video && python3 app.py"]