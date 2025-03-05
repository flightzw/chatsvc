FROM registry.cn-hangzhou.aliyuncs.com/flightzw/golang:1.23-alpine3.20 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.io make build

FROM alpine:3.20

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000

VOLUME /data/conf

CMD ["./chatsvc", "-conf", "/data/conf"]
