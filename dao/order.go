package dao

import (
	"context"

	"gorm.io/gorm"

	"mall/model"
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
func (dao *OrderDao) CreateOrder(order *model.Order) error {
	return dao.DB.Create(&order).Error
}

// ListOrderByCondition 获取订单List
func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Order("created_at desc").Find(&orders).Error
	return
}

// GetOrderById 获取订单详情
func (dao *OrderDao) GetOrderById(id uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=?", id).
		First(&order).Error
	return
}

// GetOrderByOrderNum 通过orderNum获取订单
func (dao *OrderDao) GetOrderByOrderNum(OrderNum uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("order_num=?", OrderNum).
		First(&order).Error
	return
}

// CloseOrderById 关闭订单
func (dao *OrderDao) CloseOrderById(id uint) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", id).Update("type", 3).Error
}

// UpdateOrderById 更新订单详情
func (dao *OrderDao) UpdateOrderById(id uint, order *model.Order) error {
	return dao.DB.Where("id=?", id).Updates(order).Error
}
