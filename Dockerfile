FROM golang:1.12

WORKDIR  .
COPY . .

RUN mkdir gopath
ENV GOPATH /go/gopath
#

ENV GOOS=linux GOARCH=amd64 GO111MODULE=on

RUN go build main.go
CMD ["./main"]
