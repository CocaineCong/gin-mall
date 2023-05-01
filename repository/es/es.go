package es

import (
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"

	"mall/conf"
)

var EsClient *elastic.Client

// InitEs 初始化es
func InitEs() {
	esConn := fmt.Sprintf("http://%s:%s", conf.EsHost, conf.EsPort)
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConn))
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}

// EsHookLog 初始化log日志
func EsHookLog() *elogrus.ElasticHook {
	hook, err := elogrus.NewElasticHook(EsClient, conf.EsHost, logrus.DebugLevel, conf.EsIndex)
	if err != nil {
		log.Panic(err)
	}
	return hook
}
