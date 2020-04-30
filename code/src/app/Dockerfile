FROM ubuntu:16.04

RUN mkdir -p /work/project

WORKDIR /work/project
ADD ./app.tar.gz /work/project

RUN apt-get update 
RUN apt-get install -y curl vim cron wget
RUN apt-get install -y net-tools

RUN apt-get install -y supervisor

ENTRYPOINT  ./start.sh

EXPOSE 8080
