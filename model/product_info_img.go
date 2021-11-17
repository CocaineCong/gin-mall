package model

import "github.com/jinzhu/gorm"

//商品图片模型
type ProductInfoImg struct {
	gorm.Model
	Product Product `gorm:"ForeignKey:ProductID"`
	ProductID  uint `gorm:"not null"`
	ImgPath   string
}
