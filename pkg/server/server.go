package server

import (
	"docman/cfg"
	"docman/pkg/casbin"
	"docman/pkg/global"
	"docman/pkg/kit"
	"docman/pkg/log"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := strings.TrimPrefix(c.Request.RequestURI, cfg.Config.Server.Api)
				method := c.Request.Method
				cause := fmt.Sprintf("%v", err)
				log.Error(
					fmt.Sprintf("url: %s, method: %s cause:%s", url, method, cause))
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}

func verifyToken(c *gin.Context) {
	token := c.Request.Header.Get(global.TOKEN)
	if len(strings.TrimSpace(token)) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"401": "Unauthorized access / 未授权访问"})
		return
	}
	subject, exp, err := kit.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access / 未授权访问"})
		return
	}
	if exp.Before(time.Now()) {
		fmt.Println("过期时间", exp)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User identity expired / 用户身份过期"})
	}
	subs := subject.Roles
	obj := c.Request.RequestURI
	act := c.Request.Method
	for i := range subs {
		sub := subs[i]
		ok, err := casbin.Effect.Enforce(sub, obj, act)
		fmt.Printf("sub: %s, obj: %s, act: %s, ok: %v, err: %v", sub, obj, act, ok, err)
		if err != nil {
			log.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden access / 禁止访问"})
			return
		}
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden access / 禁止访问"})
			return
		}
	}
}

func requestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		skipUrls := cfg.Config.CasbinRules.SkipUrls
		url := strings.TrimPrefix(c.Request.URL.RequestURI(), cfg.Config.Server.Api)
		currUrl := url + "::" + c.Request.Method
		contains := func(paths []string, target string) bool {
			for _, path := range paths {
				if strings.Contains(path, target) {
					return true
				}
			}
			return false
		}
		if !contains(skipUrls, currUrl) {
			verifyToken(c)
		}
		c.Next()
	}
}
func Run(initRoutes ...func()) error {
	Inst = gin.Default()
	Inst.Use(requestHandler())
	Inst.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 not found"})
	})
	Inst.Use(ErrorHandler())
	for _, route := range initRoutes {
		route()
	}
	if err := Inst.Run(fmt.Sprintf(":%d", cfg.Config.Server.Port)); err != nil {
		return errors.New("server failed")
	}
	return nil
}
