FROM ubuntu:16.04

RUN mkdir -p /work

WORKDIR /work
ADD app.tar.gz /work

RUN apt-get update && apt-get install -y curl vim cron wget
RUN apt-get install -y net-tools

RUN sudo apt-get install supervisor

ENTRYPOINT  bash

EXPOSE 8080
