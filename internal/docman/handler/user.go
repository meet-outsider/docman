// @description 用户接口，包括用户的增删改查，用户登录注册等接口
// @author outsider
// @date 2023-04-01
// @updated 2023-05-02
package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"docman/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	biz biz.IUserBiz
}
type User struct {
	data.User
	Roles []uint `binding:"required"`
}

func NewUserHandler(biz biz.IUserBiz) *UserHandler {
	return &UserHandler{biz}
}

func (s *UserHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	s.biz.GetByID(ctx, uint(id))
}

func (s *UserHandler) GetByUsername(ctx *gin.Context) {
	s.biz.GetByUsername(ctx, ctx.Param("username"))
}

func (s *UserHandler) ListByUsername(ctx *gin.Context) {
	s.biz.ListByUsername(ctx, ctx.Param("username"))
}

func (s *UserHandler) List(ctx *gin.Context) {
	page, limit, err := kit.GetPage(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.biz.List(ctx, page, limit)
}

// Update 更新用户
func (s *UserHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	var param data.UserInput
	param.ID = uint(id)
	ok := kit.UnmarshalJSON(ctx, &param)
	if !ok {
		return
	}
	s.biz.Update(ctx, &param.User)
}

func (s *UserHandler) DeleteByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	s.biz.DeleteByID(context, uint(id))
}

func (s *UserHandler) Save(ctx *gin.Context) {
	var param data.UserInput
	// parameter check
	ok := kit.BindJson(ctx, &param)
	if !ok {
		return
	}
	if len(param.Roles) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "roles is required"})
		return
	}
	s.biz.Save(ctx, &param)
}

func (s *UserHandler) DeleteByIDs(ctx *gin.Context) {
	var ids model.IDs
	ok := kit.BindJson(ctx, &ids)
	if !ok {
		return
	}
	s.biz.DeleteByIDs(ctx, ids.IDs)
}
