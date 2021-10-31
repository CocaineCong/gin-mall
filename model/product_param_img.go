package model

import "github.com/jinzhu/gorm"

type ProductParamImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
