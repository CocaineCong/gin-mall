package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// ListCategory 分类列表
func (dao *CategoryDao) ListCategory() (category []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&category).Error
	return
}

// CreateProduct 创建商品
func (dao *CategoryDao) CreateCategory(category *model.Category) error {
	return dao.DB.Model(&model.Category{}).Create(&category).Error
}
