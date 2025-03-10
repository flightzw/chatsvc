FROM registry.cn-hangzhou.aliyuncs.com/flightzw/golang:1.23-alpine3.20 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.io make build

FROM registry.cn-hangzhou.aliyuncs.com/flightzw/alpine:3.20

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000

VOLUME /data/conf

CMD ["./chatsvc", "-conf", "/data/conf"]
