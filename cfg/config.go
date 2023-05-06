package cfg

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Env     string
		Version string
		Port    uint
		Api     string
	}
	Database struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
	Logger struct {
		Path       string
		MaxAge     int `mapstructure:"max-age"`
		MaxSize    int `mapstructure:"max-size"`
		MaxBackups int `mapstructure:"max-backups"`
	}
	Jwt struct {
		Secret string
	}
	CasbinRules struct {
		SkipUrls []string `mapstructure:"skip-urls"`
	} `mapstructure:"casbin-rules"`
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
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	// 将配置文件中的值保存到全局结构体对象中
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s ", err))
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
