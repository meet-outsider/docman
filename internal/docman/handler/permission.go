package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	biz biz.IPermissionBiz
}

func NewPermissionHandler(biz biz.IPermissionBiz) *PermissionHandler {
	return &PermissionHandler{biz}
}

func (h *PermissionHandler) Create(c *gin.Context) {
	var perm data.Permission
	ok := kit.BindJson(c, &perm)
	if !ok {
		return
	}
	h.biz.Save(c, &perm)
}

func (h *PermissionHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid perm ID"})
		return
	}
	h.biz.GetByID(c, uint(id))
}

func (h *PermissionHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	h.biz.GetByName(c, name)
}

func (h *PermissionHandler) List(c *gin.Context) {
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
	h.biz.List(c, pageNum, pageSize)
}

func (h *PermissionHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}
	h.biz.DeleteById(c, uint(id))
}
