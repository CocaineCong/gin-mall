package dao

import (
	"context"

	"gorm.io/gorm"

	"mall/repository/db/model"
)

type SkillGoodsDao struct {
	*gorm.DB
}

func NewSkillGoodsDao(ctx context.Context) *SkillGoodsDao {
	return &SkillGoodsDao{NewDBClient(ctx)}
}

func (dao *SkillGoodsDao) Create(in *model.SkillProduct) error {
	return dao.Model(&model.SkillProduct{}).Create(&in).Error
}

func (dao *SkillGoodsDao) CreateByList(in []*model.SkillProduct) error {
	return dao.Model(&model.SkillProduct{}).Create(&in).Error
}

func (dao *SkillGoodsDao) ListSkillGoods() (resp []*model.SkillProduct, err error) {
	err = dao.Model(&model.SkillProduct{}).Where("num > 0").Find(&resp).Error
	return
}
