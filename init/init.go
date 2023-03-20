package init

import (
	conf "docman/config"
	"docman/internal/router"
	"docman/pkg/database"
	"fmt"
)

func Init() error {
	if err := conf.Load(); err != nil {
		return fmt.Errorf("读取配置文件失败")
	}
	if err := database.Connect(); err != nil {
		return err
	}
	if err := router.Run(); err != nil {
		return err
	}
	return nil
}
