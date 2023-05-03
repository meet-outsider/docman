// @description User model and input model
// @author outsider
// @date 2023-04-01
// @updated 2023-05-02
package data

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Nickname string `json:"nickname" binding:"required,min=2,max=32"`
	Email    string `json:"email" binding:"required,email"`
	// 0: normal, 1: disabled
	Status   int8
	Phone    uint   `json:"phone" binding:"required,phone"`
	Roles    []Role `json:"roles" gorm:"many2many:user_role;"`
}

// UserInput model
type UserInput struct {
	User
	Roles []uint `json:"roles" binding:"required"`
}
