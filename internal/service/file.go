package service

import (
	"docman/internal/model"
	"docman/internal/rsp"
	"github.com/gin-gonic/gin"
)

func CreateFile(c *gin.Context) {
	var file *model.File
	if err := c.ShouldBindJSON(&file); err != nil && file != nil {
		c.JSON(500, "参数错误")
		return
	}
	save := model.File{Name: file.Name, Path: file.Path}
	if err := save.Create(); err != nil {
		rsp.Fail(c, "错误测试")
	} else {
		rsp.Ok(c, "成功", "")
	}
	return
}

func Find(c *gin.Context) {
	var find = model.File{}
	var files []model.File
	err := find.Find(&files)
	if err != nil {
		return
	}
	rsp.Ok(c, "ok", files)
	return
}
func FindFiles(c *gin.Context) {
	var find = model.File{}
	var files []model.File
	err := find.Find(&files)
	if err != nil {
		return
	}
	rsp.Ok(c, "ok", files)
	return
}
