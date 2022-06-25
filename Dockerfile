#!/usr/bin/env bash

FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN cd /app

RUN go mod download

RUN go build

EXPOSE 8080

ENTRYPOINT ["/app/hello"]