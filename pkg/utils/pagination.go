package utils

import "math"

type Pagination struct {
	Records   interface{} `json:"records"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Total     int64       `json:"total"`
	TotalPage int         `json:"total_page"`
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
