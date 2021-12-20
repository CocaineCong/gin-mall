package service

import (
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
)

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *ShowMoneyService) Show(id uint) serializer.Response {
	var user model.User
	code := e.SUCCESS
	model.DB.Model(model.User{}).Where("id=?", id).First(&user)
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
