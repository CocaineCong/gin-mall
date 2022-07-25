package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// GetProductById 通过 id 获取product
func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).
		First(&product).Error
	return
}

// ListProductByCondition 获取商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Preload("Category").Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}

// CreateProduct 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	err = dao.DB.Model(&model.Product{}).Create(&product).Error
	return
}

// CountProductByCondition 根据情况获取商品的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// DeleteProduct 删除商品
func (dao *ProductDao) DeleteProduct(pId uint) (err error) {
	err = dao.DB.Model(&model.Product{}).Delete(&model.Product{}).Error
	return
}

// UpdateProduct 更新商品
func (dao *ProductDao) UpdateProduct(pId uint, product *model.Product) (err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", pId).
		Updates(&product).Error
	return
}

// SearchProduct 搜索商品
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}
