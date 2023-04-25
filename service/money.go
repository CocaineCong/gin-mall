package service

import (
	"context"
	"sync"

	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/repository/db/dao"
	"mall/types"
)

var MoneySrvIns *MoneySrv
var MoneySrvOnce sync.Once

type MoneySrv struct {
}

func GetMoneySrv() *MoneySrv {
	MoneySrvOnce.Do(func() {
		MoneySrvIns = &MoneySrv{}
	})
	return MoneySrvIns
}

// MoneyShow 展示用户的金额
func (s *MoneySrv) MoneyShow(ctx context.Context, uId uint, req *types.ShowMoneyServiceReq) (types.Response, error) {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
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
		Data:   types.BuildMoney(user, req.Key),
		Msg:    e.GetMsg(code),
	}, nil
}
