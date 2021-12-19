package serializer

import (
	"mall/conf"
	"mall/model"
)

type Money struct {
	UserID uint `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}


func BuildMoney(item model.User,key string) Money {
	conf.Encryption.SetKey(key)
	return Money{
		UserID: item.ID,
		UserName:   item.UserName,
		UserMoney:   conf.Encryption.AesDecoding(item.Money),
	}
}

