FROM golang:1.23-bookworm

WORKDIR /app

COPY go.mod go.sum main.go .

RUN go mod download
