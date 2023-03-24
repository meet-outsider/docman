package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserBiz interface {
	Save(c *gin.Context, user *data.User)
	GetByID(c *gin.Context, id uint)
	GetByUsername(c *gin.Context, username string)
	List(c *gin.Context, pageNum int, pageSize int)
	Update(c *gin.Context, user *data.User)
	DeleteByID(c *gin.Context, u uint)
}

type userBiz struct {
	repo repo.IUserRepo
}

func NewUserBiz(repo repo.IUserRepo) IUserBiz {
	return &userBiz{repo}
}

func (s *userBiz) Save(c *gin.Context, user *data.User) {
	isExist, err := s.repo.GetByUsername(user.Username)
	if isExist != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}
	// 密码加密
	user.Password = kit.Encrypt(user.Password)
	err = s.repo.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func (s *userBiz) GetByID(c *gin.Context, id uint) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func (s *userBiz) GetByUsername(c *gin.Context, username string) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func (s *userBiz) List(c *gin.Context, page int, limit int) {
	list, total, err := s.repo.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(list, total, page, limit))
	return
}
func (s *userBiz) Update(c *gin.Context, user *data.User) {
	err := s.repo.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}
func (s *userBiz) DeleteByID(c *gin.Context, u uint) {
	err := s.repo.DeleteByID(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	return
}
