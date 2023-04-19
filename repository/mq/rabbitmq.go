package mq

import (
	"strings"

	"github.com/streadway/amqp"

	"mall/conf"
)

// RabbitMQ rabbitMQ链接单例
var RabbitMQ *amqp.Connection

// InitRabbitMQ 在中间件中初始化rabbitMQ链接
func InitRabbitMQ() {
	pathRabbitMQ := strings.Join([]string{conf.RabbitMQ, "://", conf.RabbitMQUser, ":", conf.RabbitMQPassWord, "@", conf.RabbitMQHost, ":", conf.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(pathRabbitMQ)
	if err != nil {
		panic(err)
	}
	RabbitMQ = conn
}
