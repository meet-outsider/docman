package casbin

import (
	"docman/pkg/database"
	"docman/pkg/log"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Effect *casbin.Enforcer

func NewEnforcer() error {
	a, err := gormadapter.NewAdapterByDB(database.Inst)
	if err != nil {
		return err
	}

	Effect, err = casbin.NewEnforcer("configs/rbac_model.conf", a)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	// 启用自动保存选项。
	Effect.EnableAutoSave(true)
	// 加载策略
	err = Effect.LoadPolicy()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func GetAllPolicy() [][]string {
	policy := Effect.GetPolicy()
	return policy
}
func AddPolicy(sub string, obj string, act string) error {
	_, err := Effect.AddPolicy(sub, obj, act)
	if err != nil {
		return err
	}
	return nil
}
func RemovePolicy(sub string, obj string, act string) error {
	_, err := Effect.RemovePolicy(sub, obj, act)
	if err != nil {
		return err
	}
	return nil
}
func RemoveFilteredPolicy(sub string, obj string, act string) error {
	_, err := Effect.RemoveFilteredPolicy(0, sub, obj, act)
	if err != nil {
		return err
	}
	return nil
}
func RemoveFilteredGroupingPolicy(sub string, obj string, act string) error {
	_, err := Effect.RemoveFilteredGroupingPolicy(0, sub, obj, act)
	if err != nil {
		return err
	}
	return nil
}
