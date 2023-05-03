package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	biz biz.IRoleBiz
}

func NewRoleHandler(biz biz.IRoleBiz) *RoleHandler {
	return &RoleHandler{biz}
}

func (s *RoleHandler) Create(c *gin.Context) {
	var role data.Role
	ok := kit.BindJson(c, &role)
	if !ok {
		return
	}
	s.biz.Save(c, &role)
}

func (s *RoleHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	s.biz.GetByID(c, uint(id))
}

func (s *RoleHandler) GetByName(c *gin.Context) {
	s.biz.GetByName(c, c.Param("name"))
}

func (s *RoleHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	s.biz.List(c, page, limit)
}

func (s *RoleHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	s.biz.DeleteById(c, uint(id))
}
