package service

import (
	"context"
	"sync"

	"mall/pkg/e"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/serializer"
	"mall/types"

	logging "github.com/sirupsen/logrus"
)

var CartSrvIns *CartSrv
var CartSrvOnce sync.Once

type CartSrv struct {
}

func GetCartSrv() *CartSrv {
	CartSrvOnce.Do(func() {
		CartSrvIns = &CartSrv{}
	})
	return CartSrvIns
}

// CartCreate 创建购物车
func (s *CartSrv) CartCreate(ctx context.Context, uId uint, req *types.CartServiceReq) (serializer.Response, error) {
	var product *model.Product
	code := e.SUCCESS

	// 判断有无这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ProductId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}

	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	cart, status, _ := cartDao.CreateCart(req.ProductId, uId, req.BossID)
	if status == e.ErrorProductMoreCart {
		return serializer.Response{
			Status: status,
			Msg:    e.GetMsg(status),
		}, err
	}

	userDao := dao.NewUserDao(ctx)
	boss, _ := userDao.GetUserById(req.BossID)
	return serializer.Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   serializer.BuildCart(cart, product, boss),
	}, nil
}

// CartList 购物车
func (s *CartSrv) CartList(ctx context.Context, uId uint) (serializer.Response, error) {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCarts(carts),
		Msg:    e.GetMsg(code),
	}, nil
}

// CartUpdate 修改购物车信息
func (s *CartSrv) CartUpdate(ctx context.Context, req *types.CartServiceReq) (serializer.Response, error) {
	code := e.SUCCESS

	cartDao := dao.NewCartDao(ctx)
	err := cartDao.UpdateCartNumById(req.Id, req.UserId, req.Num)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}, nil
}

// CartDelete 删除购物车
func (s *CartSrv) CartDelete(ctx context.Context, req *types.CartServiceReq) (serializer.Response, error) {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartById(req.Id, req.UserId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}, nil
}
