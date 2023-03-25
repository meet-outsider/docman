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

func (h *RoleHandler) Create(c *gin.Context) {
	var role data.Role
	ok := kit.BindJson(c, &role)
	if !ok {
		return
	}
	h.biz.Save(c, &role)
}

func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	h.biz.GetByID(c, uint(id))
}

func (h *RoleHandler) GetByName(c *gin.Context) {
	h.biz.GetByName(c, c.Param("name"))
}

func (h *RoleHandler) List(c *gin.Context) {
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
	h.biz.List(c, page, limit)
}

func (h *RoleHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	h.biz.DeleteById(c, uint(id))
}
