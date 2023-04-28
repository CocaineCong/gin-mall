package main

import (
	"fmt"

	"mall/conf"
	util "mall/pkg/utils/log"
	"mall/repository/cache"
	"mall/repository/db/dao"
	"mall/repository/es"
	"mall/repository/mq"
)

func Loading() {
	// Ek1+Ep1==Ek2+Ep2
	conf.Init()
	dao.InitMySQL()
	cache.InitCache()
	mq.InitRabbitMQ() // 如果需要接入RabbitMQ可以打开这个注释
	es.InitEs()       // 如果需要接入ELK可以打开这个注释
	util.InitLog()
	fmt.Println("加载配置完成...")
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
