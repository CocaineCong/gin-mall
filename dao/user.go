package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) GetUserById(uId uint) (user model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}