package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	logging "github.com/sirupsen/logrus"
	"FanOneMall/serializer"
)

type ShowCountService struct {

}

//订单详细
func (service *ShowCountService) Show (id string) serializer.Response{
	 code := e.SUCCESS
	 var favoriteTotal int
	 var notPayTotal int
	 var payTotal int
	 if err:=model.DB.Model(model.Favorite{}).Where("user_id=?",id).Count(&favoriteTotal).Error;err!=nil{
	 	logging.Info(err)
	 	code = e.ErrorDatabase
	 	return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	 }
	 if err := model.DB.Model(model.Order{}).Where("user_id=? AND type=?",id,1).Count(&notPayTotal).Error;err!=nil{
	 	logging.Info(err)
	 	code = e.ErrorDatabase
	 	return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	 }
	 if err :=model.DB.Model(model.Order{}).Where("user_id=? AND type=?",id,2).Count(&payTotal).Error;err!=nil{
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
		 Data:   serializer.BuildCount(favoriteTotal,notPayTotal,payTotal),
		 Msg:    e.GetMsg(code),
	 }
}