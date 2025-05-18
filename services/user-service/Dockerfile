FROM golang:1.23 as builder

WORKDIR /app

COPY proto ./proto

COPY services/user-service/go.mod services/user-service/go.sum ./
COPY services/user-service ./

RUN sed -i 's|../../proto|./proto|g' go.mod

RUN go mod download

RUN go build -o main ./cmd

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    ca-certificates \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]

EXPOSE 8080