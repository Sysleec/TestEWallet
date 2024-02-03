FROM golang:1.21-alpine

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run ./...