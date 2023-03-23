package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string        `gorm:"unique_index" binding:"required,custom"`
	Permissions []*Permission `gorm:"many2many:role_permission;"`
}
