package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

// 设置密码
func (Admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	Admin.PasswordDigest = string(bytes)
	return nil
}

// 校验密码
func (Admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Admin.PasswordDigest), []byte(password))
	return err == nil
}

// 封面地址
func (Admin *Admin) AvatarURL() string {
	//client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	//bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	//signedGetURL, _ := bucket.SignURL(admin.Avatar, oss.HTTPGet, 24*60*60)
	signedGetURL := "D:/CodeProjects/GoLandProjects/GoSuperMark/gin-mail/static/img/avatar/3.jpg"
	return signedGetURL
}
