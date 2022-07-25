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

// GetUserById 根据 id 获取用户
func (dao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(&user).Error
	return
}
