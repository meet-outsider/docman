package data

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string `json:"name"`
	Path  string `json:"path"`
	IsDir bool   `json:"isDir"`
}
