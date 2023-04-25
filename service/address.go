package service

import (
	"context"
	"strconv"
	"sync"

	util "mall/pkg/utils"
	"mall/pkg/utils/ctl"
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

func (s *AddressSrv) Create(ctx context.Context, req *types.AddressServiceReq, uId uint) (resp interface{}, err error) {
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

func (s *AddressSrv) Show(ctx context.Context, aId string) (resp interface{}, err error) {
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	address, err := addressDao.GetAddressByAid(uint(addressId))
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccessWithData(address), nil
}

func (s *AddressSrv) List(ctx context.Context, uId uint) (resp interface{}, err error) {
	addresses, err := dao.NewAddressDao(ctx).ListAddressByUid(uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(addresses, int64(len(addresses))), nil
}

func (s *AddressSrv) Delete(ctx context.Context, aId, uId uint) (resp interface{}, err error) {
	err = dao.NewAddressDao(ctx).DeleteAddressById(aId, uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *AddressSrv) Update(ctx context.Context, req *types.AddressServiceReq, uid, aid uint) (resp interface{}, err error) {
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.UpdateAddressById(aid, address)
	var addresses []*types.AddressResp
	addresses, err = addressDao.ListAddressByUid(uid)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccessWithData(addresses), nil
}
