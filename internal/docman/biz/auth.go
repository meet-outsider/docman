package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/database"
	"docman/pkg/global"
	"docman/pkg/kit"
	"docman/pkg/model"
	"fmt"
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

func (h *authBiz) Login(c *gin.Context, username string, password string) {
	user, err := h.userRepo.GetByUsername(username)
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
	fmt.Println("userInfo", userInfo)
	token := kit.GenToken(userInfo)
	c.JSON(http.StatusOK, gin.H{"token": token})
	return
}

// Registry 用户注册
func (h *authBiz) Registry(c *gin.Context, user *data.UserInput) {
	save := user.User
	err := h.userRepo.Save(&save)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	roleIDs := user.Roles
	// 保存用户角色关系
	tx := database.Inst.Begin()
	if len(roleIDs) != 0 {
		for _, id := range roleIDs {
			_, err := h.userRepo.GetByID(id)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
				return
			}
			e := database.Inst.Save(data.UserRole{
				UserID: save.ID,
				RoleID: id,
			}).Error
			if e != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "保存用户角色关系失败"})
				return
			}
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	return
}

func (h *authBiz) Info(c *gin.Context) {
	token := c.GetHeader(global.TOKEN)
	subject, _, err := kit.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
		return
	}
	userId := subject.ID
	user, err := h.userRepo.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}
