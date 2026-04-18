# syntax=docker/dockerfile:1.4

FROM golang:1.25-bullseye AS builder

WORKDIR /build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -o demo main.go

FROM alpine

WORKDIR /build
ENV TZ=Asia/Shanghai

COPY --from=builder /build/demo /build/demo
COPY --from=builder /build/configs /build/configs

CMD ["./demo"]
