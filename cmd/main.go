package main

import (
	"fmt"

	conf "github.com/CocaineCong/gin-mall/config"
	util "github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/pkg/utils/track"
	"github.com/CocaineCong/gin-mall/repository/cache"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/repository/es"
	"github.com/CocaineCong/gin-mall/repository/kafka"
	"github.com/CocaineCong/gin-mall/routes"

	_ "github.com/apache/skywalking-go"
)

func main() {
	loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动配成功...")
}

// loading一些配置
func loading() {
	conf.InitConfig()
	dao.InitMySQL()
	cache.InitCache()
	// rabbitmq.InitRabbitMQ() // 如果需要接入RabbitMQ可以打开这个注释
	es.InitEs() // 如果需要接入ELK可以打开这个注释
	kafka.InitKafka()
	track.InitJaeger()
	util.InitLog() // 如果接入ELK请进入这个func打开注释
	fmt.Println("加载配置完成...")
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
