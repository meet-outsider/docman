package router

import (
	"docman/internal/rsp"
	"docman/internal/service"
	"docman/pkg/casbin"
	"docman/pkg/server"
	"github.com/gin-gonic/gin"
)

func BindRules() {
	server.GET("/rules", func(c *gin.Context) {
		rules, _ := casbin.Rules()
		rsp.Ok(c, "ok", "rules", rules)
	})
}

func BindAuth() {
	server.POST("/login", service.Login)       //登陆
	server.GET("/info", service.Info)          //获取当前用户注册
	server.POST("/registry", service.Registry) //注册

}

func BindUser() {
	//server.POST("/login", service.Login)       //登陆
	//server.GET("/info", service.Info)          //获取当前用户注册
	//server.POST("/registry", service.Registry) //注册
}

func BindRole() {
	server.G.GET("/role/:id", nil)
	server.G.POST("/role", service.CreateRole)
	server.G.PUT("/role", nil)
	server.G.DELETE("/role/:id", nil)
	server.G.GET("/role/list", service.GetRoles)
}

func InitDocmanRoutes() {
	BindAuth()  //权限服务
	BindUser()  //用户服务
	BindRole()  //角色服务
	BindRules() //rules
}
