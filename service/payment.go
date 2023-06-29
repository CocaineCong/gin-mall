package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"gorm.io/gorm"

	"github.com/CocaineCong/gin-mall/consts"
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
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

// TODO 目前买家和卖家的支付密码要一致，这个后续优化一下。。

// PayDown 支付操作
func (s *PaymentSrv) PayDown(ctx context.Context, req *types.PaymentDownReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error {
		uId := u.Id

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
		moneyFloat, err := user.DecryptMoney(req.Key)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}
		if moneyFloat-money < 0.0 { // 金额不足进行回滚
			log.LogrusObj.Error(err)
			return errors.New("金币不足")
		}

		finMoney := fmt.Sprintf("%f", moneyFloat-money)
		user.Money = finMoney
		user.Money, err = user.EncryptMoney(req.Key)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

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

		moneyFloat, _ = boss.DecryptMoney(req.Key)
		finMoney = fmt.Sprintf("%f", moneyFloat+money)
		boss.Money = finMoney
		boss.Money, err = boss.EncryptMoney(req.Key)
		if err != nil {
			log.LogrusObj.Error(err)
			return err
		}

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

	return
}
