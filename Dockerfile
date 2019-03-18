FROM golang:1.12

WORKDIR  .
COPY . .

RUN mkdir gopath
ENV GOPATH /go/gopath
#
ARG GOARCH
COPY key.json key.json
ENV GOOGLE_APPLICATION_CREDENTIALS key.json
ENV GOOS=linux GOARCH=$GOARCH GO111MODULE=on

RUN go build main.go
CMD ["./main"]
