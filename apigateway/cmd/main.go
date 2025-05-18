package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/trevinwisaksana/online-store/apigateway/internal/handler"
)

func main() {
	log.Println("Starting API Gateway...")

	services := []handler.ServiceConfig{
		{
			Name:    "UserService",
			Address: "user-service:50051",
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := handler.RegisterServices(ctx, mux, services); err != nil {
		log.Fatalf("Failed to register services %v", err)
	}

	log.Println("Services registered successfully. Starting HTTP server on port 8080...")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve HTTP %v", err)
	}

	log.Println("API Gateway is running")
}
