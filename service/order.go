package service

import (
	"context"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"mall/cache"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/serializer"
	"strconv"
	"time"
)

const OrderTimeOutKey = "order_timeout_queue"

type OrderService struct {
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressID uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	UserID    uint `form:"user_id" json:"user_id"`
	OrderNum  uint `form:"order_num" json:"order_num"`
	Type      int  `form:"type" json:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, id uint) serializer.Response {
	code := e.SUCCESS

	order := &model.Order{
		UserID:    id,
		ProductID: service.ProductID,
		BossID:    service.BossID,
		Num:       int(service.Num),
		Money:     float64(service.Money),
		Type:      1,
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressID)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	order.AddressID = address.ID
	number, _ := util.SnowflakeID() //改用snowflake生成uuid
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(id))
	orderNum, _ := strconv.ParseUint(number+productNum+userNum, 10, 64)
	order.OrderNum = orderNum

	orderDao := dao.NewOrderDao(ctx)
	err = orderDao.CreateOrder(order)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//订单号存入Redis中，设置过期时间   设置分数是当前时间加上15分钟
	data := redis.Z{
		Score:  float64(time.Now().Unix()) + 15*time.Minute.Seconds(),
		Member: orderNum,
	}
	if err = cache.RedisClient.ZAdd(OrderTimeOutKey, data).Err(); err != nil {
		util.LogrusObj.Infoln("订单【%d】加入延迟队列失败, err: %v", orderNum, err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	var orders []*model.Order
	var total int64
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 5
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	condition["user_id"] = uId

	if service.Type == 0 {
		condition["type"] = 0
	} else {
		condition["type"] = service.Type
	}
	orders, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(total))
}

func (service *OrderService) Show(ctx context.Context, uId string) serializer.Response {
	code := e.SUCCESS

	orderId, _ := strconv.Atoi(uId)
	orderDao := dao.NewOrderDao(ctx)
	order, _ := orderDao.GetOrderById(uint(orderId))

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(order.AddressID)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductID)
	if err != nil {
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

func (service *OrderService) Delete(ctx context.Context, oId string) serializer.Response {
	code := e.SUCCESS

	orderDao := dao.NewOrderDao(ctx)
	orderId, _ := strconv.Atoi(oId)
	// 删除订单实际上是修改订单状态为关闭
	err := orderDao.CloseOrderById(uint(orderId))
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
