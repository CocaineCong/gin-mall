package cache

import (
	"fmt"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"

	conf "mall/config"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

func InitCache() {
	Redis()
}

// Redis 在中间件中初始化redis链接
func Redis() {
	rConfig := conf.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDbName,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
