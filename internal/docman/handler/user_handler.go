package handler

import (
	"docman/internal/docman/model"
	"docman/internal/docman/service"
	"docman/internal/pkg/rsp"
	"docman/pkg/global"
	"docman/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *userHandler {
	return &userHandler{svc}
}

func (h *userHandler) Login(c *gin.Context) {
	var params map[string]string
	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	username := params["username"]
	password := params["password"]
	if len(username) == 0 || len(password) == 0 {
		rsp.Fail(c, "参数校验失败")
		return
	}
	token := utils.GenToken(username)
	rsp.Ok(c, "登陆成功，获取token", "token", token)
	return
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.svc.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rsp.Created(c, "created")
}

func (h *userHandler) GetUserInfo(c *gin.Context) {
	token := c.GetHeader(global.TOKEN)
	userId, exp, err := utils.ParseToken(token)
	if err != nil {
		rsp.Fail(c, "token解析错误")
		return
	}
	rsp.Ok(c, "获取userId成功", "username", userId, "exp:", exp)
	return
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rsp.Fail(c, "Invalid user ID")
		return
	}

	user, err := h.svc.GetUserByID(uint(id))
	if err != nil {
		rsp.Ok(c, "User not found")
		return
	}
	rsp.Ok(c, "ok", "user", user)
}

func (h *userHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.svc.GetUserByUsername(username)
	if err != nil {
		rsp.Ok(c, "User not found")
		return
	}
	rsp.Ok(c, "ok", "user", user)
}

func (h *userHandler) GetUsers(c *gin.Context) {
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

	users, count, err := h.svc.GetUsers(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rsp.Ok(c, "ok", utils.BuildPagination(users, count, pageNum, pageSize))
}
