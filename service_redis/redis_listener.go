package service_redis

import (
	"context"
	"github.com/go-redis/redis"
	"mall/cache"
	"mall/dao"
	util "mall/pkg/utils"
	"mall/service"
	"strconv"
	"time"
)

func OrderTimeOutListener() { //todo: 订单超时队列里有普通商品下单和 秒杀商品下单 最好区分一下
	// 在循环外创建mysql连接 减少资源消耗
	orderDao := dao.NewOrderDao(context.Background())
	// 协程启动
	go func() {
		for {
			opt := redis.ZRangeBy{
				Min:    strconv.Itoa(0),
				Max:    strconv.Itoa(int(time.Now().Unix())),
				Offset: 0,
				Count:  1, // 一次返回多少数据
			}

			// 获取过期时间截止至当前时间段内的订单
			orderList, err := cache.RedisClient.ZRangeByScore(service.OrderTimeOutKey, opt).Result()
			if err != nil {
				util.LogrusObj.Info("redis错误 err:", err)
			}

			// 如果没检测到过期的订单就间隔100毫秒后再检测
			if len(orderList) == 0 {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			// 根据订单号获取订单信息
			orderNum, _ := strconv.ParseUint(orderList[0], 10, 64)
			orderInfo, _ := orderDao.GetOrderByOrderNum(uint(orderNum))

			// 校验订单状态
			if orderInfo.Type != 1 {
				util.LogrusObj.Info("订单已经不是未支付状态 关闭订单失败")
				return
			}

			// mysql关闭订单
			if err = orderDao.CloseOrderById(orderInfo.ID); err != nil {
				util.LogrusObj.Info("mysql关闭订单失败")
				return
			}

			// redis 恢复库存 todo: 订单超时队列里有普通商品下单和 秒杀商品下单 最好区分一下
			//cache.RedisClient.HIncrBy("SK"+strconv.Itoa(int(orderInfo.SkillGoodId)), "num", 1)

			// 删除订单下单缓存
			if err = service.DeleteOrderCache(int(orderInfo.ProductID), int(orderInfo.UserID)); err != nil {
				util.LogrusObj.Info("redis删除订单下单缓存错误 err:", err)
				return
			}

			// redis 删除超时队列的订单缓存
			if err := cache.RedisClient.ZRem(service.OrderTimeOutKey, orderList).Err(); err != nil {
				util.LogrusObj.Info("redis删除超时队列的订单缓存错误 err:", err)
			}

		}

	}()
}
