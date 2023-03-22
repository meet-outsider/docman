package server

import (
	"docman/config"
	"docman/pkg/casbin"
	"docman/pkg/global"
	"docman/pkg/log"
	"docman/pkg/utils/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"time"
)

func requestHandler(c *gin.Context) {
	verifyToken(c)
}

func verifyToken(c *gin.Context) {
	token := c.Request.Header.Get(global.TOKEN)
	if len(strings.TrimSpace(token)) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"401": "Unauthorized access"})
		return
	}
	sub, exp, err := jwt.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"401": "Unauthorized access"})
		return
	}
	if exp.Before(time.Now()) {
		fmt.Println("过期时间", exp)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"401": "User identity information expired"})
	}
	obj := c.Request.RequestURI
	act := c.Request.Method
	ok, err := casbin.Effect.Enforce(sub, obj, act)
	if err != nil || !ok {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"403": "Forbidden access"})
		return
	}
}

func responseHandler(c *gin.Context) {
	if len(c.Errors) > 0 {
		err := c.Errors.Last()
		statusCode := http.StatusInternalServerError
		if err.Type == gin.ErrorTypePublic {
		}
		c.JSON(statusCode, gin.H{
			"error": err.Err.Error(),
		})
		c.Errors = nil
	}
}

func contains(paths []string, target string) bool {
	for _, path := range paths {
		if strings.Contains(path, target) {
			return true
		}
	}
	return false
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		urls := config.Config.CasbinRules.SkipUrls
		if !contains(urls, c.Request.URL.RequestURI()) {
			requestHandler(c)
		}
		c.Next()
		responseHandler(c)
	}
}

// 自定义校验方法
func customValidator(fl validator.FieldLevel) bool {
	// 校验逻辑
	return false
}

func Run(initRoutes ...func()) error {
	G = gin.Default()
	// 注册自定义校验方法
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("custom", customValidator)
	}
	G.Use(authMiddleware())

	// 路由初始化
	for _, route := range initRoutes {
		route()
	}
	if err := G.Run(fmt.Sprintf(":%d", config.Config.Server.Port)); err != nil {
		return errors.New("server failed")
	}
	return nil
}
