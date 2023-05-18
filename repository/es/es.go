package es

import (
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"

	conf "github.com/CocaineCong/gin-mall/config"
)

var EsClient *elastic.Client

// InitEs 初始化es
func InitEs() {
	eConfig := conf.Config.Es
	esConn := fmt.Sprintf("http://%s:%s", eConfig.EsHost, eConfig.EsPort)
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConn))
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}

// EsHookLog 初始化log日志
func EsHookLog() *elogrus.ElasticHook {
	eConfig := conf.Config.Es
	hook, err := elogrus.NewElasticHook(EsClient, eConfig.EsHost, logrus.DebugLevel, eConfig.EsIndex)
	if err != nil {
		log.Panic(err)
	}
	return hook
}
