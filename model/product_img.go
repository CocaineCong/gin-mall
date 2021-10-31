package model

import "github.com/jinzhu/gorm"

type ProductImg struct {
	gorm.Model
	ProductID  uint
	ImgPath    string
	BossID     int
	BossName   string
	BossAvatar string
}
