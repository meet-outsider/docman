package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Id   int    `gorm:"primaryKey"`
	Name string `validate:"required"`
}
