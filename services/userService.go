package services

import (
	"MygarmProject/models"
	"MygarmProject/repositories"
	"strconv"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) LoginUser(user *models.User) error {
	return s.userRepo.LoginUser(user)
}

func (s *UserService) GetUserByID(userID uint) ([]models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(user *models.User) (int64, error) {
	return s.userRepo.UpdateUserByID(user)
}

func (s *UserService) DeleteUserByID(userID string) (int64, error) {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0, err
	}
	return s.userRepo.DeleteUserByID(uint(id))
}
