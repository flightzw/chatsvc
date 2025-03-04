FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM alpine:3.18

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000

VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
