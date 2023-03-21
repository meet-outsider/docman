package init

import (
	conf "docman/config"
	"docman/internal/router"
	"docman/pkg/casbin"
	"docman/pkg/database"
	"docman/pkg/log"
	"errors"
	"fmt"
)

func Init() error {
	if err := conf.Load(); err != nil {
		log.Error(err.Error())
		return fmt.Errorf("读取配置文件失败")
	}
	if err := database.Connect(); err != nil {
		log.Error(err.Error())
		return errors.New("db disconnected")
	}
	if err := casbin.NewEnforcer(); err != nil {
		log.Error(err.Error())
		return errors.New("casbin init failed")
	}
	if err := router.Run(); err != nil {
		log.Error(err.Error())
		return errors.New("server run failed")
	}
	return nil
}
