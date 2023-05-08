package server

import (
	"docman/cfg"
	"docman/pkg/casbin"
	"docman/pkg/global"
	"docman/pkg/kit"
	"docman/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var Inst *gin.Engine

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := strings.TrimPrefix(ctx.Request.RequestURI, cfg.Config.Server.Api)
				method := ctx.Request.Method
				cause := fmt.Sprintf("%v", err)
				log.Error(
					fmt.Sprintf("url: %s, method: %s cause:%s", url, method, cause))
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		ctx.Next()
	}
}

func verifyToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get(global.TOKEN)
	if len(strings.TrimSpace(token)) == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"401": "Unauthorized access / 未授权访问"})
		return
	}
	subject, exp, err := kit.ParseToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access / 未授权访问"})
		return
	}
	if exp.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User identity expired / 用户身份过期"})
	}
	subs := subject.Roles
	obj := ctx.Request.RequestURI
	act := ctx.Request.Method
	for i := range subs {
		sub := subs[i]
		ok, err := casbin.Effect.Enforce(sub, obj, act)
		if err != nil {
			log.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden access / 禁止访问"})
			return
		}
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden access / 禁止访问"})
			return
		}
	}
}

func requestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skipUrls := cfg.Config.CasbinRules.SkipUrls
		url := strings.TrimPrefix(ctx.Request.URL.RequestURI(), cfg.Config.Server.Api)
		currUrl := url + "::" + ctx.Request.Method
		contains := func(paths []string, target string) bool {
			for _, path := range paths {
				if strings.Contains(path, target) {
					return true
				}
			}
			return false
		}
		if !contains(skipUrls, currUrl) {
			verifyToken(ctx)
		}
		ctx.Next()
	}
}
func Init() {
	Inst = gin.Default()
	Inst.Use(requestHandler())
	Inst.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "404 not found"})
	})
	Inst.Use(ErrorHandler())
}
