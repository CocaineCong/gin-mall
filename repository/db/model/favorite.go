package model

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"`
}
