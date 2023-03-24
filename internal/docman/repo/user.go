package repo

import (
	"docman/internal/docman/data"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Save(user *data.User) (*data.User, error)
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

func (r *userRepo) Save(user *data.User) (*data.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetByID(id uint) (*data.User, error) {
	var user data.User
	if err := r.db.Debug().Preload("Roles").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetByUsername(username string) (*data.User, error) {
	var user data.User
	if err := r.db.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) List(page int, limit int) ([]*data.User, int64, error) {
	var users []*data.User
	tx := r.db.Debug().Preload("Roles").Offset((page - 1) * limit).Limit(limit).Find(&users)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return users, tx.RowsAffected, nil
}
func (r *userRepo) Update(user *data.User) error {
	return r.db.Updates(user).Error
}
func (r *userRepo) DeleteByID(u uint) error {
	return r.db.Delete(&data.User{}, u).Error
}
