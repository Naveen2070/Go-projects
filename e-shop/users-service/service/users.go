package service

import (
	"context"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
)

type UserServiceServer struct {
	userspb.UnimplementedUserServiceServer
	users []userspb.Users
}

// HealthServiceServer is the implementation of the HealthService
type HealthServiceServer struct {
	userspb.UnimplementedHealthServiceServer
}

// Check implements the HealthService Check method
func (h *HealthServiceServer) Check(ctx context.Context, req *userspb.HealthCheckRequest) (*userspb.HealthCheckResponse, error) {
	return &userspb.HealthCheckResponse{Status: "SERVING"}, nil
}
