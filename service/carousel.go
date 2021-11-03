package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	logging "github.com/sirupsen/logrus"
	"FanOneMall/serializer"
)

type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

//ListCarouselsService 视频列表服务
type ListCarouselsService struct {

}

//创建轮播图
func (service *CreateCarouselService) Create() serializer.Response {
	carousel := model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&carousel).Error
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
		Data:   serializer.BuildCarousel(carousel),
		Msg:    e.GetMsg(code),
	}
}


//视频列表
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
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCarousels(carousels),
		Msg:    e.GetMsg(code),
	}
}
