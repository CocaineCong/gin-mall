package service

import (
	"context"
	"mime/multipart"
	"strings"
	"sync"
	"time"

	"mall/conf"
	"mall/consts"
	"mall/pkg/e"
	util "mall/pkg/utils"
	dao2 "mall/repository/db/dao"
	model2 "mall/repository/db/model"
	"mall/serializer"
	"mall/types"

	logging "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) serializer.Response {
	var user *model2.User
	code := e.SUCCESS
	if req.Key == "" || len(req.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	util.Encrypt.SetKey(req.Key)
	userDao := dao2.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
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
	user = &model2.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model2.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if conf.UploadModel == consts.UploadModelOss {
		user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	} else {
		user.Avatar = "avatar.JPG"
	}
	// 创建用户
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

// Login 用户登陆函数
func (s *UserSrv) Login(ctx context.Context, req *types.UserServiceReq) serializer.Response {
	var user *model2.User
	code := e.SUCCESS
	userDao := dao2.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if !exist { // 如果查询不到，返回相应的错误
		logging.Info(err)
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(req.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, req.UserName, 0)
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

// Update 用户修改信息
func (s *UserSrv) Update(ctx context.Context, uId uint, req *types.UserServiceReq) serializer.Response {
	var user *model2.User
	var err error
	code := e.SUCCESS
	// 找到用户
	userDao := dao2.NewUserDao(ctx)
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
	if req.NickName != "" {
		user.NickName = req.NickName
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

func (s *UserSrv) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64, req *types.UserServiceReq) serializer.Response {
	code := e.SUCCESS
	var user *model2.User
	var err error

	userDao := dao2.NewUserDao(ctx)
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
	var path string
	if conf.UploadModel == consts.UploadModelLocal { // 兼容两种存储方式
		path, err = util.UploadAvatarToLocalStatic(file, uId, user.UserName)
	} else {
		path, err = util.UploadToQiNiu(file, fileSize)
	}
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}

	user.Avatar = path
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
func (s *UserSrv) Send(ctx context.Context, id uint, req *types.SendEmailServiceReq) serializer.Response {
	code := e.SUCCESS
	var address string
	var notice *model2.Notice

	token, err := util.GenerateEmailToken(id, req.OperationType, req.Email, req.Password)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	noticeDao := dao2.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(req.OperationType)
	if err != nil {
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
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", req.Email)
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

// Valid 验证内容
func (s *UserSrv) Valid(ctx context.Context, token string, req *types.ValidEmailServiceReq) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS

	// 验证token
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

	// 获取该用户信息
	userDao := dao2.NewUserDao(ctx)
	user, err := userDao.GetUserById(userID)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		// 1:绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 2：解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		// 3：修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userID, user)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 成功则返回用户的信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
