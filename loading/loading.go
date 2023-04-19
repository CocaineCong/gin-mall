package loading

import (
	"mall/repository/cache"
	"mall/repository/db/dao"
)

func Loading() {
	dao.InitMySQL()
	cache.InitCache()
	lib.OpenCache()
	rabbitmq.InitRabbitMQ()
	// kfk.KafkaInit()
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
