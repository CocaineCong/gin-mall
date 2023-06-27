package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

// GetAddressByAid 根据 AddressId 获取 Address
func (dao *AddressDao) GetAddressByAid(aId, uId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("id = ? AND user_id = ?", aId, uId).First(&address).
		Error

	return
}

// ListAddressByUid 根据 User Id 获取User
func (dao *AddressDao) ListAddressByUid(uid uint) (r []*types.AddressResp, err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("user_id = ?", uid).
		Order("created_at desc").
		Select("id, user_id, name, phone, address, UNIX_TIMESTAMP(created_at)").
		Find(&r).Error

	return
}

// CreateAddress 创建地址
func (dao *AddressDao) CreateAddress(address *model.Address) (err error) {
	return dao.DB.Model(&model.Address{}).
		Create(&address).Error
}

// DeleteAddressById 根据 id 删除地址
func (dao *AddressDao) DeleteAddressById(aId, uId uint) (err error) {
	return dao.DB.Where("id = ? AND user_id = ?", aId, uId).
		Delete(&model.Address{}).Error
}

// UpdateAddressById 通过 id 修改地址信息
func (dao *AddressDao) UpdateAddressById(aId uint, address *model.Address) (err error) {
	return dao.DB.Model(&model.Address{}).
		Where("id=?", aId).
		Updates(address).Error
}
