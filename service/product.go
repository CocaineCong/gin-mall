package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	logging "github.com/sirupsen/logrus"

	"mall/conf"
	"mall/consts"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/serializer"
	"mall/types"
)

var ProductSrvIns *ProductSrv
var ProductSrvOnce sync.Once

type ProductSrv struct {
}

func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}

// 商品
func (s *ProductSrv) Show(ctx context.Context, id string) serializer.Response {
	code := e.SUCCESS

	pId, _ := strconv.Atoi(id)

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
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

// 创建商品
func (s *ProductSrv) Create(ctx context.Context, uId uint, files []*multipart.FileHeader, req *types.ProductServiceReq) serializer.Response {
	var boss *model.User
	var err error
	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	var path string
	if conf.UploadModel == consts.UploadModelLocal {
		path, err = util.UploadProductToLocalStatic(tmp, uId, req.Name)
	} else {
		path, err = util.UploadToQiNiu(tmp, files[0].Size)
	}
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}
	product := &model.Product{
		Name:          req.Name,
		CategoryID:    uint(req.CategoryID),
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       path,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		if conf.UploadModel == consts.UploadModelLocal {
			path, err = util.UploadProductToLocalStatic(tmp, uId, req.Name+num)
		} else {
			path, err = util.UploadToQiNiu(tmp, file.Size)
		}
		if err != nil {
			code = e.ErrorUploadFile
			return serializer.Response{
				Status: code,
				Data:   e.GetMsg(code),
				Error:  path,
			}
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}

	wg.Wait()

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

func (s *ProductSrv) ProductList(ctx context.Context, req *types.ProductServiceReq) serializer.Response {
	var products []*model.Product
	var total int64
	code := e.SUCCESS

	if req.PageSize == 0 {
		req.PageSize = 15
	}
	condition := make(map[string]interface{})
	if req.CategoryID != 0 {
		condition["category_id"] = req.CategoryID
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, req.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// ProductDelete 删除商品
func (s *ProductSrv) ProductDelete(ctx context.Context, pId string) serializer.Response {
	code := e.SUCCESS

	productDao := dao.NewProductDao(ctx)
	productId, _ := strconv.Atoi(pId)
	err := productDao.DeleteProduct(uint(productId))
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

// 更新商品
func (s *ProductSrv) Update(ctx context.Context, pId string, req *types.ProductServiceReq) serializer.Response {
	code := e.SUCCESS
	productDao := dao.NewProductDao(ctx)

	productId, _ := strconv.Atoi(pId)
	product := &model.Product{
		Name:       req.Name,
		CategoryID: uint(req.CategoryID),
		Title:      req.Title,
		Info:       req.Info,
		// ImgPath:       service.ImgPath,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		OnSale:        req.OnSale,
	}
	err := productDao.UpdateProduct(uint(productId), product)
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

// 搜索商品
func (s *ProductSrv) Search(ctx context.Context, req *types.ProductServiceReq) serializer.Response {
	code := e.SUCCESS
	if req.PageSize == 0 {
		req.PageSize = 15
	}

	productDao := dao.NewProductDao(ctx)
	products, err := productDao.SearchProduct(req.Info, req.BasePage)
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

// List 获取商品列表图片
func (s *ProductSrv) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImgByProductId(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
