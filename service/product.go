package service

import (
	logging "github.com/sirupsen/logrus"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"mime/multipart"
)

//更新商品的服务
type ProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
	PageNum       int    `form:"pageNum"`
	PageSize      int    `form:"pageSize"`
}


type ListProductImgService struct {

}

// 商品
func (service *ProductService) Show(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
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
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

//创建商品
func (service *ProductService)Create(id uint,files []*multipart.FileHeader) serializer.Response {
	var boss model.User
	model.DB.Model(&model.User{}).Where("id = ?",id).First(&boss)
	tmp,_ := files[0].Open()
	status , info := UploadToQiNiu(tmp,files[0].Size)
	if status != 200 {
		return serializer.Response{
			Status:  status  ,
			Data:      e.GetMsg(status),
			Error:info,
		}
	}
	product := model.Product{
		Name:          service.Name,
		CategoryID:    uint(service.CategoryID),
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       info,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		Num:           service.Num,
		OnSale:        true,
		BossID:        int(id),
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	code := e.SUCCESS
	err := model.DB.Create(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	for _,file := range files {
		tmp, _ := file.Open()
		status, info := UploadToQiNiu(tmp, file.Size)
		if status != 200 {
			return serializer.Response{
				Status: status,
				Data:   e.GetMsg(status),
				Error:  info,
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   info,
		}
		err = model.DB.Create(&productImg).Error
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

func (service *ProductService) List() serializer.Response {
	var products []model.Product
	var total int64
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if service.CategoryID == 0 {
		if err := model.DB.Model(model.Product{}).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Offset((service.PageNum - 1) * service.PageSize).
			Limit(service.PageSize).Find(&products).
			Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(model.Product{}).Preload("Category").
			Where("category_id = ?", service.CategoryID).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Model(model.Product{}).Preload("Category").
			Where("category_id=?", service.CategoryID).
			Offset((service.PageNum - 1) * service.PageSize).
			Limit(service.PageSize).
			Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	var productList  []serializer.Product
	for _, item := range products {
		products := serializer.BuildProduct(item)
		productList = append(productList, products)
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

//删除商品
func (service *ProductService) Delete(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&product).Error
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
	}
}

//更新商品
func (service *ProductService) Update(id string) serializer.Response {
	var product model.Product
	model.DB.Model(&model.Product{}).First(&product,id)
	product.Name=service.Name
	product.CategoryID=uint(service.CategoryID)
	product.Title=service.Title
	product.Info=service.Info
	product.ImgPath=service.ImgPath
	product.Price=service.Price
	product.DiscountPrice=service.DiscountPrice
	product.OnSale=service.OnSale
	code := e.SUCCESS
	err := model.DB.Save(&product).Error
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
	}
}

//搜索商品
func (service *ProductService) Search() serializer.Response {
	var products []model.Product
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	err := model.DB.Where("name LIKE ? OR info LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Offset((service.PageNum - 1) * service.PageSize).
		Limit(service.PageSize).Find(&products).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(len(products)))
}

//获取商品列表图片
func (service *ListProductImgService) List(id string)serializer.Response {
	var productImgList []model.ProductImg
	model.DB.Model(model.ProductImg{}).Where("product_id=?",id).Find(&productImgList)
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgList),uint(len(productImgList)))
}