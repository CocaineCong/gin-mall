package service

import (
	"context"
	"sync"

	logging "github.com/sirupsen/logrus"

	"mall/pkg/e"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/types"
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
func (s *CartSrv) CartCreate(ctx context.Context, uId uint, req *types.CartServiceReq) (types.Response, error) {
	var product *model.Product
	code := e.SUCCESS

	// 判断有无这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ProductId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}

	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	cart, status, _ := cartDao.CreateCart(req.ProductId, uId, req.BossID)
	if status == e.ErrorProductMoreCart {
		return types.Response{
			Status: status,
			Msg:    e.GetMsg(status),
		}, err
	}

	userDao := dao.NewUserDao(ctx)
	boss, _ := userDao.GetUserById(req.BossID)
	return types.Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   types.BuildCart(cart, product, boss),
	}, nil
}

// CartList 购物车
func (s *CartSrv) CartList(ctx context.Context, uId uint) (types.Response, error) {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return types.Response{
		Status: code,
		Data:   types.BuildCarts(carts),
		Msg:    e.GetMsg(code),
	}, nil
}

// CartUpdate 修改购物车信息
func (s *CartSrv) CartUpdate(ctx context.Context, req *types.CartServiceReq) (types.Response, error) {
	code := e.SUCCESS

	cartDao := dao.NewCartDao(ctx)
	err := cartDao.UpdateCartNumById(req.Id, req.UserId, req.Num)
	if err != nil {
		logging.Info(err)
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

// CartDelete 删除购物车
func (s *CartSrv) CartDelete(ctx context.Context, req *types.CartServiceReq) (types.Response, error) {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartById(req.Id, req.UserId)
	if err != nil {
		logging.Info(err)
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
