package data

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"unique_index" binding:"required"`
	Permissions []Permission `gorm:"many2many:role_permission;"`
}
