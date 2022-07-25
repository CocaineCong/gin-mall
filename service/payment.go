package service

import (
	"context"
	"errors"
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
	OrderId   uint    `form:"order_id" json:"order_id"`
	Money     float64 `form:"money" json:"money"`
	OrderNo   string  `form:"orderNo" json:"orderNo"`
	ProductID int     `form:"product_id" json:"product_id"`
	PayTime   string  `form:"payTime" json:"payTime" `
	Sign      string  `form:"sign" json:"sign" `
	BossID    int     `form:"boss_id" json:"boss_id"`
	BossName  string  `form:"boss_name" json:"boss_name"`
	Num       int     `form:"num" json:"num"`
	Key       string  `form:"key" json:"key"`
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.SUCCESS
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderById(service.OrderId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对钱进行解密。减去订单。再进行加密。
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	// 金额不足进行回滚
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	// 更新用户金额失败，回滚
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	boss, err := userDao.GetUserById(uint(service.BossID))

	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)

	err = userDao.UpdateUserById(uint(service.BossID), boss)
	// 更新boss金额失败，回滚
	if err != nil {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(service.ProductID))
	product.Num -= num
	// 更新商品数量减少失败，回滚
	err = productDao.UpdateProduct(uint(service.ProductID), product)
	if err != nil {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = orderDao.DeleteOrderById(service.OrderId)
	// 删除订单失败，回滚
	if err != nil {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productUser := &model.Product{
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		Num:           num,
		OnSale:        false,
		BossID:        int(uId),
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}
	// 买完商品后创建成了自己的商品失败。订单失败，回滚
	err = productDao.CreateProduct(productUser)
	if err != nil {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
