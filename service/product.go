package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	"mall/conf"
	"mall/consts"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/repository/db/dao"
	"mall/repository/db/model"
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

// ProductShow 商品
func (s *ProductSrv) ProductShow(ctx context.Context, req *types.ProductServiceReq) (resp interface{}, err error) {

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return types.Response{
		Status: code,
		Data:   types.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}, nil
}

// 创建商品
func (s *ProductSrv) ProductCreate(ctx context.Context, uId uint, files []*multipart.FileHeader, req *types.ProductServiceReq) (types.Response, error) {
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
		return types.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}, err
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
		util.LogrusObj.Error(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
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
			return types.Response{
				Status: code,
				Data:   e.GetMsg(code),
				Error:  path,
			}, err
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.ERROR
			return types.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}, err
		}
		wg.Done()
	}

	wg.Wait()

	return types.Response{
		Status: code,
		Data:   types.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *ProductSrv) ProductList(ctx context.Context, req *types.ProductServiceReq) (types.Response, error) {
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
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}, err
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, req.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return types.BuildListResponse(types.BuildProducts(products), uint(total)), nil
}

// ProductDelete 删除商品
func (s *ProductSrv) ProductDelete(ctx context.Context, req *types.ProductServiceReq) (types.Response, error) {
	code := e.SUCCESS

	err := dao.NewProductDao(ctx).DeleteProduct(req.ID)
	if err != nil {
		util.LogrusObj.Error(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return types.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}, nil
}

// 更新商品
func (s *ProductSrv) ProductUpdate(ctx context.Context, req *types.ProductServiceReq) (types.Response, error) {
	code := e.SUCCESS
	productDao := dao.NewProductDao(ctx)

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
	err := productDao.UpdateProduct(req.ID, product)
	if err != nil {
		util.LogrusObj.Error(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return types.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}, nil
}

// 搜索商品
func (s *ProductSrv) ProductSearch(ctx context.Context, req *types.ProductServiceReq) (types.Response, error) {
	code := e.SUCCESS
	if req.PageSize == 0 {
		req.PageSize = 15
	}

	productDao := dao.NewProductDao(ctx)
	products, err := productDao.SearchProduct(req.Info, req.BasePage)
	if err != nil {
		util.LogrusObj.Error(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return types.BuildListResponse(types.BuildProducts(products), uint(len(products))), nil
}

// ProductImgList 获取商品列表图片
func (s *ProductSrv) ProductImgList(ctx context.Context, req *types.ProductServiceReq) (types.Response, error) {
	productImgDao := dao.NewProductImgDao(ctx)
	productImgs, _ := productImgDao.ListProductImgByProductId(req.ID)
	return types.BuildListResponse(types.BuildProductImgs(productImgs), uint(len(productImgs))), nil
}
