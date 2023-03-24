package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IPermissionBiz interface {
	Save(c *gin.Context, perm *data.Permission)
	GetByID(c *gin.Context, id uint)
	GetByName(c *gin.Context, name string)
	List(c *gin.Context, pageNum int, pageSize int)
	DeleteById(c *gin.Context, id uint)
}

type permissionBiz struct {
	repo repo.IPermissionRepo
}

func NewPermissionBiz(repo repo.IPermissionRepo) IPermissionBiz {
	return &permissionBiz{repo}
}

func (s *permissionBiz) Save(c *gin.Context, perm *data.Permission) {
	isExist, _ := s.repo.GetByName(perm.Name)
	if isExist != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "权限已存在"})
	}
	savedperm, err := s.repo.Save(perm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"permission": savedperm})
}

func (s *permissionBiz) GetByID(c *gin.Context, id uint) {
	permission, err := s.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "权限不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"permission": permission})
}

func (s *permissionBiz) GetByName(c *gin.Context, name string) {
	permission, err := s.repo.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "权限不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"permission": permission})
}

func (s *permissionBiz) List(c *gin.Context, page int, limit int) {
	list, i, err := s.repo.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(list, i, page, limit))
}

func (s *permissionBiz) DeleteById(c *gin.Context, id uint) {
	err := s.repo.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
