package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewNewCarouselDao(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListCarousel() (r []*types.ListCarouselResp, err error) {
	err = dao.DB.Model(&model.Carousel{}).
		Select("id, img_path, product_id, UNIX_TIMESTAMP(created_at)").
		Find(&r).Error

	return
}
