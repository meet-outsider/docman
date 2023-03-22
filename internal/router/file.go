package router

import (
	"docman/internal/service"
	"docman/pkg/log"
	"docman/pkg/server"
	"github.com/gin-gonic/gin"
)

type Point struct {
	X int
	Y int
}

func BindFile() {
	server.G.GET("/file/:id", service.Find)       // 根据id查询
	server.G.GET("/file/list", service.FindFiles) // 查询文件列表

	server.G.POST("/file", service.CreateFile)   // 创建文件
	server.G.DELETE("/file/:id", service.Delete) //删除文件
	server.G.PUT("/file", nil)                   //修改文件
	server.G.GET("/file/clear", func(context *gin.Context) {
		log.Debug("debug")
		log.Info("info", Point{
			X: 10,
			Y: 20,
		})
		log.Warn("warn")
		log.Error("error")
		context.String(200, "name")
	})

}
