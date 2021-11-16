package model

import (
	"github.com/jinzhu/gorm"
)

//Order 订单信息
type Order struct {
	gorm.Model
	UserID       uint `gorm:"ForeignKey:UserID"`
	ProductID    uint `gorm:"ForeignKey:ProductID"`
	BossID		 uint `gorm:"ForeignKey:BossID"`
	AddressID    uint `gorm:"ForeignKey:AddressID"`
	Num          uint
	OrderNum     uint64
	Type         uint
	Money 		 int
}
