FROM golang:1.23 AS builder

WORKDIR /app

# Copy proto directory first
COPY proto ./proto

# Copy service files
COPY apigateway/go.mod apigateway/go.sum ./
COPY apigateway ./

# Update the replace directive to use the correct path
RUN sed -i 's|replace github.com/trevinwisaksana/online-store/proto => .*|replace github.com/trevinwisaksana/online-store/proto => ./proto|g' go.mod

RUN go mod download

RUN go build -o main ./cmd

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]