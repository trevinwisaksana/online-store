package handler

import (
	"context"

	pb "github.com/trevinwisaksana/online-store/proto/user"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func NewServer() pb.UserServiceServer {
	return &server{}
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	resp := &pb.GetUserResponse{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "sample@mail.com",
	}

	return resp, nil
}

func RegisterServices(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, NewServer())
}
