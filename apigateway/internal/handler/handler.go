package handler

import (
	"context"

	pb "github.com/trevinwisaksana/online-store/proto/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type ServiceConfig struct {
	Name    string
	Address string
}

func RegisterServices(ctx context.Context, mux *runtime.ServeMux, services []ServiceConfig) error {
	for _, service := range services {
		opts := []grpc.DialOption{grpc.WithInsecure()}
		var err error
		switch service.Name {
		case "UserService":
			err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, service.Address, opts)
		default:
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}
