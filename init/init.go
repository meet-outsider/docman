package init

import (
	conf "docman/config"
)

func Init() error {
	// 读取yaml配置文件
	err := conf.ReadYaml()
	if err != nil {
		return err
	}
	return nil
}
