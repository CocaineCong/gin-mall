package service

import (
	"context"
	"strconv"
	"sync"

	logging "github.com/sirupsen/logrus"

	"mall/pkg/e"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/serializer"
	"mall/types"
)

var AddressLogicIns *AddressLogic
var AddressLogicOnce sync.Once

type AddressLogic struct {
}

func GetAddressLogic() *AddressLogic {
	AddressLogicOnce.Do(func() {
		AddressLogicIns = &AddressLogic{}
	})
	return AddressLogicIns
}

func (s *AddressLogic) Create(ctx context.Context, req *types.AddressServiceReq, uId uint) (resp interface{}, err error) {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uId,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	addressDao = dao.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, err = addressDao.ListAddressByUid(uId)
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
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *AddressLogic) Show(ctx context.Context, aId string) (resp interface{}, err error) {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	address, err := addressDao.GetAddressByAid(uint(addressId))
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
		Data:   serializer.BuildAddress(address),
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *AddressLogic) List(ctx context.Context, uId uint) (resp interface{}, err error) {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.ListAddressByUid(uId)
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
		Data:   serializer.BuildAddresses(address),
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *AddressLogic) Delete(ctx context.Context, aId, uId uint) (serializer.Response, error) {
	addressDao := dao.NewAddressDao(ctx)
	code := e.SUCCESS
	err := addressDao.DeleteAddressById(aId, uId)
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

func (s *AddressLogic) Update(ctx context.Context, req *types.AddressServiceReq, uid, aid uint) (serializer.Response, error) {
	code := e.SUCCESS

	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err := addressDao.UpdateAddressById(aid, address)
	addressDao = dao.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, err = addressDao.ListAddressByUid(uid)
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
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}, nil
}
