FROM python:3.10

RUN apt-get update
RUN apt-get -y install python3-pip
RUN apt-get -y install nano

WORKDIR /video
COPY . .
RUN pip3 install --no-cache-dir -r requirements.txt

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8 

CMD ["python3", "app.py"]