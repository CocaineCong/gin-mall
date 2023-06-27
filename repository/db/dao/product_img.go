package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

// CreateProductImg 创建商品图片
func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	err = dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error

	return
}

// ListProductImgByProductId 根据商品id获取商品图片
func (dao *ProductImgDao) ListProductImgByProductId(pId uint) (r []*types.ProductImgResp, err error) {
	err = dao.DB.Model(&model.ProductImg{}).
		Where("product_id=?", pId).
		Find(&r).Error

	return
}
