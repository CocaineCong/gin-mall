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

func (dao *ProductDao) GetProductById(id int) (product model.Product, err error) {
	err = dao.Model(&model.Product{}).Where("id=?", id).
		First(&product).Error
	return
}

func (dao *ProductDao) ListProduct(id int) (products []model.Product, err error) {
	return
}
