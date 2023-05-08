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

func (s *PermissionHandler) Create(ctx *gin.Context) {
	var perm data.Permission
	ok := kit.BindJson(ctx, &perm)
	if !ok {
		return
	}
	s.biz.Save(ctx, &perm)
}

func (s *PermissionHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid perm ID"})
		return
	}
	s.biz.GetByID(c, uint(id))
}

func (s *PermissionHandler) GetByName(ctx *gin.Context) {
	name := ctx.Param("name")
	s.biz.GetByName(ctx, name)
}

func (s *PermissionHandler) List(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	s.biz.List(ctx, pageNum, pageSize)
}

func (s *PermissionHandler) DeleteByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}
	s.biz.DeleteById(ctx, uint(id))
}
