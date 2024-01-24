FROM golang:1.21.6-alpine3.19

ENV CGO_ENABLED=0

RUN apk update &&\
    apk upgrade &&\
    apk add --update figlet &&\
    go install github.com/cespare/reflex@latest
