package service

import (
	"docman/internal/model"
	"docman/internal/rsp"
	"docman/pkg/database"
	"docman/pkg/global"
	"docman/pkg/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
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
	token := jwt.GenToken(username)
	rsp.Ok(c, "登陆成功，获取token", "token", token)
	return
}
func Registry(c *gin.Context) {
	var userParams model.User
	if err := c.BindJSON(&userParams); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if len(userParams.Username) == 0 || len(userParams.Password) == 0 {
		rsp.Fail(c, "参数校验失败")
		return
	}
	tx := database.Inst.Debug().Where("username = ?", userParams.Username).Find(&model.User{})
	fmt.Println(tx.RowsAffected)
	if tx.RowsAffected > 0 {
		rsp.Fail(c, "用户已存在")
		return
	}
	database.Inst.Create(&userParams)
	rsp.Ok(c, "注册成功")
}
func Info(c *gin.Context) {
	token := c.GetHeader(global.TOKEN)
	userId, exp, err := jwt.ParseToken(token)
	if err != nil {
		rsp.Fail(c, "token解析错误")
		return
	}
	rsp.Ok(c, "获取userId成功", "userId", userId, "exp:", exp)
	return
}
func Logout(c *gin.Context) {

}
