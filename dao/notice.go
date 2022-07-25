package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 通过id获取notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}

// CreateNotice 创建notice
func (dao *NoticeDao) CreateNotice(notice *model.Notice) error {
	return dao.DB.Model(&model.Notice{}).Create(&notice).Error
}
