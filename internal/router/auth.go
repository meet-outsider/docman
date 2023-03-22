package router

import (
	"docman/internal/rsp"
	"docman/pkg/casbin"
	"docman/pkg/global"
	"docman/pkg/log"
	"docman/pkg/server"
	"docman/pkg/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BindAuth() {
	server.POST("/login", func(c *gin.Context) {
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
	server.POST("/info", func(c *gin.Context) {
		token := c.GetHeader(global.TOKEN)
		userId, exp, err := jwt.ParseToken(token)
		if err != nil {
			rsp.Fail(c, "token解析错误")
			return
		}
		rsp.Ok(c, "获取userId成功", "userId", userId, "exp:", exp)
		return
	})
	server.POST("/policy", func(c *gin.Context) {

	})
	server.GET("/auth", func(c *gin.Context) {
		sub := "admin"
		obj := "document1"
		act := "read"
		ok, err := casbin.Effect.Enforce(sub, obj, act)
		if err != nil {
			println(err.Error())
			log.Error(err.Error())
		}
		if ok {
			rsp.Ok(c, fmt.Sprintf("%s has permission to %s %s\n", sub, act, obj))
		} else {
			rsp.Fail(c, fmt.Sprintf("%s does not have permission to %s %s\n", sub, act, obj))
		}
		return
	})
	server.GET("/cas", func(c *gin.Context) {
		_, e := casbin.Effect.RemovePolicy("admin", "document1", "read")
		if e != nil {
		}
		_, err := casbin.Effect.AddPolicy("admin", "document1", "rrr")
		if err != nil {
			rsp.Fail(c, "save fail")
		} else {
			rsp.Ok(c, "ok")
		}
		return
	})
}
