package init

import (
	"docman/cfg"
	"docman/internal/docman"
	"docman/pkg/casbin"
	"docman/pkg/database"
	"docman/pkg/kit"
	"docman/pkg/log"
	"docman/pkg/server"
	"errors"
	"fmt"
)

func Init() error {
	/**
	1、加载配置文件
	2、链接数据库
	3、启动casbin权限控制服务
	4、启动应用
	*/
	if err := cfg.Load(); err != nil {
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
	if err := kit.Init(); err != nil {
		log.Error(err.Error())
		return errors.New("validator init failed")
	}
	if err := server.Run(docman.InitRoutes, func() {
	}); err != nil {
		log.Error(err.Error())
		return errors.New("server run failed")
	}
	return nil
}
