package service

import (
	"context"

	logging "github.com/sirupsen/logrus"

	"mall/pkg/e"
	"mall/repository/db/dao"
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
