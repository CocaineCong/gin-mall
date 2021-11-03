package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	logging "github.com/sirupsen/logrus"
	"FanOneMall/serializer"
)

//公告详情服务
type ShowNoticeService struct {
}

//公告创建服务
type CreateNoticeService struct {
	Text string `form:"text" json:"text"`
}

//修改公共
type UpdateNoticeService struct {
	NoticeID uint   `form:"notice_id" json:"notice_id"`
	Text     string `form:"text" json:"text"`
}

//Show公告详情服务
func (service *ShowNoticeService) Show() serializer.Response {
	var notice model.Notice
	code := e.SUCCESS
	if err := model.DB.First(&notice, 1).Error; err != nil {
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
		Data:   serializer.BuildNotice(notice),
		Msg:    e.GetMsg(code),
	}
}

//创建公告
func (service *CreateNoticeService) Create() serializer.Response {
	notice := model.Notice{
		Text: service.Text,
	}
	code := e.SUCCESS
	err := model.DB.Create(&notice).Error
	if err != nil {
		logging.Info(err)
		code := e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//公告update
func (service *UpdateNoticeService) Update() serializer.Response {
	var notice model.Notice
	code := e.SUCCESS
	if err := model.DB.First(&notice, service.NoticeID).Error; err != nil {
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
	}
}
