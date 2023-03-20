package model

import (
	"docman/pkg/database"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `validate:"required"`
	Path string `validate:"required"`
}

func (f *File) Find(files *[]File) (e error) {
	database.Db.Find(&files, nil)
	return
}
func (f *File) Update() (e error) {
	if err := database.Db.Create(f).Error; err != nil {
		e = err
	}
	return
}
func (f *File) Create() (e error) {
	if err := database.Db.Create(f).Error; err != nil {
		e = err
	}
	return
}
func (f *File) Delete() (e error) {
	if err := database.Db.Create(f).Error; err != nil {
		e = err
	}
	return
}
