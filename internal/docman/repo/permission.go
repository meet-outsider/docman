package repo

import (
	"docman/internal/docman/data"
	"gorm.io/gorm"
)

type IPermissionRepo interface {
	Save(perm *data.Permission) (*data.Permission, error)
	GetByID(id uint) (*data.Permission, error)
	GetByName(name string) (*data.Permission, error)
	List(pageNum int, pageSize int) ([]*data.Permission, int64, error)
	DeleteById(id uint) error
}

type permissionRepo struct {
	db *gorm.DB
}

func NewPermRepo(db *gorm.DB) IPermissionRepo {
	return &permissionRepo{db: db}
}

func (r *permissionRepo) Save(perm *data.Permission) (*data.Permission, error) {
	if err := r.db.Create(perm).Error; err != nil {
		return nil, err
	}
	return perm, nil
}

func (r *permissionRepo) GetByID(id uint) (*data.Permission, error) {
	var perm data.Permission
	if err := r.db.Where("id = ?", id).First(&perm).Error; err != nil {
		return nil, err
	}
	return &perm, nil
}

func (r *permissionRepo) GetByName(name string) (*data.Permission, error) {
	var perm data.Permission
	if err := r.db.Where("name = ?", name).First(&perm).Error; err != nil {
		return nil, err
	}
	return &perm, nil
}
func (r *permissionRepo) List(pageNum int, pageSize int) ([]*data.Permission, int64, error) {
	var perms []*data.Permission
	tx := r.db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&perms)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return perms, tx.RowsAffected, nil
}

func (r *permissionRepo) DeleteById(id uint) error {
	var perm data.Permission
	return r.db.Where("id = ?", id).Delete(&perm).Error
}
