package data

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Id    int `gorm:"primaryKey"`
	Name  string
	Path  string
	IsDir bool
}
