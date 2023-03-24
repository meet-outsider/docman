package data

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name  string `validate:"required"`
	Path  string `validate:"required"`
	IsDir bool   `validate:"required"`
}
