package repo

import (
	"docman/internal/docman/data"
	"gorm.io/gorm"
)

type IRoleRepo interface {
	Save(role *data.Role) (*data.Role, error)
	GetByID(id uint) (*data.Role, error)
	GetByName(name string) (*data.Role, error)
	List(pageNum int, pageSize int) ([]*data.Role, int64, error)
	DeleteById(id uint) error
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) IRoleRepo {
	return &roleRepo{db: db}
}

func (r *roleRepo) Save(role *data.Role) (*data.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepo) GetByID(id uint) (*data.Role, error) {
	var role data.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) GetByName(name string) (*data.Role, error) {
	var role data.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) List(pageNum int, pageSize int) ([]*data.Role, int64, error) {
	var roles []*data.Role
	tx := r.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&roles)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return roles, tx.RowsAffected, nil
}

func (r *roleRepo) DeleteById(id uint) error {
	var role data.Role
	return r.db.Where("id = ?", id).Delete(&role).Error
}
