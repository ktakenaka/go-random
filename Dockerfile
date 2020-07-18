FROM golang:1.14

RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src
