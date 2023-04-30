package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"strconv"
	"sync"
	"time"

	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/streadway/amqp"

	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/log"
	"mall/repository/cache"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/repository/mq"
	"mall/types"
)

var SkillProductSrvIns *SkillProductSrv
var SkillProductSrvOnce sync.Once

type SkillProductSrv struct {
}

func GetSkillProductSrv() *SkillProductSrv {
	SkillProductSrvOnce.Do(func() {
		SkillProductSrvIns = &SkillProductSrv{}
	})
	return SkillProductSrvIns
}

func (s *SkillProductSrv) Import(ctx context.Context, file multipart.File) (resp interface{}, err error) {
	xlFile, err := xlsx.OpenReader(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	rows := xlFile.GetRows("Sheet1")
	length := len(rows[1:])
	skillGoods := make([]*model.SkillProduct, length, length)
	for index, colCell := range rows {
		if index == 0 {
			continue
		}
		pId, _ := strconv.Atoi(colCell[0])
		bId, _ := strconv.Atoi(colCell[1])
		num, _ := strconv.Atoi(colCell[3])
		money, _ := strconv.ParseFloat(colCell[4], 64)
		skillGood := &model.SkillProduct{
			ProductId: uint(pId),
			BossId:    uint(bId),
			Title:     colCell[2],
			Money:     money,
			Num:       num,
		}
		skillGoods[index-1] = skillGood
	}
	err = dao.NewSkillGoodsDao(ctx).CreateByList(skillGoods)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}

// 直接放到这里，初始化秒杀商品信息，将mysql的信息存入redis中
func (s *SkillProductSrv) InitSkillGoods(ctx context.Context) (interface{}, error) {
	skillGoods, _ := dao.NewSkillGoodsDao(ctx).ListSkillGoods()
	r := cache.RedisClient
	// 加载到redis
	for i := range skillGoods {
		fmt.Println(*skillGoods[i])
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "num", skillGoods[i].Num)
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "money", skillGoods[i].Money)
	}
	return nil, nil
}

func (s *SkillProductSrv) SkillProduct(ctx context.Context, req *types.SkillProductServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	mo, _ := cache.RedisClient.HGet("SK"+strconv.Itoa(int(req.SkillProductId)), "money").Float64()
	sk := &model.SkillProduct2MQ{
		ProductId:      req.ProductId,
		BossId:         req.BossId,
		UserId:         uId,
		AddressId:      req.AddressId,
		Key:            req.Key,
		Money:          mo,
		SkillProductId: req.SkillProductId,
	}
	err = RedissonSecKillGoods(sk)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

// 加锁
func RedissonSecKillGoods(sk *model.SkillProduct2MQ) error {
	p := strconv.Itoa(int(sk.SkillProductId))
	uuid := getUuid(p)
	_, err := cache.RedisClient.Del(p).Result()
	lockSuccess, err := cache.RedisClient.SetNX(p, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return errors.New("get lock fail")
	} else {
		fmt.Println("get lock success")
	}
	_ = SendSecKillGoodsToMQ(sk)
	value, _ := cache.RedisClient.Get(p).Result()
	if value == uuid { // compare value,if equal then del
		_, err := cache.RedisClient.Del(p).Result()
		if err != nil {
			fmt.Println("unlock fail")
			return nil
		} else {
			fmt.Println("unlock success")
		}
	}
	return nil
}

// 传送到MQ
func SendSecKillGoodsToMQ(sk *model.SkillProduct2MQ) error {
	ch, err := mq.RabbitMQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ err:" + err.Error())
		return err
	}
	q, err := ch.QueueDeclare("skill_goods", true, false, false, false, nil)
	if err != nil {
		err = errors.New("rabbitMQ err:" + err.Error())
		return err
	}
	body, _ := json.Marshal(sk)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ err:" + err.Error())
		return err
	}
	log.Printf("Sent %s", body)
	return nil
}

func getUuid(gid string) string {
	codeLen := 8
	// 1. 定义原始字符串
	rawStr := "jkwangagDGFHGSERKILMJHSNOPQR546413890_"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	// 随机从中获取
	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String() + gid
}

// 取消订单的操作,redis的商品回退
