package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryName string
}
