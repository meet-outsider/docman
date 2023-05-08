package docman

import (
	"docman/cfg"
	"docman/internal/docman/biz"
	"docman/internal/docman/handler"
	"docman/internal/docman/repo"
	"docman/pkg/casbin"
	"docman/pkg/database"
	"docman/pkg/kit"
	"docman/pkg/model"
	"docman/pkg/server"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	prefix := cfg.Config.Server.Api
	v1 := server.Inst.Group(prefix)

	userRepo := repo.NewUserRepo(database.Inst)
	userBiz := biz.NewUserBiz(userRepo)
	userHandler := handler.NewUserHandler(userBiz)
	// user
	{
		v1.GET("/user/:id", userHandler.GetByID)
		v1.GET("/user/username/:username", userHandler.GetByUsername)
		v1.GET("/users/username/:username", userHandler.ListByUsername)
		v1.GET("/users", userHandler.List)
		v1.POST("/user", userHandler.Save)
		v1.PUT("/user/:id", userHandler.Update)
		v1.DELETE("/user/:id", userHandler.DeleteByID)
		v1.DELETE("/users", userHandler.DeleteByIDs)

	}
	roleRepo := repo.NewRoleRepo(database.Inst)
	roleBiz := biz.NewRoleBiz(roleRepo)
	roleHandler := handler.NewRoleHandler(roleBiz)
	// role
	{
		v1.GET("/role/:id", roleHandler.GetByID)
		v1.GET("/role/name/:name", roleHandler.GetByName)
		v1.GET("/roles", roleHandler.List)
		v1.POST("/role", roleHandler.Create)
		v1.DELETE("/role/:id", roleHandler.DeleteByID)
	}
	permissionRepo := repo.NewPermRepo(database.Inst)
	permissionBiz := biz.NewPermissionBiz(permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionBiz)
	// permission
	{
		v1.GET("/perimssions", permissionHandler.List)
		v1.GET("/:id", permissionHandler.GetByID)
		v1.GET("/name/:name", permissionHandler.GetByName)
		v1.POST("/perimssion", permissionHandler.Create)
		v1.PUT("/perimssion/:id", nil)
		v1.DELETE("/:id", permissionHandler.DeleteByID)
	}
	authHandler := handler.NewAuthHandler(biz.NewAuthBiz(userRepo))
	{
		v1.POST("/login", authHandler.Login)
		v1.POST("/registry", authHandler.Registry)
		v1.GET("/info", authHandler.Info) //get current user info
	}
	// casbin
	{
		v1.GET("/casbin/policies", ListPolicies)
		v1.POST("/casbin/policy", savePolicy)
		v1.DELETE("/casbin/policy", deletePolicy)
	}
	// flowable
	flowableHandler := handler.NewFlowableHandler(kit.NewFlowable("http://localhost:9000/flowable-rest", "rest-admin", "test"))
	{
		v1.GET("/flowable/users", flowableHandler.GetUsers)
	}

}

func ListPolicies(ctx *gin.Context) {
	rules := casbin.GetAllPolicy()
	ctx.JSON(http.StatusOK, map[string]any{
		"policies": rules,
	})
}

func savePolicy(ctx *gin.Context) {
	req := model.Policy{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"msg": err.Error(),
		})
		return
	}
	fmt.Printf("%+v\n", req)
	if err := casbin.AddPolicy(req.Sub, req.Obj, req.Act); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func deletePolicy(ctx *gin.Context) {
	req := model.Policy{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"msg": err.Error(),
		})
		return
	}
	casbin.GetAllPolicy()
	if err := casbin.RemovePolicy(req.Sub, req.Obj, req.Act); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
