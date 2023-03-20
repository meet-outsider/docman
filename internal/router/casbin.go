package router

import (
	"docman/internal/rsp"
	"docman/pkg/global"
	"docman/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func BindCasbin() {
	Gin.GET("/login/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			rsp.Fail(c, "智能为整数", nil)
			return
		}
		token := jwt.GenToken(id)
		rsp.Ok(c, "登陆成功，获取token", token)
		return
	})
	Gin.POST("/info", func(c *gin.Context) {
		token := c.GetHeader(global.TOKEN)
		fmt.Println("token", token)
		userId, exp, err := jwt.ParseToken(token)
		if err != nil {
			rsp.Fail(c, "token解析错误", err)
			return
		}
		rsp.Ok(c, "获取userId成功", "userId", userId, "exp:", exp)
		return
	})
}
