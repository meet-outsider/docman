package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `binding:"required,min=4,max=32"gorm:"unique_index"`
	Password string `binding:"required,min=6,max=32"`
	Nickname string `binding:"required,min=2,max=32"`
	Email    string `binding:"required,email"`
	Status   int8
	Phone    uint   `binding:"required,phone"`
	Roles    []Role `gorm:"many2many:user_role;"`
}
