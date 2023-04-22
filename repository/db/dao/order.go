package dao

import (
	"context"

	"gorm.io/gorm"

	model2 "mall/repository/db/model"
	"mall/types"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

// CreateOrder 创建订单
func (dao *OrderDao) CreateOrder(order *model2.Order) error {
	return dao.DB.Create(&order).Error
}

// ListOrderByCondition 获取订单List
func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page *types.BasePage) (orders []*model2.Order, total int64, err error) {
	err = dao.DB.Model(&model2.Order{}).Where(condition).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model2.Order{}).Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Order("created_at desc").Find(&orders).Error
	return
}

// GetOrderById 获取订单详情
func (dao *OrderDao) GetOrderById(id, uId uint) (order *model2.Order, err error) {
	err = dao.DB.Model(&model2.Order{}).
		Where("id = ? AND user_id = ?", id, uId).
		First(&order).Error
	return
}

// DeleteOrderById 获取订单详情
func (dao *OrderDao) DeleteOrderById(id, uId uint) error {
	return dao.DB.Where("id=? AND uId = ?", id, uId).Delete(&model2.Order{}).Error
}

// UpdateOrderById 更新订单详情
func (dao *OrderDao) UpdateOrderById(id uint, order *model2.Order) error {
	return dao.DB.Where("id=?", id).Updates(order).Error
}
