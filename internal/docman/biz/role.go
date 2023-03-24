package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IRoleBiz interface {
	Save(c *gin.Context, role *data.Role)
	GetByID(c *gin.Context, id uint)
	GetByName(c *gin.Context, name string)
	List(c *gin.Context, pageNum int, pageSize int)
	DeleteById(c *gin.Context, id uint)
}

type roleBiz struct {
	repo repo.IRoleRepo
}

func NewRoleBiz(repo repo.IRoleRepo) IRoleBiz {
	return &roleBiz{repo}
}

func (s *roleBiz) Save(c *gin.Context, role *data.Role) {
	isExist, _ := s.repo.GetByName(role.Name)
	if isExist != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "角色已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": role})
}

func (s *roleBiz) GetByID(c *gin.Context, id uint) {
	byID, err := s.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": byID})
}

func (s *roleBiz) GetByName(c *gin.Context, name string) {
	byName, err := s.repo.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": byName})
}

func (s *roleBiz) List(c *gin.Context, pageNum int, pageSize int) {
	list, i, err := s.repo.List(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(list, i, pageNum, pageSize))
}

func (s *roleBiz) DeleteById(c *gin.Context, id uint) {
	err := s.repo.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
