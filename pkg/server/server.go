package server

import (
	"docman/config"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run(initRoutes ...func()) error {
	G = gin.Default()
	G.Use(func(c *gin.Context) {
		fmt.Println("x-token", c.Request.Header.Get("x-token"))
		fmt.Println("method", c.Request.Method)
		//log.Info("Before request", zap.Any("request", c.Request))
		c.Next()
		//log.Info("After request", zap.Any("response", c.Writer.Status()))
	})
	// 路由初始化
	for _, route := range initRoutes {
		route()
	}
	if err := G.Run(fmt.Sprintf(":%d", config.Config.Server.Port)); err != nil {
		return errors.New("server failed")
	}
	return nil
}
