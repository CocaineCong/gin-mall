package main

import (
	"mall/conf"
	"mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_=r.Run(conf.HttpPort)
}
