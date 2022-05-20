package service

import (
	logging "github.com/sirupsen/logrus"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
)

type FavoritesService struct {
	ProductID uint `form:"product_id" json:"product_id"`
	BossID    uint ` form:"boss_id" json:"boss_id"`
	PageNum     	int 	  `form:"pageNum"`
	PageSize    	int 	  `form:"pageSize"`
}


//商品收藏夹
func (service *FavoritesService) Show(id uint) serializer.Response {
	var favorites []model.Favorite
	var total int64
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if err := model.DB.Model(&favorites).Preload("User").
		Where("user_id=?", id).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err := model.DB.Model(model.User{}).Preload("User").Where("user_id=?", id).
		Offset((service.PageNum - 1) * service.PageSize).
		Limit(service.PageSize).Find(&favorites).Error
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
func (service *FavoritesService) Create(id uint) serializer.Response {
	var favorite model.Favorite
	var user model.User
	var boss model.User
	var product model.Product
	code := e.SUCCESS
	model.DB.Where("user_id=? AND product_id=?",id, service.ProductID).Find(&favorite)
	model.DB.Model(model.User{}).First(&user,id)
	model.DB.Model(model.User{}).First(&boss,service.BossID)
	model.DB.Model(model.Product{}).First(&product,service.ProductID)
	if favorite == (model.Favorite{}) {
		favorite = model.Favorite{
			UserID:     id,
			User:		user,
			ProductID:  service.ProductID,
			Product:	product,
			BossID:     service.BossID,
			Boss:		boss,
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
func (service *FavoritesService) Delete(uid uint,pid string) serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS
	err := model.DB.Where("user_id=? AND product_id=?", uid, pid).Find(&favorite).Error
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
