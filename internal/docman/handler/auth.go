package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthHandler struct {
	biz biz.IAuthBiz
}

func NewAuthHandler(biz biz.IAuthBiz) *AuthHandler {
	return &AuthHandler{biz}
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var params map[string]string
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": kit.Translate(err)})
		return
	}
	username := params["username"]
	password := params["password"]
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不能为空"})
		return
	}
	fmt.Println("handler user", username, password)
	h.biz.Login(c, username, password)
}

// Registry 用户注册
func (h *AuthHandler) Registry(c *gin.Context) {
	var user data.UserInput
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": kit.Translate(err)})
		return
	}
	h.biz.Registry(c, &user)
}

func (h *AuthHandler) Info(c *gin.Context) {
	h.biz.Info(c)
}
