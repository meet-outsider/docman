package model

import "gorm.io/gorm"

type CasbinRule struct {
	gorm.Model
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Ptype string `validate:"required"json:"ptype"`
	V0    string `validate:"required"`
}
