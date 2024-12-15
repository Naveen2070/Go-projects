package service

import (
	"ExpenseTracker/app/model"
	utilities "ExpenseTracker/app/utils"
	"errors"

	"github.com/google/uuid"
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

func (s *AuthService) InitializeTwoFactorAuth(userId uuid.UUID) (utilities.Result, error) {
	result, err := utilities.SetupTwoFactorAuth(userId)
	if err != nil {
		return result, err
	}

	_, err = s.UserService.UpdateUser(userId, model.UserPayload{
		TfaSecret: result.SECRET,
	})

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *AuthService) VerifyTwoFactorAuth(userId uuid.UUID, code string) (bool, error) {
	user, err := s.UserService.GetUserByID(userId)
	if err != nil {
		return false, err
	}

	isVaild := utilities.VerifyTwoFactorAuth(string(user.TfaSecret), code)

	return isVaild, nil
}
