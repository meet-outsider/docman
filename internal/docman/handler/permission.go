package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PermissionHandler struct {
	biz biz.IPermissionBiz
}

func NewPermissionHandler(biz biz.IPermissionBiz) *PermissionHandler {
	return &PermissionHandler{biz}
}

func (s *PermissionHandler) Create(c *gin.Context) {
	var perm data.Permission
	ok := kit.BindJson(c, &perm)
	if !ok {
		return
	}
	s.biz.Save(c, &perm)
}

func (s *PermissionHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid perm ID"})
		return
	}
	s.biz.GetByID(c, uint(id))
}

func (s *PermissionHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	s.biz.GetByName(c, name)
}

func (s *PermissionHandler) List(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	s.biz.List(c, pageNum, pageSize)
}

func (s *PermissionHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}
	s.biz.DeleteById(c, uint(id))
}
