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
	v := server.Inst.Group(prefix)

	userRepo := repo.NewUserRepo(database.Inst)
	userBiz := biz.NewUserBiz(userRepo)
	userHandler := handler.NewUserHandler(userBiz)
	userV := v.Group("/users")
	{
		userV.GET("/:id", userHandler.GetByID)                      //根据id查询
		userV.PUT("/:id", userHandler.Update)                       //修改用户信息
		userV.DELETE("/:id", userHandler.DeleteByID)                //根据id删除
		userV.POST("", userHandler.Save)                            //创建用户
		userV.GET("/username/:username", userHandler.GetByUsername) //根据name查询
		userV.GET("", userHandler.List)                             //查询用户列表

	}
	roleRepo := repo.NewRoleRepo(database.Inst)
	roleBiz := biz.NewRoleBiz(roleRepo)
	roleHandler := handler.NewRoleHandler(roleBiz)
	roleV := v.Group("/roles")
	{
		roleV.POST("", roleHandler.Create)              //创建角色
		roleV.GET("/:id", roleHandler.GetByID)          //根据ID查询
		roleV.GET("/name/:name", roleHandler.GetByName) //根据name查询
		roleV.GET("", roleHandler.List)                 //查询所有角色
		roleV.DELETE("/:id", roleHandler.DeleteByID)    //根据ID删除
	}
	permissionRepo := repo.NewPermRepo(database.Inst)
	permissionBiz := biz.NewPermissionBiz(permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionBiz)
	permV := v.Group("/permissions")
	{
		permV.GET("", permissionHandler.List)                 //查询所有权限
		permV.POST("", permissionHandler.Create)              //创建
		permV.PUT("/:id", nil)                                //根据ID更新
		permV.GET("/:id", permissionHandler.GetByID)          //根据ID查询
		permV.GET("/name/:name", permissionHandler.GetByName) //根据name查询
		permV.DELETE("/:id", permissionHandler.DeleteByID)    //根据ID删除角色
	}
	authHandler := handler.NewAuthHandler(biz.NewAuthBiz(userRepo))
	{
		v.POST("/login", authHandler.Login)       //登陆
		v.POST("/registry", authHandler.Registry) //注册用户
		v.GET("/info", authHandler.Info)          //获取当前用户信息
	}
	casbinv := v.Group("/casbin")
	{
		casbinv.GET("/policies", func(c *gin.Context) {
			rules := casbin.GetAllPolicy()
			c.JSON(http.StatusOK, map[string]any{
				"policies": rules,
			})
		})
	}
}
