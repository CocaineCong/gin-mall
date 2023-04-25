package main

import (
	"fmt"

	"mall/conf"
	"mall/routes"
)

func main() {
	Loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
	fmt.Println("启动配成功...")
}
