package service

import (
	"context"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
	"mall/conf"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/serializer"
	"mime/multipart"
	"strings"
	"time"
)

// UserService 管理用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type SendEmailService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailService struct {
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额
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
	err = userDao.CreateUser(user)
	if err != nil {
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
func (service UserService) Login(ctx context.Context) serializer.Response {
	var user model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist { //如果查询不到，返回相应的错误
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
func (service UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user model.User
	var err error
	code := e.SUCCESS
	//找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if service.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserById(uId, user)
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

func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.SUCCESS
	var user model.User
	var err error

	status, info := UploadToQiNiu(file, fileSize)
	if status != 200 {
		return serializer.Response{
			Status: status,
			Data:   e.GetMsg(status),
			Error:  info,
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = info
	err = userDao.UpdateUserById(uId, user)
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
	if err := dao.DB.First(&notice, service.OperationType).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
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
		if err := dao.DB.Table("user").Where("id=?", userID).Update("email", email).Error; err != nil {
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
		if err := dao.DB.Table("user").Where("id=?", userID).Update("email", "").Error; err != nil {
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
	if err := dao.DB.First(&user, userID).Error; err != nil {
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
		if err := dao.DB.Save(&user).Error; err != nil {
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
