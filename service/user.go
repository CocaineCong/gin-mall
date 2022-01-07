package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
	"mall/conf"
	"mall/model"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/serializer"
	"mime/multipart"
	"strings"
	"time"
)

//UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=10"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Key      string `form:"key" json:"key" binding:"required"`
}

type UserLoginService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

//用户修改信息的服务
type UserUpdateService struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=5,max=10"`
}

type UploadAvatarService struct {
}

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailService struct {
}

func (service UserRegisterService) Register() serializer.Response {
	var user model.User
	var count int64
	code := e.SUCCESS
	if service.Key=="" || len(service.Key)!=16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	conf.Encryption.SetKey(service.Key)
	model.DB.Model(&model.User{}).Where("user_name=?",service.UserName).Count(&count)
	if count == 1 {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		Nickname: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Money:    conf.Encryption.AesEncoding("10000"),
	}
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}


//Login 用户登陆函数
func (service UserLoginService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

//Update 用户修改信息
func (service UserUpdateService) Update(id uint) serializer.Response {
	var user model.User
	code := e.SUCCESS
	//找到用户
	err := model.DB.First(&user, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if service.NickName != "" {
		user.Nickname = service.NickName
	}
	err = model.DB.Save(&user).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}


func (service *UploadAvatarService) Post(id uint, file multipart.File,fileSize int64) serializer.Response {
	var user model.User
	code := e.SUCCESS
	status, info := UploadToQiNiu(file, fileSize)
	if status != 200 {
		return serializer.Response{
			Status: status,
			Data:   e.GetMsg(status),
			Error:  info,
		}
	}
	model.DB.Where("id=?", id).First(&user)
	user.Avatar = info
	model.DB.Save(&user)
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// Send 发送邮件
func (service *SendEmailService) Send(id uint) serializer.Response {
	code := e.SUCCESS
	var address string
	var notice model.Notice
	token, err := util.GenerateEmailToken(id, service.OperationType, service.Email, service.Password)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//数据库里 对应邮件id = operation_type+1
	if err := model.DB.First(&notice, service.OperationType+1).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token
	fmt.Println(address)
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "VaildAddress", address, -1)
	fmt.Println(mailText)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "FanOne")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		logging.Info(err)
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//验证内容
func (service ValidEmailService) Valid(token string) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	//验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			logging.Info(err)
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}

	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if operationType == 1 {
		//1:绑定邮箱
		if err := model.DB.Table("user").Where("id=?", userID).Update("email", email).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if operationType == 2 {
		//2：解绑邮箱
		if err := model.DB.Table("user").Where("id=?", userID).Update("email", "").Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user, userID).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//3：修改密码
	if operationType == 3 {
		// 加密密码
		if err := user.SetPassword(password); err != nil {
			logging.Info(err)
			code = e.ErrorFailEncryption
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err := model.DB.Save(&user).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.UpdatePasswordSuccess
		//返回修改密码成功信息
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//返回用户信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
