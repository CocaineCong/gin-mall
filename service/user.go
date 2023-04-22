package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"sync"
	"time"

	"mall/conf"
	"mall/consts"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/pkg/utils/ctl"
	"mall/pkg/utils/email"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/types"
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

// UserRegister 用户注册
func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	var user *model.User
	if req.Key == "" || len(req.Key) != 16 {
		return nil, errors.New("密钥长度不足")
	}
	util.Encrypt.SetKey(req.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	if exist {
		return nil, errors.New("用户已经存在了")
	}
	user = &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		util.LogrusObj.Error(err)
		return
	}

	if conf.UploadModel == consts.UploadModelOss {
		user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	} else {
		user.Avatar = "avatar.JPG"
	}

	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}

// UserLogin 用户登陆函数
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if !exist { // 如果查询不到，返回相应的错误
		util.LogrusObj.Error(err)
		return nil, errors.New("用户不存在")
	}

	if user.CheckPassword(req.Password) == false {
		return nil, errors.New("账号/密码不正确")
	}

	token, err := util.GenerateToken(user.ID, req.UserName, 0)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	userResp := &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	userTokenResp := ctl.TokenData{
		User:         userResp,
		AccessToken:  token,
		RefreshToken: token, // TODO 加上 refresh token
	}

	return ctl.RespSuccessWithData(userTokenResp), nil
}

// UserInfoUpdate 用户修改信息
func (s *UserSrv) UserInfoUpdate(ctx context.Context, uId uint, req *types.UserInfoUpdateReq) (resp interface{}, err error) {
	var user *model.User
	// 找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	if req.NickName != "" {
		user.NickName = req.NickName
	}

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	return ctl.RespSuccess(), nil
}

// UserAvatarUpload 更新头像
func (s *UserSrv) UserAvatarUpload(ctx context.Context, uId uint, file multipart.File, fileSize int64, req *types.UserServiceReq) (resp interface{}, err error) {

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	var path string
	if conf.UploadModel == consts.UploadModelLocal { // 兼容两种存储方式
		path, err = util.UploadAvatarToLocalStatic(file, uId, user.UserName)
	} else {
		path, err = util.UploadToQiNiu(file, fileSize)
	}
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	return ctl.RespSuccess(), nil
}

// SendEmail 发送邮件
func (s *UserSrv) SendEmail(ctx context.Context, id uint, req *types.SendEmailServiceReq) (resp interface{}, err error) {
	var address string
	token, err := util.GenerateEmailToken(id, req.OperationType, req.Email, req.Password)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	address = conf.ValidEmail + token
	mailText := fmt.Sprintf(consts.EmailOperationMap[req.OperationType], address)
	if err = email.SendEmail(mailText, req.Email); err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

// Valid 验证内容
func (s *UserSrv) Valid(ctx context.Context, token string, req *types.ValidEmailServiceReq) (resp interface{}, err error) {
	var userId uint
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
			util.LogrusObj.Error(err)
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.SUCCESS {
		return ctl.RespSuccess(code), nil
	}

	// 获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ErrorDatabase
		return
	}

	switch operationType {
	case consts.EmailOperationBinding:
		user.Email = email
	case consts.EmailOperationNoBinding:
		user.Email = ""
	case consts.EmailOperationUpdatePassword:
		err = user.SetPassword(password)
		if err != nil {
			return
		}
	default:
		return nil, errors.New("操作不符合")
	}

	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		return nil, err
	}

	userResp := &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	// 成功则返回用户的信息
	return ctl.RespSuccessWithData(userResp), err
}
