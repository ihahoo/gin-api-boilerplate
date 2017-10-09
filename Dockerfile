FROM golang:1.9-alpine 

ENV appdir /go/src/gin-api-boilerplate
ARG port=8080

RUN mkdir -p $appdir
WORKDIR $appdir

ADD Makefile .
ADD ./bin ./bin
ADD ./data/config/app.json ./data/config/app.json

RUN echo http://mirrors.aliyun.com/alpine/v3.6/main > /etc/apk/repositories \
    && echo http://mirrors.aliyun.com/alpine/v3.6/main >> /etc/apk/repositories \
    && apk add --no-cache make

EXPOSE $port
