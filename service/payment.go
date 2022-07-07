package service

import (
	"fmt"
	logging "github.com/sirupsen/logrus"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/serializer"
	"strconv"
)

type OrderPay struct {
	Money     float64 `form:"money" json:"money"`
	OrderNo   string  `form:"orderNo" json:"orderNo"`
	ProductID int     `form:"product_id" json:"product_id"`
	PayTime   string  `form:"payTime" json:"payTime" `
	Sign      string  `form:"sign" json:"sign" `
	BossID    int     `form:"boss_id" json:"boss_id"`
	BossName  string  `form:"boss_name" json:"boss_name"`
	Num       int     `form:"num" json:"num"`
	Key       string  `json:"key" form:"key"`
}

func (service *OrderPay) PayDown(id uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	var order model.Order
	code := e.SUCCESS
	err := dao.DB.Where("user_id=? AND product_id=?", id, service.ProductID).Find(&order).Error
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
	numFloat := float64(num)
	money = money * numFloat
	err = dao.DB.First(&user, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)
	err = dao.DB.Save(&user).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var boss model.User
	err = dao.DB.First(&boss, service.BossID).Error
	moneyStr = util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	err = dao.DB.Save(&boss).Error
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
	dao.DB.First(&product).Where("product_id=? AND boss_id=?", service.ProductID, service.BossID)
	product.Num -= num
	err = dao.DB.Save(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = dao.DB.Find(&order).Where("boss_id=? AND user_id=? AND product_id=?", service.BossID, id, service.ProductID).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = dao.DB.Delete(&order).Error
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
	err = dao.DB.Where("id=?", service.ProductID).Find(&productTemp).Error
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
	dao.DB.Find(&productBoss).Where("id=?", id)
	var productTest model.Product
	productTest.Num = service.Num
	productTest.BossID = int(id)
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
	err = dao.DB.Create(&productTest).Error
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
