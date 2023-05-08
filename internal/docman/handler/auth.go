// @description 用户登录注册相关接口
// @author outsider
// @date 2023-04-01
// @updated 2023-05-02
package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
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
func (s *AuthHandler) Login(ctx *gin.Context) {
	var params map[string]string
	ok := kit.BindJson(ctx, &params)
	if !ok {
		return
	}
	username := params["username"]
	password := params["password"]
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不能为空"})
		return
	}
	s.biz.Login(ctx, username, password)
}

// Registry 用户注册
func (s *AuthHandler) Registry(ctx *gin.Context) {
	var user data.UserInput
	ok := kit.BindJson(ctx, &user)
	if !ok {
		return
	}
	// 注册用户为默认角色
	user.Roles = []uint{1}
	s.biz.Registry(ctx, &user)
}

func (s *AuthHandler) Info(ctx *gin.Context) {
	s.biz.Info(ctx)
}
