FROM golang:1.13.5 AS build

WORKDIR /tcs

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

ADD go.mod .
ADD go.sum .
RUN go mod tidy

COPY . .
RUN mkdir -p dist \
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode "external" -extldflags "-static"' -o ./dist/tcs ./cmd/tcs/main.go \
    && mkdir -p /app \
    && mv dist/* /app/


FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
COPY --from=build /app/ .
USER root
CMD ["./tcs"]