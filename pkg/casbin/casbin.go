package casbin

import (
	"docman/pkg/database"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Effect *casbin.Enforcer

func NewEnforcer() error {
	a, err := gormadapter.NewAdapterByDB(database.Db)
	if err != nil {
		return err
	}

	Effect, err = casbin.NewEnforcer("configs/rbac_model.conf", a)
	if err != nil {
		return err
	}

	return nil
}
