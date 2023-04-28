package service

import (
	"context"
	"errors"
	"sync"

	"mall/pkg/e"
	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/log"
	"mall/repository/db/dao"
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
func (s *CartSrv) CartCreate(ctx context.Context, uId uint, req *types.CartServiceReq) (resp interface{}, err error) {
	// 判断有无这个商品
	_, err = dao.NewProductDao(ctx).GetProductById(req.ProductId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	_, status, _ := cartDao.CreateCart(req.ProductId, uId, req.BossID)
	if status == e.ErrorProductMoreCart {
		err = errors.New(e.GetMsg(status))
		return
	}
	return ctl.RespSuccess(), nil
}

// CartList 购物车
func (s *CartSrv) CartList(ctx context.Context, uId uint, req *types.CartListReq) (resp interface{}, err error) {
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(carts, int64(len(carts))), nil // TODO 无分页，之后考虑要不要加
}

// CartUpdate 修改购物车信息
func (s *CartSrv) CartUpdate(ctx context.Context, req *types.UpdateCartServiceReq) (resp interface{}, err error) {
	err = dao.NewCartDao(ctx).UpdateCartNumById(req.Id, req.UserId, req.Num)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}

// CartDelete 删除购物车
func (s *CartSrv) CartDelete(ctx context.Context, req *types.CartServiceReq) (resp interface{}, err error) {
	err = dao.NewCartDao(ctx).DeleteCartById(req.Id, req.UserId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}
