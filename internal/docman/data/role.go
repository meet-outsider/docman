package data

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"unique_index" binding:"required"`
	Users       []User       `gorm:"many2many:user_role;"`
	Permissions []Permission `gorm:"many2many:role_permission;"`
}
