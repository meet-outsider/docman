package service

import (
	"docman/internal/docman/model"
	"docman/internal/docman/repository"
)

type RoleService interface {
	CreateRole(role *model.Role) (*model.Role, error)
	GetRoleByID(id uint) (*model.Role, error)
	GetRoleByName(name string) (*model.Role, error)
	GetRoles(pageNum int, pageSize int) ([]*model.Role, error)
}

type roleService struct {
	repo repository.RoleRepo
}

func NewRoleService(repo repository.RoleRepo) RoleService {
	return &roleService{repo}
}

func (s *roleService) CreateRole(user *model.Role) (*model.Role, error) {
	return s.repo.Save(user)
}

func (s *roleService) GetRoleByID(id uint) (*model.Role, error) {
	return s.repo.FindByID(id)
}

func (s *roleService) GetRoleByName(name string) (*model.Role, error) {
	return s.repo.FindByName(name)
}

func (s *roleService) GetRoles(pageNum int, pageSize int) ([]*model.Role, error) {
	return s.repo.FindAll(pageNum, pageSize)
}
