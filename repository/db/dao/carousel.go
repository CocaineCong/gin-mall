package dao

import (
	"context"

	"gorm.io/gorm"

	"mall/repository/db/model"
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

func (dao *CarouselDao) ListCarousel() (carousels []*model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}
