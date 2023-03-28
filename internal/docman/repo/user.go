package repo

import (
	"docman/internal/docman/data"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Save(user *data.User) error
	GetByID(id uint) (*data.User, error)
	GetByUsername(username string) (*data.User, error)
	List(page int, limit int) ([]*data.User, int64, error)
	Update(user *data.User) error
	DeleteByID(u uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Save(user *data.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetByID(id uint) (*data.User, error) {
	var user data.User
	if err := r.db.Model(&user).Preload("Roles").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByUsername(username string) (*data.User, error) {
	var user data.User
	if err := r.db.Where("username = ?", username).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) List(page int, limit int) ([]*data.User, int64, error) {
	var users []*data.User
	var count int64
	r.db.Model(&data.User{}).Count(&count)
	tx := r.db.Preload("Roles").Offset((page - 1) * limit).Limit(limit).Find(&users)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	return users, count, nil
}
func (r *userRepo) Update(user *data.User) error {
	return r.db.Updates(user).Error
}
func (r *userRepo) DeleteByID(u uint) error {
	return r.db.Delete(&data.User{}, u).Error
}
