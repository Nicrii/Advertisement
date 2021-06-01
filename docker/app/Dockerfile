FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY ./../../src/go.mod ./
COPY ./../../src/go.sum ./

RUN go mod download

COPY ./../../src/ .

RUN go build -o ./main

CMD ["./main"]