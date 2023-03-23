package server

import (
	"docman/config"
	"docman/pkg/casbin"
	"docman/pkg/global"
	"docman/pkg/log"
	"docman/pkg/utils"
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
				log.Error(fmt.Sprintf("url: %s, method: %s cause:%s", c.Request.RequestURI, c.Request.Method, err))
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"401": "Unauthorized access"})
		return
	}
	sub, exp, err := utils.ParseToken(token)
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
	if err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"403": "Forbidden access"})
		return
	}
	if !ok {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"403": "Forbidden access"})
		return
	}
}

func requestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		urls := config.Config.CasbinRules.SkipUrls
		currUrl := c.Request.URL.RequestURI() + "::" + c.Request.Method
		contains := func(paths []string, target string) bool {
			for _, path := range paths {
				if strings.Contains(path, target) {
					return true
				}
			}
			return false
		}
		if !contains(urls, currUrl) {
			verifyToken(c)
		}
		c.Next()
	}
}

func Run(initRoutes ...func()) error {
	Inst = gin.Default()
	//G.Use(requestHandler())
	Inst.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "not found")
	})
	Inst.Use(ErrorHandler())
	for _, route := range initRoutes {
		route()
	}
	if err := Inst.Run(fmt.Sprintf(":%d", config.Config.Server.Port)); err != nil {
		return errors.New("server failed")
	}
	return nil
}
