package service

import (
	"context"
	"sync"

	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/log"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/types"
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

func (s *AddressSrv) AddressCreate(ctx context.Context, req *types.AddressCreateReq, uId uint) (resp interface{}, err error) {
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uId,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *AddressSrv) AddressGet(ctx context.Context, req *types.AddressGetReq) (resp interface{}, err error) {
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(req.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccessWithData(address), nil
}

func (s *AddressSrv) AddressList(ctx context.Context, req *types.AddressListReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	addresses, err := dao.NewAddressDao(ctx).
		ListAddressByUid(u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(addresses, int64(len(addresses))), nil
}

func (s *AddressSrv) AddressDelete(ctx context.Context, req *types.AddressDeleteReq, uId uint) (resp interface{}, err error) {
	err = dao.NewAddressDao(ctx).DeleteAddressById(req.Id, uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *AddressSrv) AddressUpdate(ctx context.Context, req *types.AddressServiceReq, uid, aid uint) (resp interface{}, err error) {
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.UpdateAddressById(aid, address)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var addresses []*types.AddressResp
	addresses, err = addressDao.ListAddressByUid(uid)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccessWithData(addresses), nil
}
