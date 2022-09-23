FROM golang:1.18.6-alpine3.16

ENV CGO_ENABLED=0

RUN apk update &&\
    apk upgrade &&\
    apk add --update figlet &&\
    go install github.com/cespare/reflex@latest
