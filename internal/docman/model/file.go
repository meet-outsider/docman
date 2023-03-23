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
		tx = database.Inst.Find(f)
	} else {
		tx = database.Inst.Find(files)
	}
	if tx.RowsAffected == 0 || files == nil {
		err = errors.New("无数据")
	}
	return
}
func (f *File) Update() error {
	return f.Update()
	//return database.Inst.Model(f).Where("id", 1).Update("", "").Error
}
func (f *File) Save() error {
	return database.Inst.Create(f).Error
}
func (f *File) Delete() error {
	return database.Inst.Delete(f).Error
}

// Clear 删除30天之前的逻辑删除字段
func (f *File) Clear() (err error) {
	return database.Inst.Unscoped().Delete(f).
		Where("created_at < ?", time.Now().AddDate(0, 0, -30)).Error
}
