package model

type Pagination struct {
	Records   any
	Page      int
	Limit     int
	Total     int64
	TotalPage int
}
