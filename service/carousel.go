package service

import (
	logging "github.com/sirupsen/logrus"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
)

type ListCarouselsService struct {

}

func (service *ListCarouselsService) List() serializer.Response {
	var carousels []model.Carousel
	code := e.SUCCESS
	if err := model.DB.Find(&carousels).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels),uint(len(carousels)))
}

