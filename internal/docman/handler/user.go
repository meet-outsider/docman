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

func (s *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	s.biz.GetByID(c, uint(id))
}

func (s *UserHandler) GetByUsername(c *gin.Context) {
	s.biz.GetByUsername(c, c.Param("username"))
}

func (s *UserHandler) ListByUsername(c *gin.Context) {
	s.biz.ListByUsername(c, c.Param("username"))
}

func (s *UserHandler) List(c *gin.Context) {
	page, limit, err := kit.GetPage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.biz.List(c, page, limit)
}

// Update 更新用户
func (s *UserHandler) Update(c *gin.Context) {
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
	s.biz.Update(c, &param.User)
}

func (s *UserHandler) DeleteByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	s.biz.DeleteByID(context, uint(id))
}

func (s *UserHandler) Save(c *gin.Context) {
	var param data.UserInput
	// parameter check
	ok := kit.BindJson(c, &param)
	if !ok {
		return
	}
	if len(param.Roles) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roles is required"})
		return
	}
	s.biz.Save(c, &param)
}

func (s *UserHandler) DeleteByIDs(c *gin.Context) {
	var ids model.IDs
	ok := kit.BindJson(c, &ids)
	if !ok {
		return
	}
	s.biz.DeleteByIDs(c, ids.IDs)
}
