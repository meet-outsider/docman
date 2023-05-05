package repo

import (
	"docman/internal/docman/data"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Save(user *data.User) error
	GetByID(id uint) (*data.User, error)
	GetByUsername(username string) (*data.User, error)
	List(page int, limit int, user data.User) ([]*data.User, int64, error)
	Update(user *data.User) error
	DeleteByID(id uint) error
	DeleteByIDs(ids []uint) error
}

var fileds = []string{"id", "username", "email", "nickname", "status", "phone", "created_at", "updated_at"}

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

func (r *userRepo) List(page int, limit int, user data.User) ([]*data.User, int64, error) {
	var users []*data.User
	var count int64
	query := r.db.Model(&data.User{})

	// dynamic query
	// check string is blank
	if len(user.Username) > 0 {
		query = query.Where("username LIKE ?", "%"+user.Username+"%")
	}
	if len(user.Email) > 0 {
		query = query.Where("email LIKE ?", "%"+user.Email+"%")
	}

	err := query.Preload("Roles").Select(fileds).Count(&count).Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *userRepo) Update(user *data.User) error {
	return r.db.Updates(user).Error
}

func (r *userRepo) DeleteByID(id uint) error {
	return r.db.Delete(&data.User{}, id).Error
}

func (r *userRepo) DeleteByIDs(ids []uint) error {
	var users []data.User
	for _, id := range ids {
		users = append(users, data.User{Model: gorm.Model{ID: id}})
	}
	return r.db.Delete(&users).Error
}
