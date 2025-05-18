package main

import (
	"log"
	"net"

	"github.com/trevinwisaksana/online-store/user-service/internal/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting gRPC server...")

	s := grpc.NewServer()
	handler.RegisterServices(s)
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Listening on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Println("gRPC server is running")
}
