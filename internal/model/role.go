package model

import (
	"docman/pkg/database"
	"errors"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"unique_index" binding:"required,custom"`
}

func (r *Role) List(roles *[]Role) error {
	var tx *gorm.DB
	if roles == nil {
		// 查单个
		tx = database.Inst.Find(&r)
	} else {
		tx = database.Inst.Find(&roles)
	}
	if tx.RowsAffected == 0 || roles == nil {
		return errors.New("无数据")
	}
	return nil
}

func (r *Role) Get() error {
	return database.Inst.Where(r).Find(r).Error
}

func (r *Role) Create() error {
	return database.Inst.Create(r).Error
}
func (r *Role) SaveBatches(array []Role) error {
	return database.Inst.CreateInBatches(array, len(array)).Error
}
func (r *Role) Update() error {
	return database.Inst.Updates(r).Error
}

func (r Role) Delete() error {
	return database.Inst.Delete(r).Error
}
