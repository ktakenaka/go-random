FROM golang:1.16

RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0 \
  go get github.com/google/wire/cmd/wire \
  go get github.com/golang/mock/gomock \
  GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.3

# stop using realize temporally, it's not bettery efficient.
# RUN go get github.com/oxequa/realize
# CMD ["realize", "start"]
