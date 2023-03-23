package repository

import (
	"docman/internal/docman/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Save(user *model.User) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindAll(pageNum int, pageSize int) ([]*model.User, int64, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) Save(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindAll(pageNum int, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	tx := r.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return users, tx.RowsAffected, nil
}
