FROM golang:1.13.6-alpine

ENV CGO_ENABLED=0

WORKDIR /go/src

COPY ./src /go/src

RUN apk update && apk add git
#RUN go get -u github.com/labstack/echo/...