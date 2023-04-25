package service

import (
	"context"
	"sync"

	util "mall/pkg/utils"
	"mall/pkg/utils/ctl"
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
func (s *MoneySrv) MoneyShow(ctx context.Context, uId uint, req *types.ShowMoneyServiceReq) (resp interface{}, err error) {
	user, err := dao.NewUserDao(ctx).GetUserById(uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	util.Encrypt.SetKey(req.Key)
	mResp := &types.MoneyResp{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserMoney: util.Encrypt.AesDecoding(user.Money),
	}
	return ctl.RespSuccessWithData(mResp), nil
}
