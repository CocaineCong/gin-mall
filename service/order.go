package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"mall/cache"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

//Create
type OrderService struct {
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressID uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	UserID    uint `form:"user_id" json:"user_id"`
	OrderNum  uint `form:"order_num" json:"order_num"`
	PageNum   int  `form:"pageNum"`
	PageSize  int  `form:"pageSize"`
	Type      int  `form:"type" json:"type"`
}

func (service *OrderService) Create(id uint) serializer.Response {
	var product model.Product
	dao.DB.First(&product, service.ProductID)
	order := model.Order{
		UserID:    id,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1,
	}
	address := model.Address{}
	code := e.SUCCESS
	if err := dao.DB.First(&address, service.AddressID).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressID = address.ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(id))
	number = number + productNum + userNum
	orderNum, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		logging.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.OrderNum = orderNum
	err = dao.DB.Create(&order).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//订单号存入Redis中，设置过期时间
	data := redis.Z{
		Score:  float64(time.Now().Unix()) + 15*time.Minute.Seconds(),
		Member: orderNum,
	}
	cache.RedisClient.ZAdd("3", data)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) List(id uint) serializer.Response {
	var orders []model.Order
	var total int64
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 5
	}
	if service.Type == 0 {
		if err := dao.DB.Model(&orders).Where("user_id=?", id).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		if err := dao.DB.Where("user_id = ?", id).Offset((service.PageNum - 1) * service.PageSize).
			Limit(service.PageSize).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := dao.DB.Model(&orders).Where("user_id=? AND type = ?", id, service.Type).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		if err := dao.DB.Where("user_id=? AND type=?", id, service.Type).
			Offset((service.PageNum - 1) * service.PageSize).
			Limit(service.PageSize).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total))
}

func (service *OrderService) Show(id string) serializer.Response {
	var order model.Order
	var product model.Product
	var address model.Address
	code := e.SUCCESS
	if err := dao.DB.Model(&model.Order{}).First(&order, id).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	dao.DB.Where("id = ?", order.AddressID).First(&address)
	if err := dao.DB.Where("id=?", order.ProductID).First(&product).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ErrorNotExistProduct
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) Delete(id string) serializer.Response {
	var order model.Order
	code := e.SUCCESS
	err := dao.DB.First(&order, id).Error
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
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
