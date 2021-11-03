package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/serializer"
	logging "github.com/sirupsen/logrus"
)

type ShowFavoritesService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

type CreateFavoritesService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
}

type DeleteFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
}

//商品收藏夹
func (service *ShowFavoritesService) Show(id string) serializer.Response {
	var favorites []model.Favorite
	total := 0
	code := e.SUCCESS
	if service.Limit == 0 {
		service.Limit = 12
	}
	if err := model.DB.Model(&favorites).Where("user_id=?", id).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err := model.DB.Where("user_id=?", id).Limit(service.Limit).Offset(&service.Start).Find(&favorites).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(favorites), uint(total))
}

//创建收藏夹
func (service *CreateFavoritesService) Create() serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS
	model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite)
	if favorite == (model.Favorite{}) {
		favorite = model.Favorite{
			UserID:    service.UserID,
			ProductID: service.ProductID,
			BossID:    service.BossID,
		}
		if err := model.DB.Create(&favorite).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//删除收藏夹
func (service *DeleteFavoriteService) Delete() serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS
	err := model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&favorite).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}
