package router

import (
	"docman/internal/service"
)

func BindFile() {
	Gin.GET("/create", service.CreateFile)
	Gin.GET("/list", service.FindFiles)
}
