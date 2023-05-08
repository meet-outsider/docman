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

func (s *RoleHandler) Create(ctx *gin.Context) {
	var role data.Role
	ok := kit.BindJson(ctx, &role)
	if !ok {
		return
	}
	s.biz.Save(ctx, &role)
}

func (s *RoleHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	s.biz.GetByID(ctx, uint(id))
}

func (s *RoleHandler) GetByName(ctx *gin.Context) {
	s.biz.GetByName(ctx, ctx.Param("name"))
}

func (s *RoleHandler) List(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	s.biz.List(ctx, page, limit)
}

func (s *RoleHandler) DeleteByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}
	s.biz.DeleteById(ctx, uint(id))
}
