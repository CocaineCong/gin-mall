package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string //`gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"unique"`
	Status         string
	Limit          int    // 0 非管理员  1 管理员
	Type           int    // 0表示用户  1表示商家
	Avatar         string `gorm:"size:1000"`
	Monery 		   int
}

const (
	PassWordCost        = 12         //密码加密难度
	Active       string = "active"   //激活用户
	Inactive     string = "inactive" //未激活用户
	Suspend      string = "suspend"  //被封禁用户
)

//GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

//AvatarUrl 封面地址
func (user *User) AvatarURL() string {
	//client ,_ := oss.New(os.Getenv("OSS_END_POINT"),os.Getenv("OSS_ACCESS_KEY_ID"),os.Getenv("OSS_ACCESS_KEY_SECRET"))
	//bucket,_ := client.Bucket(os.Getenv("OSS_BUCKET"))
	//signedGetURL,_:=bucket.SignURL(user.Avatar,oss.HTTPGet,24*60*60)
	//client ,_ := oss.New("oss-cn-beijing.aliyuncs.com" ,"LTAI4G9m5vcxdmrjLDuG5Uuf","DNmjvI8n5rHFwhtIyPNtgsYBLyLaVa")
	//bucket,_ := client.Bucket("gomall1" )
	//signedGetURL,_:=bucket.SignURL(user.Avatar,oss.HTTPGet,24*60*60)
	signedGetURL := user.Avatar
	fmt.Println("xxxxxxxxxx", signedGetURL)
	return signedGetURL
}
