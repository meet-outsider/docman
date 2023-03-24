package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type roleHandler struct {
	biz biz.IRoleBiz
}

func NewRoleHandler(biz biz.IRoleBiz) *roleHandler {
	return &roleHandler{biz}
}

func (h *roleHandler) Create(c *gin.Context) {
	var role data.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedRole, err := h.biz.Save(&role)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, savedRole)
}

func (h *roleHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.biz.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *roleHandler) GetByName(c *gin.Context) {
	name := c.Param("name")

	role, err := h.biz.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *roleHandler) List(c *gin.Context) {
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

	roles, total, err := h.biz.List(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(roles, total, pageNum, pageSize))

}

func (h *roleHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.biz.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, "删除成功！")
}
