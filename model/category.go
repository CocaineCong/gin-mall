package model

import "github.com/jinzhu/gorm"

// Category 分类模型
type Category struct {
	gorm.Model
	CategoryID   uint
	CategoryName string
	// Num          uint
}
