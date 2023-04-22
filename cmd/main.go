package main

import (
	"mall/conf"
	"mall/loading"
	"mall/routes"
)

func main() {
	loading.Loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
