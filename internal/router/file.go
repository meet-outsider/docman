package router

import (
	"docman/internal/service"
)

func BindFile() {
	Gin.GET("/file/:id", service.Find)       // 根据id查询
	Gin.GET("/file/list", service.FindFiles) // 查询文件列表

	Gin.POST("/file", service.CreateFile)   // 创建文件
	Gin.DELETE("/file/:id", service.Delete) //删除文件
	Gin.PUT("/file", nil)                   //修改文件
	Gin.PUT("/file/clear", nil)

}
