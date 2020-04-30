FROM ubuntu:16.04

RUN mkdir -p /work/project

WORKDIR /work/project
ADD ./app.tar.gz /work/project

RUN apt-get update 
RUN apt-get install -y curl vim cron wget
RUN apt-get install -y net-tools

RUN apt-get install -y supervisor

RUN cp ./doc/app.conf /etc/supervisor/conf.d
RUN touch /var/log/supervisor/app.log

ENTRYPOINT supervisord -c /etc/supervisor/supervisord.conf && tail -f /var/log/supervisor/app.log 

EXPOSE 8080
