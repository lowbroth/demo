FROM golang:1.25.4-alpine3.21 AS builder

WORKDIR /build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN go build -ldflags="-s -w" -o  demo main.go
RUN go build -o  demo main.go

FROM alpine

WORKDIR /build
ENV TZ=Asia/Shanghai

COPY --from=builder /build/demo /build/demo
COPY --from=builder /build/configs /build/configs
EXPOSE 8080

CMD ["./demo"]

