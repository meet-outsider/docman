package router

import (
	"docman/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func Init() {
	BindFile()
	BindCasbin()
}
func Run() (e error) {
	Gin = gin.New()
	Init()
	if err := Gin.Run(fmt.Sprintf(":%d", config.Config.Server.Port)); err != nil {
		return err
	}
	return nil
}
