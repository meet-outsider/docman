package model

import "gorm.io/gorm"

type CasbinRule struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Ptype string `validate:"required"`
	V0 string `validate:"required"`
}