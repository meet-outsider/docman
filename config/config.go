package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Env     string `mapstructure:"env"`
		Version string `mapstructure:"version"`
		Port    uint   `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

var Config config

func ReadYaml() error {
	var port = Config.Server.Port
	var env = Config.Server.Env
	// 设置配置文件路径和文件名
	viper.SetConfigName(fmt.Sprintf("docman.%s.yaml", Config.Server.Env))
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将配置文件中的值保存到全局结构体对象中
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct： %s \n", err))
	}
	fmt.Println(Config.Server.Port)
	fmt.Println(port)
	if Config.Server.Port == 0 && port == 0 {
		return fmt.Errorf("port is unset")
	}
	if port != 0 {
		Config.Server.Port = port
	}
	Config.Server.Env = env

	return nil
}
