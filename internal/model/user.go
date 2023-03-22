package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Password string
	Nickname string
	Email    string
	status   int8
	phone    uint
	Roles    []*Role `gorm:"many2many:user_roles;"`
}

func CreateUser(db *gorm.DB, roleName, username, password string) error {
	role := &Role{Name: roleName}
	if err := db.Create(role).Error; err != nil {
		return err
	}

	user := &User{Username: username, Password: password}
	user.Roles = append(user.Roles, role)
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
