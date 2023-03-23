package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Password string `binding:"required"`
	Nickname string `binding:"required"`
	Email    string `binding:"required"`
	status   int8
	phone    uint    `binding:"required"`
	Roles    []*Role `gorm:"many2many:user_roles;"`
}
