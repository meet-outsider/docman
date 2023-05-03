// @description: 用户业务逻辑
// @author: yuanzichao
// @date: 2021-06-11 16:56:00
// @version V1
package biz

import (
	"docman/internal/docman/data"
	"docman/internal/docman/repo"
	"docman/pkg/database"
	"docman/pkg/kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IUserBiz 用户业务逻辑接口
type IUserBiz interface {
	Save(c *gin.Context, user *data.UserInput)
	GetByID(c *gin.Context, id uint)
	GetByUsername(c *gin.Context, username string)
	ListByUsername(c *gin.Context, username string)
	List(c *gin.Context, pageNum int, pageSize int)
	Update(c *gin.Context, user *data.User)
	DeleteByID(c *gin.Context, id uint)
	DeleteByIDs(c *gin.Context, ids []uint)
}

// userBiz 用户业务逻辑实现
type userBiz struct {
	repo repo.IUserRepo
}

func (*userBiz) ListByUsername(c *gin.Context, username string) {
	users := make([]data.User, 0)
	database.Inst.Where("username like ?", "%"+username+"%").Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func NewUserBiz(repo repo.IUserRepo) IUserBiz {
	return &userBiz{repo}
}

func (s *userBiz) Save(c *gin.Context, userInput *data.UserInput) {
	var user = userInput.User
	isExist, _ := s.repo.GetByUsername(user.Username)
	if isExist != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}
	// password encrypt
	user.Password = kit.Encrypt(user.Password)
	// save user
	if err := s.repo.Save(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// save relation of user and role
	var userRoles []data.UserRole
	for _, roleID := range userInput.Roles {
		userRoles = append(userRoles, data.UserRole{
			UserID: user.ID,
			RoleID: roleID,
		})
	}
	if err := database.Inst.Model(&data.UserRole{}).Save(&userRoles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *userBiz) GetByID(c *gin.Context, id uint) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *userBiz) GetByUsername(c *gin.Context, username string) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *userBiz) List(c *gin.Context, page int, limit int) {
	var request struct {
		User data.User `json:"user"`
	}
	user := &request.User
	ok := kit.UnmarshalJSON(c, &request)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	list, total, err := s.repo.List(page, limit, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kit.BuildPagination(list, total, page, limit))
}

func (s *userBiz) Update(c *gin.Context, user *data.User) {
	err := s.repo.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *userBiz) DeleteByID(c *gin.Context, u uint) {
	err := s.repo.DeleteByID(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func (s *userBiz) DeleteByIDs(c *gin.Context, ids []uint) {

	err := s.repo.DeleteByIDs(ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
