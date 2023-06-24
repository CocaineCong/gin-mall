package service

import (
	"context"
	"sync"

	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	util "github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
)

var AddressSrvIns *AddressSrv
var AddressSrvOnce sync.Once

type AddressSrv struct {
}

func GetAddressSrv() *AddressSrv {
	AddressSrvOnce.Do(func() {
		AddressSrvIns = &AddressSrv{}
	})
	return AddressSrvIns
}

func (s *AddressSrv) AddressCreate(ctx context.Context, req *types.AddressCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  u.Id,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *AddressSrv) AddressShow(ctx context.Context, req *types.AddressGetReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	address, err := dao.NewAddressDao(ctx).GetAddressByAid(req.Id, u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	resp = &types.AddressResp{
		ID:        address.ID,
		UserID:    address.UserID,
		Name:      address.Name,
		Phone:     address.Phone,
		Address:   address.Address,
		CreatedAt: address.CreatedAt.Unix(),
	}

	return
}

func (s *AddressSrv) AddressList(ctx context.Context, req *types.AddressListReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	resp, err = dao.NewAddressDao(ctx).
		ListAddressByUid(u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (s *AddressSrv) AddressDelete(ctx context.Context, req *types.AddressDeleteReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewAddressDao(ctx).DeleteAddressById(req.Id, u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return
}

func (s *AddressSrv) AddressUpdate(ctx context.Context, req *types.AddressServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  u.Id,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.UpdateAddressById(req.Id, address)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var addresses []*types.AddressResp
	addresses, err = addressDao.ListAddressByUid(u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	resp = &types.DataListResp{
		Item:  addresses,
		Total: int64(len(addresses)),
	}

	return

}
