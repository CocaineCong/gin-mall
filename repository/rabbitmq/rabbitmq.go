package rabbitmq

import (
	"strings"

	"github.com/streadway/amqp"

	conf "mall/config"
)

// RabbitMQ rabbitMQ链接单例
var RabbitMQ *amqp.Connection

// InitRabbitMQ 在中间件中初始化rabbitMQ链接
func InitRabbitMQ() {
	rConfig := conf.Config.RabbitMq
	pathRabbitMQ := strings.Join([]string{rConfig.RabbitMQ, "://", rConfig.RabbitMQUser, ":", rConfig.RabbitMQPassWord, "@", rConfig.RabbitMQHost, ":", rConfig.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(pathRabbitMQ)
	if err != nil {
		panic(err)
	}
	RabbitMQ = conn
}
