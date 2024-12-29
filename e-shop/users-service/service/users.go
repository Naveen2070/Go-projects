package service

import (
	"context"
	"time"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
)

type UserServiceServer struct {
	userspb.UnimplementedUserServiceServer
	users []*userspb.Users
}

// HealthServiceServer is the implementation of the HealthService
type HealthServiceServer struct {
	userspb.UnimplementedHealthServiceServer
}

// Check implements the HealthService Check method
func (h *HealthServiceServer) Check(ctx context.Context, req *userspb.HealthCheckRequest) (*userspb.HealthCheckResponse, error) {
	return &userspb.HealthCheckResponse{Status: "SERVING"}, nil
}

func genUUID() string {
	uuid := make([]byte, 16)
	return string(uuid)
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.Users, error) {
	user := &userspb.Users{
		UserId:      genUUID(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Username:    req.Username,
		Email:       req.Email,
		LastUpdated: time.Now().String(),
		CreatedAt:   time.Now().String(),
	}
	s.users = append(s.users, user)
	return user, nil
}
