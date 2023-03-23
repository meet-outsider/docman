package service

import (
	"docman/internal/docman/model"
	"docman/internal/pkg/rsp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func CreateFile(c *gin.Context) {
	var file model.File
	if err := c.BindJSON(&file); err != nil {
		c.JSON(500, "请求body参数错误")
		return
	}
	save := model.File{Name: file.Name, Path: file.Path}
	if err := save.Save(); err != nil {
		rsp.Fail(c, "新增失败")
	} else {
		rsp.Ok(c, "创建成功")
	}
	return
}

func Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var orm = model.File{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	if err := orm.Find(nil); err != nil {
		rsp.Fail(c, err.Error())
		return
	}
	rsp.Ok(c, "ok", "file", orm)
	return
}
func FindFiles(c *gin.Context) {
	var orm = model.File{}
	var files []model.File
	if err := orm.Find(&files); err != nil {
		rsp.Fail(c, err.Error())
		return
	}
	rsp.Ok(c, "ok", "files", files)
	return
}
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var orm = model.File{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	if err := orm.Find(nil); err != nil {
		rsp.Ok(c, err.Error())
		return
	}
	if err := orm.Delete().Error; err != nil {
		rsp.Fail(c, "删除失败！")
		return
	}
	rsp.Ok(c, "删除成功！")
	return
}
