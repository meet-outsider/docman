package model

import (
	cas "docman/pkg/casbin"
	"docman/pkg/database"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"unique_index"`
}

func CheckPermission(e *casbin.Enforcer, username, obj, act string) (bool, error) {
	roles := make([]string, 0)

	user := &User{}
	if err := database.Inst.Where("username = ?", username).Preload("Roles").First(user).Error; err != nil {
		return false, err
	}
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}
	return cas.Effect.Enforce(username, obj, act, roles)

}
