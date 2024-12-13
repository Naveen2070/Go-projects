package service

import (
	"ExpenseTracker/app/model"
	utilities "ExpenseTracker/app/utils"
	"errors"
)

type AuthService struct {
	UserService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		UserService: userService,
	}
}

func (s *AuthService) Register(user model.UserPayload) (bool, error) {
	return s.UserService.CreateUser(user)
}

func (s *AuthService) Login(user model.AuthPayload) (string, error) {
	email := user.Email

	result, err := s.UserService.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	isValid := utilities.NewPassFactory().ComparePassword(result.Password, user.Password)
	if !isValid {
		return "", errors.New("invalid credentials")
	}

	token, err := utilities.GenerateToken(model.User{
		ID:       result.ID,
		Username: result.Username,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
