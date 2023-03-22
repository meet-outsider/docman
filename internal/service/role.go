package service

import (
	"docman/internal/model"
	"docman/internal/rsp"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role = model.Role{}
	err := c.ShouldBindJSON(&role)
	fmt.Println("role ", role)
	if err != nil {
		rsp.Fail(c, err.Error())
		return
	}
	err = role.Create()
	if err != nil {
		rsp.Fail(c, err.Error())
		return
	}
	rsp.Ok(c, "添加角色成功")
	return
}
func GetRoles(c *gin.Context) {
	var roles *[]model.Role
	role := model.Role{}
	err := role.List(roles)
	if err != nil {
		rsp.Fail(c, err.Error())
	}
	rsp.Ok(c, "添加角色成功", "roles", roles)
	return
}
