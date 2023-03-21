package router

import (
	"docman/internal/rsp"
	"docman/pkg/global"
	"docman/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func BindAuth() {
	Gin.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if len(username) == 0 || len(password) == 0 {
			rsp.Fail(c, "参数校验失败")
			return
		}
		token := jwt.GenToken(11)
		rsp.Ok(c, "登陆成功，获取token", "token", token)
		return
	})
	Gin.POST("/info", func(c *gin.Context) {
		token := c.GetHeader(global.TOKEN)
		userId, exp, err := jwt.ParseToken(token)
		if err != nil {
			rsp.Fail(c, "token解析错误")
			return
		}
		rsp.Ok(c, "获取userId成功", "userId", userId, "exp:", exp)
		return
	})
}
