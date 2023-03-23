package repository

import (
	"docman/internal/docman/model"
	"gorm.io/gorm"
)

type RoleRepo interface {
	Save(role *model.Role) (*model.Role, error)
	FindByID(id uint) (*model.Role, error)
	FindByName(name string) (*model.Role, error)
	FindAll(pageNum int, pageSize int) ([]*model.Role, error)
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepo {
	return &roleRepo{db}
}

func (r *roleRepo) Save(role *model.Role) (*model.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepo) FindByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) FindByName(name string) (*model.Role, error) {
	var role model.Role
	if err := r.db.Where("rolename = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) FindAll(pageNum int, pageSize int) ([]*model.Role, error) {
	var roles []*model.Role
	if err := r.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
