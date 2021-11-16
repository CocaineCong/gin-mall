package model

import (
	"github.com/jinzhu/gorm"
)

//商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint `gorm:"ForeignKey:CategoryID"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale 		  bool `gorm:"default:false"`
	Num 		  int
	BossID        int
	BossName      string
	BossAvatar    string
}
