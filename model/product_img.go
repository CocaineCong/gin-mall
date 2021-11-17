package model

import "github.com/jinzhu/gorm"

type ProductImg struct {
	gorm.Model
	Product Product `gorm:"ForeignKey:ProductID"`
	ProductID  uint `gorm:"not null"`
	ImgPath    string
}
