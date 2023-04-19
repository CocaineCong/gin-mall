package main

import (
	"mall/conf"
	"mall/routes"
	"mall/service_mq"
	"mall/service_redis"
)

func main() {
	// Ek1+Ep1==Ek2+Ep2
	conf.Init()

	// redis启动监听 订单超时队列
	service_redis.OrderTimeOutListener()

	// mq启动监听 秒杀订单队列
	service_mq.MQ2SkillGoodConsumer()

	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
