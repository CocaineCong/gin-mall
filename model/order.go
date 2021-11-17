package model

import (
	"github.com/jinzhu/gorm"
)

//Order 订单信息
type Order struct {
	gorm.Model
	User 		 User 	 	`gorm:"ForeignKey:UserID"`
	UserID       uint 		`gorm:"not null"`
	Product    	 Product 	`gorm:"ForeignKey:ProductID"`
	ProductID    uint 		`gorm:"not null"`
	Boss		 User 		`gorm:"ForeignKey:BossID"`
	BossID		 uint 		`gorm:"not null"`
	Address 	 Address 	`gorm:"ForeignKey:AddressID"`
	AddressID    uint 		`gorm:"not null"`
	Num          uint
	OrderNum     uint64
	Type         uint
	Money 		 int
}
