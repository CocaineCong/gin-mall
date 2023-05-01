package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"mall/conf"
	"mall/consts"
)

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

// SetPassword 设置密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 头像地址
func (admin *Admin) AvatarURL() string {
	if conf.UploadModel == consts.UploadModelOss {
		return admin.Avatar
	}
	return conf.PhotoHost + conf.HttpPort + conf.AvatarPath + admin.Avatar
}
