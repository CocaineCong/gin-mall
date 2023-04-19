package serializer

import (
	"context"

	"mall/conf"
	"mall/consts"
	dao2 "mall/repository/db/dao"
	model2 "mall/repository/db/model"
)

type Order struct {
	ID            uint   `json:"id"`
	OrderNum      uint64 `json:"order_num"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	BossID        uint   `json:"boss_id"`
	Num           uint   `json:"num"`
	AddressName   string `json:"address_name"`
	AddressPhone  string `json:"address_phone"`
	Address       string `json:"address"`
	Type          uint   `json:"type"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

func BuildOrder(item1 *model2.Order, item2 *model2.Product, item3 *model2.Address) Order {
	o := Order{
		ID:            item1.ID,
		OrderNum:      item1.OrderNum,
		CreatedAt:     item1.CreatedAt.Unix(),
		UpdatedAt:     item1.UpdatedAt.Unix(),
		UserID:        item1.UserID,
		ProductID:     item1.ProductID,
		BossID:        item1.BossID,
		Num:           uint(item1.Num),
		AddressName:   item3.Name,
		AddressPhone:  item3.Phone,
		Address:       item3.Address,
		Type:          item1.Type,
		Name:          item2.Name,
		ImgPath:       conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + item2.ImgPath,
		DiscountPrice: item2.DiscountPrice,
	}

	if conf.UploadModel == consts.UploadModelOss {
		o.ImgPath = item2.ImgPath
	}

	return o
}

func BuildOrders(ctx context.Context, items []*model2.Order) (orders []Order) {
	productDao := dao2.NewProductDao(ctx)
	addressDao := dao2.NewAddressDao(ctx)

	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressID)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)
	}
	return orders
}
