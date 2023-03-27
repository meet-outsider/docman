package data

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"unique_index" binding:"required"`
	Users       []User       `json:"users" gorm:"many2many:user_role;"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permission;"`
}
