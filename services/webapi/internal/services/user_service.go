package services

import (
	"github.com/bit365/iotcloud/services/webapi/internal/models"
	"github.com/bit365/iotcloud/services/webapi/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type UserServiceImpl struct {
	repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repo}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *models.User) error {
	err := s.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) UpdateUser(user *models.User) error {
	err := s.repository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
