package service

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"io/ioutil"
	"mall/cache"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
	"os"
)



type OrderPay struct {
	MerchantNum  int    `form:"merchantNum" json:"merchantNum"`
	Money        int    `form:"money" json:"money"`
	OrderNo      string `form:"orderNo" json:"orderNo"`
	ProductID    int    `form:"product_id" json:"product_id"`
	PayTime      string `form:"payTime" json:"payTime" `
	Sign         string `form:"sign" json:"sign" `
	BuyerID      int    `form:"user_id" json:"user_id"`
	BuyerName    string `form:"buyer_name" json:"buyer_name"`
	BossID       int    `form:"boss_id" json:"boss_id"`
	BossName     string `form:"boss_name" json:"boss_name"`
	ProductMoney int    `form:"product_money" json:"product_money"`
	Num          int    `form:"num" json:"num"`
}


func (service *OrderPay) PayDowm() serializer.Response {
	var order model.Order
	code := e.SUCCESS
	err := model.DB.Where("user_id=? AND product_id=?",service.BuyerID,service.ProductID).Find(&order).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var user model.User
	money := service.Money
	num := service.Num
	money = money*num
	err = model.DB.First(&user, service.BuyerID).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Money =user.Money -money
	err = model.DB.Save(&user).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var userboss model.User
	err = model.DB.First(&userboss, service.BossID).Error
	userboss.Money += money
	err = model.DB.Save(&userboss).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var product model.Product
	model.DB.First(&product).Where("product_id=? AND boss_id=?",service.ProductID,service.BossID)
	product.Num -= num
	err = model.DB.Save(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Find(&order).Where("boss_id=? AND user_id=? AND product_id=?",service.BossID,service.BuyerID,service.ProductID).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var productTemp model.Product
	err = model.DB.Where("product_id=?",service.ProductID).Find(&productTemp).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var productBoss model.User
	model.DB.Find(&productBoss).Where("id=?",service.BuyerID)
	var productTest model.Product
	productTest.Num=service.Num
	productTest.BossID=service.BuyerID
	productTest.OnSale = false
	productTest.ImgPath = productTemp.ImgPath
	productTest.Price = productTemp.Price
	productTest.Info = productTemp.Info
	productTest.Name = productTemp.Name
	productTest.Title = productTemp.Title
	productTest.CategoryID = productTemp.CategoryID
	productTest.DiscountPrice = productTemp.DiscountPrice
	productTest.BossName = productBoss.UserName
	productTest.BossAvatar = productBoss.Avatar
	err = model.DB.Create(&productTest).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&order).Error
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
		Status:code,
		Msg:e.GetMsg(code),
	}
}