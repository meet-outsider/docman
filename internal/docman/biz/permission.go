package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"errors"
)

type IPermissionBiz interface {
	Save(perm *data.Permission) (*data.Permission, error)
	GetByID(id uint) (*data.Permission, error)
	GetByName(name string) (*data.Permission, error)
	List(pageNum int, pageSize int) ([]*data.Permission, int64, error)
	DeleteById(id uint) error
}

type permissionBiz struct {
	repo repo.IPermissionRepo
}

func NewPermissionBiz(repo repo.IPermissionRepo) IPermissionBiz {
	return &permissionBiz{repo}
}

func (s *permissionBiz) Save(perm *data.Permission) (*data.Permission, error) {
	isExist, _ := s.GetByName(perm.Name)
	if isExist != nil {
		return nil, errors.New("角色已存在")
	}
	return s.repo.Save(perm)
}

func (s *permissionBiz) GetByID(id uint) (*data.Permission, error) {
	return s.repo.GetByID(id)
}

func (s *permissionBiz) GetByName(name string) (*data.Permission, error) {
	return s.repo.GetByName(name)
}

func (s *permissionBiz) List(pageNum int, pageSize int) ([]*data.Permission, int64, error) {
	return s.repo.List(pageNum, pageSize)
}

func (s *permissionBiz) DeleteById(id uint) error {
	return s.repo.DeleteById(id)
}
