FROM ubuntu:16.04

RUN mkdir -p /work

WORKDIR /work

RUN apt-get update && apt-get install -y curl vim cron wget
RUN apt-get install -y net-tools

RUN sudo apt-get install supervisor

ENTRYPOINT start.sh

EXPOSE 8080
