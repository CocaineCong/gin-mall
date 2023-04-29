package service

import (
	"context"
	"sync"

	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/encryption"
	"mall/pkg/utils/log"
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
func (s *MoneySrv) MoneyShow(ctx context.Context, uId uint, req *types.MoneyShowReq) (resp interface{}, err error) {
	user, err := dao.NewUserDao(ctx).GetUserById(uId)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	util.Encrypt.SetKey(req.Key)
	mResp := &types.MoneyShowResp{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserMoney: util.Encrypt.AesDecoding(user.Money),
	}
	return ctl.RespSuccessWithData(mResp), nil
}
