package service

import (
	"context"
	"errors"
	"time"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
)

type UserServiceServer struct {
	userspb.UnimplementedUserServiceServer
	users []*userspb.User
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
	println(uuid)
	return string(uuid)
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.User, error) {
	user := &userspb.User{
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

func (s *UserServiceServer) GetUserById(ctx context.Context, req *userspb.GetUserByIdRequest) (*userspb.User, error) {
	for _, user := range s.users {
		if user.UserId == req.UserId {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *userspb.GetUsersRequest) (*userspb.Users, error) {
	if len(s.users) == 0 {
		return &userspb.Users{}, errors.New("no users found")
	}
	return &userspb.Users{Users: s.users}, nil
}
