package model

import (
	"github.com/jinzhu/gorm"
)

//Order 订单信息
type Order struct {
	gorm.Model
	UserID       uint
	ProductID    uint
	BossID		 uint
	AddressID    uint
	Num          uint
	OrderNum     uint64
	Type         uint
	Money 		 int
}

