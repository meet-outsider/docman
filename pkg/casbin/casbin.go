package casbin

import (
	"docman/pkg/database"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var Effect *casbin.Enforcer

func NewEnforcer() error {
	a, err := gormadapter.NewAdapterByDB(database.Inst)
	if err != nil {
		return err
	}

	Effect, err = casbin.NewEnforcer("configs/rbac_model.conf", a)
	if err != nil {
		return err
	}
	// 启用自动保存选项。
	Effect.EnableAutoSave(true)
	// 加载策略
	err = Effect.LoadPolicy()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
