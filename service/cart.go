package service

import (
	"context"
	"errors"
	"sync"

	"github.com/CocaineCong/gin-mall/pkg/e"
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	util "github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/types"
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
func (s *CartSrv) CartCreate(ctx context.Context, req *types.CartCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 判断有无这个商品
	_, err = dao.NewProductDao(ctx).GetProductById(req.ProductId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	_, status, _ := cartDao.CreateCart(req.ProductId, u.Id, req.BossID)
	if status == e.ErrorProductMoreCart {
		err = errors.New(e.GetMsg(status))
		return
	}
	return
}

// CartList 购物车
func (s *CartSrv) CartList(ctx context.Context, req *types.CartListReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	carts, err := dao.NewCartDao(ctx).ListCartByUserId(u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	resp = &types.DataListResp{
		Item:  carts, // TODO 无分页，之后考虑要不要加
		Total: int64(len(carts)),
	}

	return
}

// CartUpdate 修改购物车信息
func (s *CartSrv) CartUpdate(ctx context.Context, req *types.UpdateCartServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewCartDao(ctx).UpdateCartNumById(req.Id, u.Id, req.Num)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return
}

// CartDelete 删除购物车
func (s *CartSrv) CartDelete(ctx context.Context, req *types.CartDeleteReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewCartDao(ctx).DeleteCartById(req.Id, u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return
}
