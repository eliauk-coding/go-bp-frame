FROM golang:1.16-alpine

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/go-bp-frame
COPY . .

RUN go build .

EXPOSE 8999
ENTRYPOINT ["./gobpframe"]
