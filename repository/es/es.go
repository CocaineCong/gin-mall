package es

import (
	"fmt"
	"log"

	"github.com/CocaineCong/eslogrus"
	elastic "github.com/elastic/go-elasticsearch"
	"github.com/sirupsen/logrus"

	conf "github.com/CocaineCong/gin-mall/config"
)

var EsClient *elastic.Client

// InitEs 初始化es
func InitEs() {
	eConfig := conf.Config.Es
	esConn := fmt.Sprintf("http://%s:%s", eConfig.EsHost, eConfig.EsPort)
	cfg := elastic.Config{
		Addresses: []string{esConn},
	}
	client, err := elastic.NewClient(cfg)
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}

// EsHookLog 初始化log日志
func EsHookLog() *eslogrus.ElasticHook {
	eConfig := conf.Config.Es
	hook, err := eslogrus.NewElasticHook(EsClient, eConfig.EsHost, logrus.DebugLevel, eConfig.EsIndex)
	if err != nil {
		log.Panic(err)
	}
	return hook
}
