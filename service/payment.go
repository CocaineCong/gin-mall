package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"gorm.io/gorm"

	"mall/consts"
	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/encryption"
	"mall/pkg/utils/log"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/types"
)

var PaymentSrvIns *PaymentSrv
var PaymentSrvOnce sync.Once

type PaymentSrv struct {
}

func GetPaymentSrv() *PaymentSrv {
	PaymentSrvOnce.Do(func() {
		PaymentSrvIns = &PaymentSrv{}
	})
	return PaymentSrvIns
}

func (s *PaymentSrv) PayDown(ctx context.Context, req *types.PaymentDownReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error {
		uId := u.Id
		util.Encrypt.SetKey(req.Key)

		payment, err := dao.NewOrderDaoByDB(tx).GetOrderById(req.OrderId, uId)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}
		money := payment.Money
		num := payment.Num
		money = money * float64(num)

		userDao := dao.NewUserDaoByDB(tx)
		user, err := userDao.GetUserById(uId)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

		// 对钱进行解密。减去订单。再进行加密。
		moneyStr := util.Encrypt.AesDecoding(user.Money)
		moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
		if moneyFloat-money < 0.0 { // 金额不足进行回滚
			log.LogrusObj.Error(err)
			return errors.New("金币不足")
		}

		finMoney := fmt.Sprintf("%f", moneyFloat-money)
		user.Money = util.Encrypt.AesEncoding(finMoney)

		err = userDao.UpdateUserById(uId, user)
		if err != nil { // 更新用户金额失败，回滚
			log.LogrusObj.Error(err)
			return err
		}
		boss, err := userDao.GetUserById(uint(req.BossID))
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}
		moneyStr = util.Encrypt.AesDecoding(boss.Money)
		moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
		finMoney = fmt.Sprintf("%f", moneyFloat+money)
		boss.Money = util.Encrypt.AesEncoding(finMoney)

		err = userDao.UpdateUserById(uint(req.BossID), boss)
		if err != nil { // 更新boss金额失败，回滚
			log.LogrusObj.Error(err)
			return err
		}

		productDao := dao.NewProductDaoByDB(tx)
		product, err := productDao.GetProductById(uint(req.ProductID))
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}
		product.Num -= num
		err = productDao.UpdateProduct(uint(req.ProductID), product)
		if err != nil { // 更新商品数量减少失败，回滚
			log.LogrusObj.Error(err)
			return err
		}

		// 更新订单状态
		payment.Type = consts.OrderTypePendingShipping
		err = dao.NewOrderDaoByDB(tx).UpdateOrderById(req.OrderId, uId, payment)
		if err != nil { // 更新订单失败，回滚
			log.LogrusObj.Error(err)
			return err
		}

		productUser := model.Product{
			Name:          product.Name,
			CategoryID:    product.CategoryID,
			Title:         product.Title,
			Info:          product.Info,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Num:           num,
			OnSale:        false,
			BossID:        uId,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}

		err = productDao.CreateProduct(&productUser)
		if err != nil { // 买完商品后创建成了自己的商品失败。订单失败，回滚
			log.LogrusObj.Error(err)
			return err
		}

		return nil

	})

	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}
