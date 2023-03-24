package docman

import (
	"docman/cfg"
	"docman/internal/docman/biz"
	"docman/internal/docman/handler"
	"docman/internal/docman/repo"
	"docman/pkg/casbin"
	"docman/pkg/database"
	"docman/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() {
	prefix := cfg.Config.Server.Api
	v1 := server.Inst.Group(prefix)

	userRepo := repo.NewUserRepo(database.Inst)
	userBiz := biz.NewUserRepo(userRepo)
	userHandler := handler.NewUserHandler(userBiz)
	userV1 := v1.Group("/users")
	{
		v1.POST("/login", userHandler.Login)       //登陆
		v1.POST("/registry", userHandler.Registry) //注册用户
		v1.GET("/info", userHandler.Info)          //获取当前用户信息
	}
	{
		userV1.GET("/:id", userHandler.GetByID)                      //根据id查询
		userV1.PUT("/:id", userHandler.Update)                       //修改用户信息
		userV1.DELETE("/:id", userHandler.DeleteByID)                //根据id删除
		userV1.POST("", userHandler.Registry)                        //创建用户
		userV1.GET("/username/:username", userHandler.GetByUsername) //根据name查询
		userV1.GET("", userHandler.List)                             //查询用户列表

	}
	roleRepo := repo.NewRoleRepo(database.Inst)
	roleBiz := biz.NewRoleBiz(roleRepo)
	roleHandler := handler.NewRoleHandler(roleBiz)
	roleV1 := v1.Group("/roles")
	{
		roleV1.POST("", roleHandler.Create)              //创建角色
		roleV1.GET("/:id", roleHandler.GetByID)          //根据ID查询
		roleV1.GET("/name/:name", roleHandler.GetByName) //根据name查询
		roleV1.GET("", roleHandler.List)                 //查询所有角色
		roleV1.DELETE("/:id", roleHandler.DeleteByID)    //根据ID删除
	}
	permissionRepo := repo.NewPermRepo(database.Inst)
	permissionBiz := biz.NewPermissionBiz(permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionBiz)
	permV1 := v1.Group("/permissions")
	{
		permV1.GET("", permissionHandler.List)                 //查询所有权限
		permV1.POST("", permissionHandler.Create)              //创建
		permV1.PUT("/:id", nil)                                //根据ID更新
		permV1.GET("/:id", permissionHandler.GetByID)          //根据ID查询
		permV1.GET("/name/:name", permissionHandler.GetByName) //根据name查询
		permV1.DELETE("/:id", permissionHandler.DeleteByID)    //根据ID删除角色
	}
	casbinV1 := v1.Group("/casbin")
	{
		casbinV1.GET("/policies", func(c *gin.Context) {
			rules := casbin.GetAllPolicy()
			c.JSON(http.StatusOK, map[string]any{
				"policies": rules,
			})
		})
	}
}
