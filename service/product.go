package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	"mall/conf"
	"mall/consts"
	"mall/pkg/utils/ctl"
	"mall/pkg/utils/log"
	util "mall/pkg/utils/upload"
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
	product, err := dao.NewProductDao(ctx).ShowProductById(req.ID)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccessWithData(product), nil
}

// 创建商品
func (s *ProductSrv) ProductCreate(ctx context.Context, uId uint, files []*multipart.FileHeader, req *types.ProductCreateReq) (resp interface{}, err error) {
	boss, _ := dao.NewUserDao(ctx).GetUserById(uId)
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	var path string
	if conf.UploadModel == consts.UploadModelLocal {
		path, err = util.UploadProductToLocalStatic(tmp, uId, req.Name)
	} else {
		path, err = util.UploadToQiNiu(tmp, files[0].Size)
	}
	if err != nil {
		log.LogrusObj.Error(err)
		return
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
		log.LogrusObj.Error(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		tmp, _ = file.Open()
		if conf.UploadModel == consts.UploadModelLocal {
			path, err = util.UploadProductToLocalStatic(tmp, uId, req.Name+num)
		} else {
			path, err = util.UploadToQiNiu(tmp, file.Size)
		}
		if err != nil {
			log.LogrusObj.Error(err)
			return
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = dao.NewProductImgDaoByDB(productDao.DB).CreateProductImg(productImg)
		if err != nil {
			log.LogrusObj.Error(err)
			return
		}
		wg.Done()
	}

	wg.Wait()

	return ctl.RespSuccess(), nil
}

func (s *ProductSrv) ProductList(ctx context.Context, req *types.ProductListReq) (resp interface{}, err error) {
	var products []*types.ProductResp
	var total int64
	condition := make(map[string]interface{})
	if req.CategoryID != 0 {
		condition["category_id"] = req.CategoryID
	}
	productDao := dao.NewProductDao(ctx)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, req.BasePage)
		wg.Done()
	}()
	wg.Wait()
	total, err = productDao.CountProductByCondition(condition)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(products, total), nil
}

// ProductDelete 删除商品
func (s *ProductSrv) ProductDelete(ctx context.Context, req *types.ProductDeleteReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	err = dao.NewProductDao(ctx).DeleteProduct(req.ID, u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

// 更新商品
func (s *ProductSrv) ProductUpdate(ctx context.Context, req *types.ProductServiceReq) (resp interface{}, err error) {
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
	err = dao.NewProductDao(ctx).UpdateProduct(req.ID, product)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}

// 搜索商品
// TODO 后续用脚本同步数据MySQL到ES，用ES进行搜索
func (s *ProductSrv) ProductSearch(ctx context.Context, req *types.ProductServiceReq) (resp interface{}, err error) {
	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(req.Info, req.BasePage)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(products, count), nil
}

// ProductImgList 获取商品列表图片
func (s *ProductSrv) ProductImgList(ctx context.Context, req *types.ProductServiceReq) (resp interface{}, err error) {
	productImgs, _ := dao.NewProductImgDao(ctx).ListProductImgByProductId(req.ID)
	return ctl.RespList(productImgs, int64(len(productImgs))), nil
}
