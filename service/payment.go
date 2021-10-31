package service

import (
	"FanOneMall/conf"
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//初始化支付的服务
type InitPayService struct {
	OrderNum string `form:"order_num" json:"order_num"`
	PayType string `form:"pay_type" json:"pay_type"`
	Amount string `form:"amount" json:"amount"`
}

// PayOrderInfo
type PayOrderInfo struct {
	ID string `json:"id"`			//渠道唯一ID
	PayURL string `json:"payUrl"`	//支付页URL
}

type Result struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data PayOrderInfo `json:"data"`
}

//接受FM支付回调接口
type ConfirmPayService struct {
	MerchantNum 	string `form:"merchantNum" json:"merchantNum"`
	OrderNo 		string `form:"orderNo" json:"orderNo"`
	PlatformOrderNo string `form:"platformOrderNo" json:"platformOrderNo"`
	Amount          string `form:"amount" json:"amount" `
	ActualPayAmount string `form:"actualPayAmount" json:"actualPayAmount" `
	State           int    `form:"state" json:"state" `
	Attch           string `form:"attch" json:"attch" `
	PayTime         string `form:"payTime" json:"payTime" `
	Sign            string `form:"sign" json:"sign" `
}

type OrderPay struct {
	MerchantNum 	int `form:"merchantNum" json:"merchantNum"`
	Money 			int `form:"money" json:"money"`
	OrderNo 		string `form:"orderNo" json:"orderNo"`
	PlatformOrderNo string `form:"platformOrderNo" json:"platformOrderNo"`
	Amount          int `form:"amount" json:"amount" `
	ActualPayAmount string `form:"actualPayAmount" json:"actualPayAmount" `
	State           int    `form:"state" json:"state" `
	Attch           string `form:"attch" json:"attch" `
	ProductID 		int `form:"product_id" json:"product_id"`
	PayTime         string `form:"payTime" json:"payTime" `
	Sign            string `form:"sign" json:"sign" `
	BuyerID 		int    `form:"user_id" json:"user_id"`
	BuyerName 		string `form:"buyer_name" json:"buyer_name"`
	BossID 			int    `form:"boss_id" json:"boss_id"`
	BossName		string `form:"boss_name" json:"boss_name"`
	ProductMoney	int    `form:"product_money" json:"product_money"`
	Num 			int    `form:"num" json:"num"`
}

func (service *InitPayService) Init() serializer.Response{
	code := e.SUCCESS
	var buff bytes.Buffer
	buff.WriteString("1361954603746197504")
	buff.WriteString(service.OrderNum)
	buff.WriteString(service.Amount)
	buff.WriteString("http://localhost:3000/api/v1/payments")
	buff.WriteString("9aeca5434094c80b67db4d1907089af3")
	fmt.Println("Buff Test")
	signTemp :=fmt.Sprintf("%s",buff)
	fmt.Println(signTemp)

	sign:=fmt.Sprintf("%x",md5.Sum(buff.Bytes()))
	fmt.Println("sign",sign)
	returnURL := "http://localhost:8080/home#/"
	//构造请求参数
	buff.Reset()
	buff.WriteString("sign=")
	buff.WriteString(sign)
	buff.WriteString("&amount=")
	buff.WriteString(service.Amount)
	buff.WriteString("&orderNo=")
	buff.WriteString(service.OrderNum)
	buff.WriteString("&payType=")
	buff.WriteString(service.PayType)
	buff.WriteString("&merchantNum=")
	buff.WriteString("1361954603746197504")
	buff.WriteString("&notifyUrl=")
	buff.WriteString("http://localhost:3000/api/v1/payments")
	buff.WriteString("&returnUrl=")
	buff.WriteString(returnURL)
	buff.WriteString("&attch=")
	buff.WriteString("alipay")
	//fmt.Println("PayTestBuff")
	//fmt.Println(&buff)
	//调用渠道接口
	resp, err := http.Post("http://api-15pl2fs7g0000.zhifu.fm.it88168.com/api/startOrder", "application/x-www-form-urlencoded", &buff)
	if err != nil {
		logging.Info(err)
		code = e.ErrorCallApi
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//读取所有响应数据
	var data []byte
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		logging.Info(err)
		code = e.ErrorReadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//解析渠道返回的JSON
	var result Result
	if err = json.Unmarshal(data, &result); err != nil {
		logging.Info(err)
		code = e.ErrorUnmarshalJson
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: result.Code,
		Msg:    result.Msg,
		Data:   result.Data,
	}
}

// 接收FM支付回调 详情请查阅FM支付文档
func (service *ConfirmPayService) Confirm() {
	if service.Attch == "alipay" {
		if service.State == 1 {
			if err := model.DB.Model(model.Order{}).Where("order_num=?", service.OrderNo).Update("type", 2).Error; err != nil {
				logging.Info(err)
			}
			if err := conf.RedisClient.ZRem(os.Getenv("REDIS_ZSET_KEY"), service.OrderNo).Err(); err != nil {
				logging.Info(err)
			}
		}
	}
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
	user.Monery =user.Monery-money
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
	userboss.Monery += money
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
	productTest.OnSale = "1"
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