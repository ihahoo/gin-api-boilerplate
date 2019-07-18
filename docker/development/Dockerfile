FROM golang:1.12-alpine

ENV appdir /go/src/gin-api-boilerplate
RUN mkdir -p $appdir
WORKDIR $appdir

RUN echo https://mirrors.tuna.tsinghua.edu.cn/alpine/v3.10/main > /etc/apk/repositories \
    && apk add --no-cache git make \
    && go get -v github.com/pilu/fresh
