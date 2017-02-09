FROM golang:1.7.5

RUN go get github.com/laher/goxc

ENV USER root
WORKDIR /go/src/github.com/yuuki/gokc

ADD . /go/src/github.com/yuuki/gokc
