package service_mq

import (
	"context"
	"encoding/json"
	"errors"
	"mall/cache"
	"mall/dao"
	"mall/model"
	util "mall/pkg/utils"
	"mall/service"
	"strconv"
)

//type SkillGoodMessage struct {
//	*model.SkillGood2MQ
//	OrderNum string //分布式唯一订单号
//}

func MQ2SkillGoodConsumer() {
	// 通过mq连接创建channel
	ch, err := model.MQ.Channel()
	// 获取同一个队列名的队列
	q, err := ch.QueueDeclare("skill_goods", true, false, false, false, nil)
	// 设置在收到ack确认之前 控制分配给consumer的消息个数
	err = ch.Qos(1, 0, false)
	//从队列取出元素
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)

	if err != nil {
		err = errors.New("rabbitMQ err:" + err.Error())
		// MQ消费端启动失败就panic
		util.LogrusObj.Panicln(err)
	}
	// 创建mysql连接
	orderDao := dao.NewOrderDao(context.Background())

	// 启动协程 开始消费mq队列
	go func() {
		for d := range msgs {
			var p *service.SkillGoodMessage
			_ = json.Unmarshal(d.Body, p)
			// 获取Uuid订单号
			OrderNum, _ := strconv.ParseUint(p.OrderNum, 10, 64)

			// 创建订单
			newOrder := &model.Order{
				UserID:    p.UserId,
				ProductID: p.ProductId,
				BossID:    p.BossId,
				AddressID: p.AddressId,
				Num:       1,
				OrderNum:  OrderNum,
				Type:      1, //未支付
				Money:     p.Money,
			}

			errOrder := service.CreateSkillGoodOrder(orderDao.DB, newOrder)
			if errOrder != nil {
				return
			}
			// 释放锁
			pid := strconv.Itoa(int(p.ProductId))
			_, err = cache.RedisClient.Del(service.SkillGoodsLockKey + pid).Result()
			if err != nil {
				util.LogrusObj.Infoln("unlock fail")
			}

			/*	paymentObj := &OrderPay{   todo：支付 感觉支付和下单应该分开 后续添加第三方支付
					OrderId:   order.ID,
					Money:     order.Money * float64(order.Num),
					ProductID: order.ProductID,
					BossID:    order.BossID,
					Num:       order.Num,
					Key:       p.Key,
				}
				fmt.Println("payment", *paymentObj)
				errPay := paymentObj.PayDown(context.Background(), p.UserId)
				if errPay != nil {
					return
				}*/

			// redis扣除库存
			cache.RedisClient.HIncrBy("SK"+strconv.Itoa(int(p.SkillGoodId)), "num", -1) // 库存数量-1
			// 返回ack确认
			_ = d.Ack(false) // 确认消息,必须为false

		}

	}()

}
