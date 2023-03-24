package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/kit"
	"errors"
)

type IUserBiz interface {
	Save(user *data.User) error
	GetByID(id uint) (*data.User, error)
	GetByUsername(username string) (*data.User, error)
	List(pageNum int, pageSize int) ([]*data.User, int64, error)
	Update(user *data.User) error
	DeleteByID(u uint) error
}

type userBiz struct {
	repo repo.IUserRepo
}

func NewUserRepo(repo repo.IUserRepo) IUserBiz {
	return &userBiz{repo}
}

func (s *userBiz) Save(user *data.User) error {
	isExist, _ := s.GetByUsername(user.Username)
	if isExist != nil {
		return errors.New("用户已存在")
	}
	// 密码加密
	user.Password = kit.Encrypt(user.Password)
	_, err := s.repo.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userBiz) GetByID(id uint) (*data.User, error) {
	return s.repo.GetByID(id)
}

func (s *userBiz) GetByUsername(username string) (*data.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *userBiz) List(page int, limit int) ([]*data.User, int64, error) {
	return s.repo.List(page, limit)
}
func (s *userBiz) Update(user *data.User) error {
	return s.repo.Update(user)
}
func (s *userBiz) DeleteByID(u uint) error {
	return s.repo.DeleteByID(u)
}
