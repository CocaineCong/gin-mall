package main

import (
	"FanOneMall/conf"
	"FanOneMall/routes"
	"FanOneMall/service"
)

func main() {
	//从配置文件读入配置
	conf.Init()
	service.ListenOrder()
	//转载路由
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
