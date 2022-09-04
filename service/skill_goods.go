package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	logging "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"mall/cache"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
)

type SkillGoodsImport struct {
}

// 限购一个
type SkillGoodsService struct {
	SkillGoodsId uint   `json:"skill_goods_id" form:"skill_goods_id"`
	ProductId    uint   `json:"product_id" form:"product_id"`
	BossId       uint   `json:"boss_id" form:"boss_id"`
	AddressId    uint   `json:"address_id" form:"address_id"`
	Key          string `json:"key" form:"key"`
}

func (service *SkillGoodsImport) Import(ctx context.Context, file multipart.File) serializer.Response {
	xlFile, err := xlsx.OpenReader(file)
	if err != nil {
		logging.Info(err)
	}
	code := e.SUCCESS
	rows := xlFile.GetRows("Sheet1")
	length := len(rows[1:])
	skillGoods := make([]*model.SkillGoods, length, length)
	for index, colCell := range rows {
		if index == 0 {
			continue
		}
		pId, _ := strconv.Atoi(colCell[0])
		bId, _ := strconv.Atoi(colCell[1])
		num, _ := strconv.Atoi(colCell[3])
		money, _ := strconv.ParseFloat(colCell[4], 64)
		skillGood := &model.SkillGoods{
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
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "上传失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 直接放到这里，初始化秒杀商品信息，将mysql的信息存入redis中
func (service *SkillGoodsService) InitSkillGoods(ctx context.Context) error {
	skillGoods, _ := dao.NewSkillGoodsDao(ctx).ListSkillGoods()
	r := cache.RedisClient
	// 加载到redis
	for i := range skillGoods {
		fmt.Println(*skillGoods[i])
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "num", skillGoods[i].Num)
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "money", skillGoods[i].Money)
	}
	return nil
}

func (service *SkillGoodsService) SkillGoods(ctx context.Context, uId uint) serializer.Response {
	mo, _ := cache.RedisClient.HGet("SK"+strconv.Itoa(int(service.SkillGoodsId)), "money").Float64()
	sk := &model.SkillGood2MQ{
		ProductId:   service.ProductId,
		BossId:      service.BossId,
		UserId:      uId,
		AddressId:   service.AddressId,
		Key:         service.Key,
		Money:       mo,
		SkillGoodId: service.SkillGoodsId,
	}
	err := RedissonSecKillGoods(sk)
	if err != nil {
		return serializer.Response{}
	}
	return serializer.Response{}
}

// 加锁
func RedissonSecKillGoods(sk *model.SkillGood2MQ) error {
	p := strconv.Itoa(int(sk.ProductId))
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
	if value == uuid { //compare value,if equal then del
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
func SendSecKillGoodsToMQ(sk *model.SkillGood2MQ) error {
	ch, err := model.MQ.Channel()
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
