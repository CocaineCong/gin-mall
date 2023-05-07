package main

import (
	"fmt"

	conf "mall/config"
	util "mall/pkg/utils/log"
	"mall/pkg/utils/track"
	"mall/repository/cache"
	"mall/repository/db/dao"
	"mall/repository/kafka"
)

func Loading() {
	// Ek1+Ep1==Ek2+Ep2
	conf.InitConfig()
	dao.InitMySQL()
	cache.InitCache()
	// rabbitmq.InitRabbitMQ() // 如果需要接入RabbitMQ可以打开这个注释
	// es.InitEs() // 如果需要接入ELK可以打开这个注释
	kafka.InitKafka()
	track.InitJaeger()
	util.InitLog() // 如果接入ELK请进入这个func打开注释
	fmt.Println("加载配置完成...")
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
