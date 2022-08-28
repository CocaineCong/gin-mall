package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	logging "github.com/sirupsen/logrus"
	"mall/cache"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"mall/types"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
)

type SkillGoodsImport struct {
}

// 限购一个
type SkillGoodsService struct {
	ProductId uint `json:"product_id" form:"product_id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
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
	for index, colCell := range rows[1:] {
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
		r.HSet(strconv.Itoa(int(skillGoods[i].Id)), "num", skillGoods[i].Num)
		r.HSet(strconv.Itoa(int(skillGoods[i].Id)), "money", skillGoods[i].Money)
	}
	return nil
}

// 秒杀
func (service *SkillGoodsService) SkillGoods(ctx context.Context, uId uint) serializer.Response {
	r := cache.RedisClient

}

// 加锁
func RedissonSecKillGoods(uId, pId, bossId uint, money float64) error {
	p := strconv.Itoa(int(pId))
	uuid := getUuid(p)
	lockSuccess, err := cache.RedisClient.SetNX(p, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return errors.New("get lock fail")
	} else {
		fmt.Println("get lock success")
	}

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
func SendSecKillGoodsToMQ(uId, pId, bossId uint, money float64) error {
	infoSend := types.SkillGood2MQ{
		ProductId: pId,
		BossId:    bossId,
		UserId:    uId,
		Money:     money,
	}
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
