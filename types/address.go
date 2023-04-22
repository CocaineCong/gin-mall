package types

import (
	"mall/repository/db/model"
)

type AddressServiceReq struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type AddressResp struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

// 收货地址购物车
func BuildAddress(item *model.Address) *AddressResp {
	return &AddressResp{
		ID:       item.ID,
		UserID:   item.UserID,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		Seen:     false,
		CreateAt: item.CreatedAt.Unix(),
	}
}

// 收货地址列表
func BuildAddresses(items []*model.Address) (addresses []*AddressResp) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return addresses
}
