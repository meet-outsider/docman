package router

import (
	"docman/internal/docman/handler"
	"docman/internal/docman/repository"
	service2 "docman/internal/docman/service"
	"docman/internal/pkg/rsp"
	"docman/pkg/casbin"
	"docman/pkg/database"
	"docman/pkg/server"
	"github.com/gin-gonic/gin"
)

func BindRules() {
	server.GET("/rules", func(c *gin.Context) {
		rules, _ := casbin.Rules()
		rsp.Ok(c, "ok", "rules", rules)
	})
}

func BindUser() {
	v1 := server.Inst.Group("/user")
	userRepo := repository.NewUserRepository(database.Inst)
	userService := service2.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	{
		v1.POST("/login", userHandler.Login)                                  //登陆
		v1.POST("/registry", userHandler.CreateUser)                          //注册
		v1.GET("/getUserInfo", userHandler.GetUserInfo)                       //获取当前用户信息
		v1.GET("/getUserByID/:id", userHandler.GetUserByID)                   //根据id获取
		v1.GET("/getUserByUsername/:username", userHandler.GetUserByUsername) //根据name获取
		v1.GET("/getUsers", userHandler.GetUsers)                             //获取所有用户
	}
}

func BindRole() {
	v1 := server.Inst.Group("/role")
	userRepo := repository.NewRoleRepository(database.Inst)
	userService := service2.NewRoleService(userRepo)
	userHandler := handler.NewRoleHandler(userService)
	{
		v1.POST("/createRole", userHandler.CreateRole)
		v1.GET("/getRoleByID/:id", userHandler.GetRoleByID)
		v1.GET("/getRoleByName/:name", userHandler.GetRoleByName)
		v1.GET("/getRoles", userHandler.GetRoles)
	}
}
func BindPerm() {
	server.GET("/permission", nil)    //权限列表
	server.POST("/permission", nil)   //添加
	server.PUT("/permission", nil)    //修改
	server.DELETE("/permission", nil) //删除

}

func InitDocmanRoutes() {
	BindUser()  //用户服务
	BindRole()  //角色服务
	BindRules() //rules
}
