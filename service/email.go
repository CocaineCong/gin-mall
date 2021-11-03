package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/util"
	"FanOneMall/serializer"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
	"time"
)

type SendEmailService struct {
	UserID        uint   `form:"user_id" json:"user_id"`
	Email         string `form:"email" json:"email"`
	Password      string `form:"password" json:"password"`
	OperationType uint   `form:"operation_type" json:"operation_type"` //operationType 1.绑定邮箱 2.解绑邮箱 3.改密码
}

//send 发送邮箱
func (service *SendEmailService) Send() serializer.Response {
	code := e.SUCCESS
	var address string
	var notice model.Notice
	token, err := util.GenerateEmailToken(service.UserID, service.OperationType, service.Email, service.Password)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//数据库里面对应邮件id=operation_type+1
	if err := model.DB.First(&notice, service.OperationType+1).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = "http://localhost:8080/#/vaild/email/" + token
	mailStr := notice.Text
	//mailText := strings.Replace(mailStr, "VaildAddress", address, -1)
	mailText := mailStr + "请点击确认身份" +address
	m := mail.NewMessage()
	m.SetHeader("From", "cocainecong@163.com")
	m.SetHeader("To", service.Email)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")抄送
	m.SetHeader("Subject", "LWGG后花园")
	m.SetBody("text/html", mailText)

	d := mail.NewDialer("smtp.163.com", 465, "cocainecong@163.com", "UAZGCZCVOSIZPOLH")
	d.StartTLSPolicy = mail.MandatoryStartTLS
	fmt.Println("ERROR1",*d)
	fmt.Println("ERROR2",*m)

	//发邮件
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

//VaildEmailService 绑定、解绑邮箱和修改密码的服务
type VaildEmailService struct {
	Token string `form:"token" json:"token"`
}

//Vaild 绑定邮箱
func (service *VaildEmailService) Vaild() serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	//验证token
	if service.Token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(service.Token)
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
		//1.绑定邮箱
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
		//2.解绑邮箱
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
	//3.修改密码
	if operationType == 3 {
		//加密密码
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
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}
