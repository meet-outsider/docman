package kit

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	Records   interface{} `json:"records"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Total     int64       `json:"total"`
	TotalPage int         `json:"total_page"`
}

func PostPage(ctx *gin.Context) (page int, limit int, err error) {
	page, err = strconv.Atoi(ctx.DefaultPostForm("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	limit, err = strconv.Atoi(ctx.DefaultPostForm("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page limit"})
		return
	}
	return
}
func GetPage(ctx *gin.Context) (page int, limit int, err error) {
	page, err = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page limit"})
		return
	}
	return
}
func BuildPagination(records interface{}, total int64, page, limit int) *Pagination {
	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	return &Pagination{
		Records:   records,
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPage,
	}
}
