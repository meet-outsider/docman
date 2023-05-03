package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/global"
	"docman/pkg/kit"
	"docman/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthBiz interface {
	Login(c *gin.Context, username string, password string)
	Registry(c *gin.Context, user *data.UserInput)
	Info(c *gin.Context)
}

type authBiz struct {
	userRepo repo.IUserRepo
}

func NewAuthBiz(userRepo repo.IUserRepo) IAuthBiz {
	return &authBiz{
		userRepo: userRepo,
	}
}

func (s *authBiz) Login(c *gin.Context, username string, password string) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	// 检查密码是否正确
	if err := kit.Decrypt(password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}
	roles := make([]string, len(user.Roles))
	for i := range user.Roles {
		roles[i] = user.Roles[i].Name
	}
	userInfo := model.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Roles:    roles,
	}
	token := kit.GenToken(userInfo)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Registry 用户注册
func (s *authBiz) Registry(c *gin.Context, user *data.UserInput) {
	NewUserBiz(s.userRepo).Save(c, user)
}

func (s *authBiz) Info(c *gin.Context) {
	token := c.GetHeader(global.TOKEN)
	subject, _, err := kit.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
		return
	}
	userId := subject.ID
	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
