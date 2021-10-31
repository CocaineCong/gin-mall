package model

import "github.com/jinzhu/gorm"

type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
