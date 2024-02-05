FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.* ./

RUN go mod download

ENTRYPOINT go run ./...