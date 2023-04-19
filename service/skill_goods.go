package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
	"mall/cache"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/serializer"
	"mime/multipart"
	"strconv"
	"time"
)

const SkillGoodsLockKey = "order_timeout_queue"

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

type SkillGoodMessage struct {
	*model.SkillGood2MQ
	OrderNum string //分布式唯一订单号
}

// Import 通过excel导入商品数据
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

// InitSkillGoods 直接放到这里，初始化秒杀商品信息，将mysql的信息存入redis中
func (service *SkillGoodsService) InitSkillGoods(ctx context.Context) error {
	skillGoods, _ := dao.NewSkillGoodsDao(ctx).ListSkillGoods()
	r := cache.RedisClient
	// 加载到redis 设置商品的库存和价格
	for i := range skillGoods {
		fmt.Println(*skillGoods[i])
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "num", skillGoods[i].Num)
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "money", skillGoods[i].Money)
	}
	return nil
}

func (service *SkillGoodsService) SkillGoods(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
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
		code = e.ErrorRedissonSecKillGoods
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
		Error:  "",
	}
}

func (service *SkillGoodsService) GetSkillGoodsResult(ctx context.Context, uId uint) serializer.Response {
	// redis下单缓存中 查找订单
	code := 200
	orderNum, err := GetOrderCache(int(service.ProductId), int(uId))
	if err != nil {
		code = e.ErrorRedis
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if len(orderNum) == 0 { //没有下单缓存 说明还在mq队列中
		return serializer.Response{
			Status: code,
			Msg:    "排队中",
		}
	} else {
		return serializer.Response{ //有下单缓存 秒杀成功
			Status: code,
			Data:   nil,
			Msg:    "秒杀成功",
		}
	}
}

// RedissonSecKillGoods 加锁
func RedissonSecKillGoods(sk *model.SkillGood2MQ) error {
	productId := strconv.Itoa(int(sk.ProductId))
	userId := strconv.Itoa(int(sk.UserId))

	snowflakeID, err := util.SnowflakeID()
	if err != nil {
		util.LogrusObj.Infoln("get lock fail", err)
		return errors.New("snowflakeID 生成失败")
	}
	uuid := snowflakeID + productId + userId

	// 通过SetNX 进行redis分布式锁的获取
	lockSuccess, err := cache.RedisClient.SetNX(SkillGoodsLockKey+productId, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		util.LogrusObj.Infoln("get SkillGoodsLock fail", err)
		return errors.New("获取秒杀商品锁失败")
	}

	// 查看订单信息缓存 校验用户是否重复秒杀
	orderId, _ := GetOrderCache(int(sk.ProductId), int(sk.UserId)) // 订单key在消费端进行添加
	if len(orderId) > 0 {
		util.LogrusObj.Infoln("user ready has skillGoodOrder")
		return errors.New("user ready has skillGoodOrder")
	}

	err = SendSecKillGoodsToMQ(sk, uuid) //在分布式锁里面  发送秒杀商品消息给MQ
	if err != nil || !lockSuccess {
		util.LogrusObj.Infoln("send SecKillGoods to MQ fail", err)
		return errors.New("send SecKillGoods to MQ fail")
	}

	value, _ := cache.RedisClient.Get(SkillGoodsLockKey + productId).Result() //获取锁
	if value == uuid {                                                        //compare value,if equal then del  验证锁是不是自己加的  是的话就进行锁删除
		_, err = cache.RedisClient.Del(SkillGoodsLockKey + productId).Result()
		if err != nil {
			util.LogrusObj.Infoln("unlock fail")
			return nil
		} else {
			util.LogrusObj.Infoln("unlock success")
		}
	}
	return nil
}

// SendSecKillGoodsToMQ 传送到MQ
func SendSecKillGoodsToMQ(sk *model.SkillGood2MQ, orderUuid string) error {
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

	sKillGoodMessage := SkillGoodMessage{
		SkillGood2MQ: sk,
		OrderNum:     orderUuid,
	}

	body, _ := json.Marshal(sKillGoodMessage)
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

//func getUuid(gid string) string {
//	codeLen := 8
//	// 1. 定义原始字符串
//	rawStr := "jkwangagDGFHGSERKILMJHSNOPQR546413890_"
//	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
//	buf := make([]byte, 0, codeLen)
//	b := bytes.NewBuffer(buf)
//	// 随机从中获取
//	rand.Seed(time.Now().UnixNano())
//	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
//		randNum := rand.Intn(rawStrLen)
//		b.WriteByte(rawStr[randNum])
//	}
//	return b.String() + gid
//}

// GetOrderCache 检验redis中是否已经存在该用户的秒杀订单
func GetOrderCache(gid, userID int) (orderId string, err error) {
	if orderId, err = cache.RedisClient.Get(fmt.Sprintf("order:%d:%d", userID, gid)).Result(); err == redis.Nil {
		err = nil
	} else {
		util.LogrusObj.Infof("redis.Get() failed, err: %v", err)
		err = errors.New(e.GetMsg(e.ErrorRedis))
	}
	return
}

// CreateOrderCache 新增订单缓存
func CreateOrderCache(gid, userID int, OrderNum uint64) (err error) {
	// 订单缓存永久生效 直到删除
	if err = cache.RedisClient.Set(fmt.Sprintf("order:%d:%d", userID, gid), OrderNum, -1).Err(); err != nil {
		util.LogrusObj.Infof("redis.Set() failed, err: %v, order:%d:%d", err, userID, gid)
		err = errors.New(e.GetMsg(e.ErrorRedis))
	}
	return
}

// DeleteOrderCache 删除订单缓存 仅在撤销订单时才会调用
func DeleteOrderCache(gid, userID int) (err error) {
	if err = cache.RedisClient.Del(fmt.Sprintf("order:%d:%d", userID, gid)).Err(); err != nil {
		util.LogrusObj.Infof("redis.Del() falied, err: %v, order:%d:%d", err, userID, gid)
		err = errors.New(e.GetMsg(e.ErrorRedis))
	}
	return
}

// CreateSkillGoodOrder 创建的mysql订单涉 创建redis下单缓存 加入redis订单超时队列
func CreateSkillGoodOrder(db *gorm.DB, order *model.Order) error {
	// 查看订单信息缓存 校验用户是否重复秒杀
	orderId, _ := GetOrderCache(int(order.ProductID), int(order.UserID))
	if len(orderId) > 0 {
		return errors.New("user ready has skillGoodOrder")
	}

	orderDao := dao.NewOrderDaoByDB(db)
	err := orderDao.CreateOrder(order)

	// 创建订单信息缓存
	if err = CreateOrderCache(int(order.ProductID), int(order.UserID), order.OrderNum); err != nil {
		return err
	}
	//将订单号存入Redis,并设置过期时间  设置分数是当前时间加上15分钟
	data := redis.Z{
		Score:  float64(time.Now().Unix()) + 15*time.Minute.Seconds(),
		Member: order.OrderNum,
	}
	if err := cache.RedisClient.ZAdd(OrderTimeOutKey, data).Err(); err != nil {
		log.Printf("订单【%d】加入延迟队列失败, err: %v", order.OrderNum, err)
		return err
	}
	return nil
}
