package service

import (
	connection "ExpenseTracker/app/db"
	"ExpenseTracker/app/model"
	"time"

	"github.com/google/uuid"
)

var db = connection.ConnectDB()

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (model.User, error) {
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (s *UserService) CreateUser(user model.UserPayload) (bool, error) {
	result := db.Create(&model.User{
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (s *UserService) UpdateUser(id uuid.UUID, updatedUser model.UserPayload) (model.User, error) {
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	userToUpdate := model.User{
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		Password:  updatedUser.Password,
		UpdatedAt: time.Now(),
	}
	result = db.Model(&user).Updates(&userToUpdate)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
