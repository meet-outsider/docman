package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Nickname string `json:"nickname" binding:"required,min=2,max=32"`
	Email    string `json:"email" binding:"required,email"`
	Status   int8
	Phone    uint   `json:"phone" binding:"required,phone"`
	Roles    []Role `json:"roles" gorm:"many2many:user_role;"`
}
type UserInput struct {
	User
	Roles []uint `binding:"required"`
}
