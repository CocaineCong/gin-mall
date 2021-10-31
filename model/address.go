package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	UserID  uint `gorm:"ForeignKey:User; AssociationForeignKey:User"`
	Name    string
	Phone   string
	Address string
}
