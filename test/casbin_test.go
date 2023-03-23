package cbsbin_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/casbin/casbin/v2"
)

func Tsest(t *testing.T) {
	// 初始化 casbin Enforcer

	//a, err := gormadapter.NewAdapterByDB(database.Db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	e, err := casbin.NewEnforcer("configs/rbac_model.conf", "configs/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// 加载策略
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal(err)
	}

	// 测试用户 admin 是否有读取 document1 的权限
	sub := "admin"
	obj := "document1"
	act := "read"
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Printf("%s has permission to %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s does not have permission to %s %s\n", sub, act, obj)
	}

	// 测试用户 editor 是否有写入 document3 的权限
	sub = "editor"
	obj = "document3"
	act = "write"
	ok, err = e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Printf("%s has permission to %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s does not have permission to %s %s\n", sub, act, obj)
	}

	// 测试用户 viewer 是否有删除 document2 的权限
	sub = "viewer"
	obj = "document2"
	act = "delete"
	ok, err = e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Printf("%s has permission to %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s does not have permission to %s %s\n", sub, act, obj)
	}
}
