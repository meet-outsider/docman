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

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	h.biz.GetByID(c, uint(id))
}

func (h *UserHandler) GetByUsername(c *gin.Context) {
	h.biz.GetByUsername(c, c.Param("username"))
}

func (h *UserHandler) List(c *gin.Context) {
	page, limit, err := kit.GetPage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.biz.List(c, page, limit)
}

// Update 更新用户
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	var param data.UserInput
	param.ID = uint(id)
	ok := kit.UnmarshalJSON(c, &param)
	if !ok {
		return
	}
	h.biz.Update(c, &param.User)
}

func (h *UserHandler) DeleteByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	h.biz.DeleteByID(context, uint(id))
}

func (h *UserHandler) Save(c *gin.Context) {
	var param data.UserInput
	// 参数校验
	ok := kit.BindJson(c, &param)
	if !ok {
		return
	}
	if len(param.Roles) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roles is required"})
		return
	}
	h.biz.Save(c, &param)
}

func (h *UserHandler) DeleteByIDs(c *gin.Context) {
	var ids model.IDs
	ok := kit.BindJson(c, &ids)
	if !ok {
		return
	}
	h.biz.DeleteByIDs(c, ids.IDs)
}
