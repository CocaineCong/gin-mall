package model

import "github.com/jinzhu/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint `gorm:"not null"`
}
