package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
)

//商品图片创建的服务
type CreateImgServe struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

//商品详情图片创建的服务
type CreateInfoImgService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

//商品参数图片创建的服务
type CreateParamImgService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

//商品详情图片详情服务
type ShowInfoImgsService struct {
}

//商品参数图片详情
type ShowParamImgsService struct {
}

//创建商品图片
func (service *CreateImgServe) Create() serializer.Response {
	img := model.ProductParamImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&img).Error
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

//创建商品详情图片
func (service *CreateInfoImgService) Create() serializer.Response {
	infoImg := model.ProductParamImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&infoImg).Error
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

//创建商品参数图片
func (service *CreateParamImgService) Create() serializer.Response {
	paramImg := model.ProductParamImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	code := e.SUCCESS
	err := model.DB.Create(&paramImg).Error
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

func (service *ShowInfoImgsService) Show(id string) serializer.Response {
	var infoImgs []model.ProductInfoImg
	code := e.SUCCESS
	err := model.DB.Where("product_id=?", id).Find(&infoImgs).Error
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
		Data:   serializer.BuildInfoImgs(infoImgs),
		Msg:    e.GetMsg(code),
	}
}

//参数图片
func (service *ShowParamImgsService) Show(id string) serializer.Response {
	var paramImgs []model.ProductParamImg
	code := e.SUCCESS
	err := model.DB.Where("product_id=?", id).Find(&paramImgs).Error
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
		Data:   serializer.BuildParamImgs(paramImgs),
		Msg:    e.GetMsg(code),
	}
}
