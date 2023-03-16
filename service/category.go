package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"mall/dao"
	"mall/pkg/e"
	"mall/model"
	"mall/serializer"
)

type ListCategoriesService struct {
}

func (service *ListCategoriesService) List(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.ListCategory()
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
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCategories(categories),
	}
}

// CreateCategoryService 
type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"categoCreateCategoryry_id"`
	CategoryName string `form:"category_name" json:"category_name"`
}

// Create 创建分类
func (service *CreateCategoryService) Create(ctx context.Context) serializer.Response {
	category := model.Category{
		CategoryID:   service.CategoryID,
		CategoryName: service.CategoryName,
	}
	code := e.SUCCESS
	categoryDao := dao.NewCategoryDao(ctx)
	err := categoryDao.CreateCategory(&category)
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
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCategory(&category),
	}
}
