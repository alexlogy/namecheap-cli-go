VERSION 0.6
FROM golang:1.18.0-buster
WORKDIR /src

build:
    COPY . /src
    RUN go build -o namecheap-go namecheap.go
    SAVE ARTIFACT /src/namecheap-go AS LOCAL ./bin/namecheap-go

all:
  BUILD +build