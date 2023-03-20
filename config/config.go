package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Env     string `yaml:"env"`
		Version string `yaml:"version"`
		Port    uint   `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Logger struct {
		Path       string `yaml:"path"`
		MaxAge     int    `yaml:"max-age"`
		MaxSize    int    `yaml:"max-size"`
		MaxBackups int    `yaml:"max-backups"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

var Config config

func Load() error {
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
	if Config.Server.Port == 0 && port == 0 {
		return fmt.Errorf("port is unset")
	}
	if port != 0 {
		Config.Server.Port = port
	}
	Config.Server.Env = env

	return nil
}
