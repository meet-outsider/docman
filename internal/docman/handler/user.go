package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/kit"
	"fmt"
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
	fmt.Println("getByid ")
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("getByid ", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	h.biz.GetByID(c, uint(id))
}

func (h *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	h.biz.GetByUsername(c, username)
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
	fmt.Println(id)
	var param User
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
