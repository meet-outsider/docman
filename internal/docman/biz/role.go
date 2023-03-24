package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"errors"
)

type IRoleBiz interface {
	Save(role *data.Role) (*data.Role, error)
	GetByID(id uint) (*data.Role, error)
	GetByName(name string) (*data.Role, error)
	List(pageNum int, pageSize int) ([]*data.Role, int64, error)
	DeleteById(id uint) error
}

type roleBiz struct {
	repo repo.IRoleRepo
}

func NewRoleBiz(repo repo.IRoleRepo) IRoleBiz {
	return &roleBiz{repo}
}

func (s *roleBiz) Save(role *data.Role) (*data.Role, error) {
	isExist, _ := s.GetByName(role.Name)
	if isExist != nil {
		return nil, errors.New("角色已存在")
	}
	return s.repo.Save(role)
}

func (s *roleBiz) GetByID(id uint) (*data.Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleBiz) GetByName(name string) (*data.Role, error) {
	return s.repo.GetByName(name)
}

func (s *roleBiz) List(pageNum int, pageSize int) ([]*data.Role, int64, error) {
	return s.repo.List(pageNum, pageSize)
}

func (s *roleBiz) DeleteById(id uint) error {
	return s.repo.DeleteById(id)
}
