package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
)

// 收藏创建的服务
type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
}

// 分类列表服务
type ListCategoriesService struct {
}

//创建分类
func (service *CreateCategoryService) Create() serializer.Response {
	category := model.Category{
		CategoryID:   service.CategoryID,
		CategoryName: service.CategoryName,
	}
	code := e.SUCCESS
	err := model.DB.Create(&category).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCategory(category),
		Msg:    e.GetMsg(code),
	}
}

//分类列表
func (service *ListCategoriesService) List() serializer.Response {
	var categories []model.Category
	code := e.SUCCESS
	if err := model.DB.Find(&categories).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCategories(categories),
	}
}
