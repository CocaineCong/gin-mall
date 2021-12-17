package service

import (
	logging "github.com/sirupsen/logrus"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
)

type ListCategoriesService struct {

}


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