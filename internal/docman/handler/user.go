package handler

import (
	"docman/internal/docman/biz"
	"docman/internal/docman/data"
	"docman/pkg/database"
	"docman/pkg/global"
	"docman/pkg/kit"
	"docman/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type userHandler struct {
	biz biz.IUserBiz
}
type User struct {
	data.User
	Roles []uint `binding:"required"`
}

func NewUserHandler(biz biz.IUserBiz) *userHandler {
	return &userHandler{biz}
}

// Login 用户登录
func (h *userHandler) Login(c *gin.Context) {
	var params map[string]string
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": kit.Translate(err)})
		return
	}
	username := params["username"]
	password := params["password"]
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不能为空"})
		return
	}
	user, err := h.biz.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	// 检查密码是否正确
	if err := kit.Decrypt(password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}
	roles := make([]string, len(user.Roles))
	fmt.Println("userRoles ", user.Roles)
	fmt.Println(len(user.Roles))
	for i := range user.Roles {
		roles[i] = user.Roles[i].Name
	}
	fmt.Println("roles ", roles)

	userInfo := model.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Roles:    roles,
	}
	fmt.Println("userInfo ", userInfo)
	token := kit.GenToken(userInfo)
	c.JSON(http.StatusOK, gin.H{"token": token})
	return
}

// Registry 用户注册
func (h *userHandler) Registry(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": kit.Translate(err)})
		return
	}
	save := user.User
	err := h.biz.Save(&save)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	roleIDs := user.Roles
	// 保存用户角色关系
	tx := database.Inst.Begin()
	if len(roleIDs) != 0 {
		for _, id := range roleIDs {
			_, err := h.biz.GetByID(id)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
				return
			}
			e := database.Inst.Save(data.UserRole{
				UserID: save.ID,
				RoleID: id,
			}).Error
			if e != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "保存用户角色关系失败"})
				return
			}
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

func (h *userHandler) Info(c *gin.Context) {
	token := c.GetHeader(global.TOKEN)
	subject, _, err := kit.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
		return
	}
	userId := subject.ID
	user, err := h.biz.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return

}

func (h *userHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	user, err := h.biz.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *userHandler) GetByUsername(c *gin.Context) {
	username := c.Query("username")

	user, err := h.biz.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *userHandler) List(c *gin.Context) {
	page, limit, err2 := kit.GetPage(c)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}
	users, count, err := h.biz.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(users, count, page, limit))
}

// Update 更新用户
func (h *userHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	var param User
	if err := c.ShouldBindJSON(&param); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": kit.Translate(err)})
		//return
	}
	user := param.User
	user.ID = uint(id)
	err = h.biz.Update(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "更新成功"})
}

func (h *userHandler) DeleteByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	err = h.biz.DeleteByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
}
