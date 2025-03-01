FROM golang:1.23-bookworm

RUN apt-get update \
    && apt-get install -y \
        iputils-ping \
        iproute2 \
        dnsutils

WORKDIR /app

COPY go.mod go.sum main.go .

RUN go mod download
