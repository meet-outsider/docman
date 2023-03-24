package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type permissionHandler struct {
	svc biz.IPermissionBiz
}

func NewPermissionHandler(svc biz.IPermissionBiz) *permissionHandler {
	return &permissionHandler{svc}
}

func (h *permissionHandler) Create(c *gin.Context) {
	var perm data.Permission
	if err := c.ShouldBindJSON(&perm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedperm, err := h.svc.Save(&perm)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, savedperm)
}

func (h *permissionHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid perm ID"})
		return
	}

	perm, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "perm not found"})
		return
	}

	c.JSON(http.StatusOK, perm)
}

func (h *permissionHandler) GetByName(c *gin.Context) {
	name := c.Param("name")

	perm, err := h.svc.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "perm not found"})
		return
	}

	c.JSON(http.StatusOK, perm)
}

func (h *permissionHandler) List(c *gin.Context) {
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

	permissions, total, err := h.svc.List(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(permissions, total, pageNum, pageSize))

}

func (h *permissionHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	if err := h.svc.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "permission not found"})
		return
	}

	c.JSON(http.StatusOK, "删除成功！")
}
