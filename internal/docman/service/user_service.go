package service

import (
	"docman/internal/docman/model"
	"docman/internal/docman/repository"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUsers(pageNum int, pageSize int) ([]*model.User, int64, error)
}

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	return s.repo.Save(user)
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *userService) GetUsers(pageNum int, pageSize int) ([]*model.User, int64, error) {
	return s.repo.FindAll(pageNum, pageSize)
}
