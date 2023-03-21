package router

import (
	"docman/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func Init() {
	BindFile()
	BindAuth()
}
func Run() (e error) {
	Gin = gin.Default()

	Gin.Use(func(c *gin.Context) {
		//log.Info("Before request", zap.Any("request", c.Request))
		c.Next()
		//log.Info("After request", zap.Any("response", c.Writer.Status()))
	})

	Init()
	if err := Gin.Run(fmt.Sprintf(":%d", config.Config.Server.Port)); err != nil {
		return err
	}
	return nil
}
