package serializer

import "FanOneMall/model"

type Order struct {
	ID uint `json:"id"`
	OrderNum uint64 `json:"order_num"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	UserID uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	BossID uint `json:"boss_id"`
	Num uint `json:"num"`
	AddressName string `json:"address_name"`
	AddressPhone string `json:"address_phone"`
	Address string `json:"address"`
	Type uint `json:"type"`
	Name string `json:"name"`
	ImgPath string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

func BuildOrder(item1 model.Order, item2 model.Product, item3 model.Address) Order {
	return Order{
		ID:            item1.ID,
		OrderNum:      item1.OrderNum,
		CreatedAt:     item1.CreatedAt.Unix(),
		UpdatedAt:     item1.UpdatedAt.Unix(),
		UserID:        item1.UserID,
		ProductID:     item1.ProductID,
		BossID :	   item1.BossID,
		Num:           item1.Num,
		AddressName:   item3.Name,
		AddressPhone:  item3.Phone,
		Address:       item3.Address,
		Type:          item1.Type,
		Name:          item2.Name,
		ImgPath:       item2.ImgPath,
		DiscountPrice: item2.DiscountPrice,
	}
}

func BuildOrders(items []model.Order) (orders []Order) {
	for _, item1 := range items {
		item2 := model.Product{}
		item3 := model.Address{}
		err := model.DB.First(&item2, item1.ProductID).Error
		err = model.DB.First(&item3, item1.AddressID).Error
		if err != nil {
			continue
		}
		order := BuildOrder(item1, item2, item3)
		orders = append(orders, order)
	}
	return orders
}
