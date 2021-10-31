package service

import (
	"FanOneMall/cache"
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
	"fmt"
	"strings"
)

//展示排行的服务
type ListRankingService struct {
}

//展示家电的排行
type ListElecRankingService struct {
}

//展现配件排行
type ListAcceRankingService struct {
}

//获取排行
func (service *ListRankingService) List() serializer.Response {
	var products []model.Product
	code := e.SUCCESS
	pros, _ := cache.RedisClient.ZRevRange(cache.RankKey, 0, 9).Result()
	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id,%s)", strings.Join(pros, ","))
		err := model.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			logging.Info(err)
			code := e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProducts(products),
		Msg:    e.GetMsg(code),
	}
}

//家电排行
func (service *ListElecRankingService) List() serializer.Response {
	var products []model.Product
	code := e.SUCCESS
	pros, _ := cache.RedisClient.ZRevRange(cache.ElectricalRank, 0, 6).Result()
	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id,%s)", strings.Join(pros, ","))
		err := model.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			logging.Info(err)
			code := e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProducts(products),
		Msg:    e.GetMsg(code),
	}
}

//配件排行
func (service *ListAcceRankingService) List() serializer.Response {
	var products []model.Product
	code := e.SUCCESS
	// 从redis读取点击前十
	pros, _ := cache.RedisClient.ZRevRange(cache.AccessoryRank, 0, 6).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(pros, ","))
		err := model.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			logging.Info(err)
			code := e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
