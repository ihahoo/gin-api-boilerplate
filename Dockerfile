FROM alpine:3.10

ENV appdir /app
ARG port=8080

RUN mkdir -p $appdir
WORKDIR $appdir

ADD Makefile .
ADD ./bin ./bin
ADD ./data/config/app.json ./data/config/app.json

RUN echo https://mirrors.tuna.tsinghua.edu.cn/alpine/v3.10/main > /etc/apk/repositories \
    && apk add --no-cache make

EXPOSE $port
