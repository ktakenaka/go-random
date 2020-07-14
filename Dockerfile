FROM golang

RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src

RUN go get -u golang.org/x/vgo
RUN go get -u github.com/gin-gonic/gin
