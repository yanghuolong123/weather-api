FROM ubuntu:16.04

RUN mkdir -p /work

WORKDIR /work
ADD app.tar.gz /work

RUN apt-get update && apt-get install -y curl vim cron wget
RUN apt-get install -y net-tools

RUN apt-get install -y supervisor

RUN tar -zxvf app.tar.gz

ENTRYPOINT  ./start.sh

EXPOSE 8080
