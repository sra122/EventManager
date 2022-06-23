#!/usr/bin/env bash

FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /hello

FROM alpine:latest

WORKDIR /

RUN apk --no-cache add ca-certificates

COPY --from=builder /hello /hello

EXPOSE 8080

ENTRYPOINT ["/hello"]