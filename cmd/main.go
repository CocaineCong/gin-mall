package main

import (
	"mall/conf"
	"mall/routes"
)

func main() {
	Loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
