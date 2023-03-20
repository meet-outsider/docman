package model

import (
	"docman/pkg/database"
	"errors"
	"gorm.io/gorm"
	"time"
)

type File struct {
	gorm.Model
	Name  string `validate:"required"`
	Path  string `validate:"required"`
	IsDir bool   `validate:"required"`
}

func (f *File) Find(files *[]File) (err error) {
	var tx *gorm.DB
	if files == nil {
		// 查单个
		tx = database.Db.Find(&f)
	} else {
		tx = database.Db.Find(&files)
	}
	if tx.RowsAffected == 0 || files == nil {
		err = errors.New("无数据")
	}
	return
}
func (f *File) Update() error {
	return database.Db.Create(f).Error
}
func (f *File) Create() error {
	return database.Db.Create(f).Error
}
func (f *File) Delete() error {
	return database.Db.Delete(f).Error
}

// Clear 删除30天之前的逻辑删除字段
func (f *File) Clear() (err error) {
	return database.Db.Unscoped().Delete(f).
		Where("created_at < ?", time.Now().AddDate(0, 0, -30)).Error
}
