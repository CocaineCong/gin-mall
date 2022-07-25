package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type FavoritesDao struct {
	*gorm.DB
}

func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBClient(ctx)}
}

func NewFavoritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

// ListFavoriteByUserId 通过 user_id 获取收藏夹列表
func (dao *FavoritesDao) ListFavoriteByUserId(uId uint, pageSize, pageNum int) (favorites []*model.Favorite, total int64, err error) {
	// 总数
	err = dao.DB.Model(&model.Favorite{}).Preload("User").
		Where("user_id=?", uId).Count(&total).Error
	if err != nil {
		return
	}
	// 分页
	err = dao.DB.Model(model.Favorite{}).Preload("User").Where("user_id=?", uId).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).Find(&favorites).Error
	return
}

// CreateFavorite 创建收藏夹
func (dao *FavoritesDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = dao.DB.Create(&favorite).Error
	return
}

// FavoriteExistOrNot 判断是否存在
func (dao *FavoritesDao) FavoriteExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).
		Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if count == 0 || err != nil {
		return false, err
	}
	return true, err

}

// DeleteFavoriteById 删除收藏夹
func (dao *FavoritesDao) DeleteFavoriteById(fId uint) (err error) {
	err = dao.DB.Where("id=?", fId).Delete(&model.Favorite{}).Error
	return
}
