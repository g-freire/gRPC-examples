package health

import (
	"context"
)

type server struct {
	UnimplementedHealthServiceServer
}

func (s server) GetHealthStatus(_ context.Context, _ *GetHealthStatusRequest) (*GetHealthStatusResponse, error) {
	return &GetHealthStatusResponse{
		Name:  "scaffolding",
		Alive: true,
	}, nil
}

// NewGRPCServer returns the gRPC server to the scaffolding service.
func NewGRPCServer() HealthServiceServer {
	return &server{}
}