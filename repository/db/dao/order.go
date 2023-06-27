package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
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
func (dao *OrderDao) ListOrderByCondition(uId uint, req *types.OrderListReq) (r []*types.OrderListResp, count int64, err error) {
	// TODO 商城算是一个TOC的应用，TOC的应该是不允许join操作的，看看后续怎么改走缓存，比如走缓存，找找免费的CDN之类的
	d := dao.DB.Model(&model.Order{}).
		Where("user_id = ?", uId)
	if req.Type != 0 {
		d.Where("type = ?", req.Type)
	}
	d.Count(&count) // 总数

	db := dao.DB.Model(&model.Order{}).
		Joins("AS o LEFT JOIN product AS p ON p.id = o.product_id").
		Joins("LEFT JOIN address AS a ON a.id = o.address_id").
		Where("o.user_id = ?", uId)
	if req.Type != 0 {
		db.Where("o.type = ?", req.Type)
	}
	db.Offset((req.PageNum - 1) * req.PageSize).
		Limit(req.PageSize).Order("created_at DESC").
		Select("o.id AS id," +
			"o.order_num AS order_num," +
			"UNIX_TIMESTAMP(o.created_at) AS created_at," +
			"UNIX_TIMESTAMP(o.updated_at) AS updated_at," +
			"o.user_id AS user_id," +
			"o.product_id AS product_id," +
			"o.boss_id AS boss_id," +
			"o.num AS num," +
			"o.type AS type," +
			"p.name AS name," +
			"p.discount_price AS discount_price," +
			"p.img_path AS img_path," +
			"a.name AS address_name," +
			"a.phone AS address_phone," +
			"a.address AS address").
		Find(&r)

	return
}

func (dao *OrderDao) GetOrderById(id, uId uint) (r *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).
		Where("id = ? AND user_id = ?", id, uId).
		First(&r).Error

	return
}

// ShowOrderById 获取订单详情
func (dao *OrderDao) ShowOrderById(id, uId uint) (r *types.OrderListResp, err error) {
	err = dao.DB.Model(&model.Order{}).
		Joins("AS o LEFT JOIN product AS p ON p.id = o.product_id").
		Joins("LEFT JOIN address AS a ON a.id = o.address_id").
		Where("o.id = ? AND o.user_id = ?", id, uId).
		Select("o.id AS id," +
			"o.order_num AS order_num," +
			"UNIX_TIMESTAMP(o.created_at) AS created_at," +
			"UNIX_TIMESTAMP(o.updated_at) AS updated_at," +
			"o.user_id AS user_id," +
			"o.product_id AS product_id," +
			"o.boss_id AS boss_id," +
			"o.num AS num," +
			"o.type AS type," +
			"p.name AS name," +
			"p.discount_price AS discount_price," +
			"p.img_path AS img_path," +
			"a.name AS address_name," +
			"a.phone AS address_phone," +
			"a.address AS address").
		Find(&r).Error

	return
}

// DeleteOrderById 获取订单详情
func (dao *OrderDao) DeleteOrderById(id, uId uint) error {
	return dao.DB.Model(&model.Order{}).
		Where("id=? AND user_id = ?", id, uId).
		Delete(&model.Order{}).Error
}

// UpdateOrderById 更新订单详情
func (dao *OrderDao) UpdateOrderById(id, uId uint, order *model.Order) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, uId).
		Updates(order).Error
}
