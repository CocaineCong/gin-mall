package service

import (
	logging "github.com/sirupsen/logrus"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}


func (service *AddressService) Create(id uint) serializer.Response {
	var address model.Address
	code := e.SUCCESS
	address = model.Address{
		UserID:  id,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := model.DB.Create(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var addresses []model.Address
	err = model.DB.Model(model.Address{}).Where("user_id = ?", id).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Show(id string) serializer.Response {
	var addresses []model.Address
	code := e.SUCCESS
	err := model.DB.Where("user_id = ?", id).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Delete(id string) serializer.Response {
	var address model.Address
	code := e.SUCCESS
	err := model.DB.Where("id = ?", id).First(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Update(uid uint,aid string) serializer.Response {
	code := e.SUCCESS
	address := model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	aidInt,_ := strconv.Atoi(aid)
	address.ID = uint(aidInt)
	err := model.DB.Save(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var addresses []model.Address
	err = model.DB.Model(model.Address{}).Where("user_id = ?", uid).
		Order("created_at desc").Find(&addresses).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}
