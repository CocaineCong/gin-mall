package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"sync"

	conf "github.com/CocaineCong/gin-mall/config"
	"github.com/CocaineCong/gin-mall/consts"
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/CocaineCong/gin-mall/pkg/utils/email"
	"github.com/CocaineCong/gin-mall/pkg/utils/jwt"
	"github.com/CocaineCong/gin-mall/pkg/utils/log"
	util "github.com/CocaineCong/gin-mall/pkg/utils/upload"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/repository/db/model"
	"github.com/CocaineCong/gin-mall/types"
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
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("用户已经存在了")
		return
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
		Money:    consts.UserInitMoney,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		log.LogrusObj.Error(err)
		return
	}
	// 加密money
	money, err := user.EncryptMoney(req.Key)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	user.Money = money
	// 默认头像走的local
	user.Avatar = consts.UserDefaultAvatarLocal
	if conf.Config.System.UploadModel == consts.UploadModelOss {
		// 如果配置是走oss，则用url作为默认头像
		user.Avatar = consts.UserDefaultAvatarOss
	}

	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

// UserLogin 用户登陆函数
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if !exist { // 如果查询不到，返回相应的错误
		log.LogrusObj.Error(err)
		return nil, errors.New("用户不存在")
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("账号/密码不正确")
	}

	accessToken, refreshToken, err := jwt.GenerateToken(user.ID, req.UserName)
	if err != nil {
		log.LogrusObj.Error(err)
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

	resp = &types.UserTokenData{
		User:         userResp,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

// UserInfoUpdate 用户修改信息
func (s *UserSrv) UserInfoUpdate(ctx context.Context, req *types.UserInfoUpdateReq) (resp interface{}, err error) {
	// 找到用户
	u, _ := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	if req.NickName != "" {
		user.NickName = req.NickName
	}

	err = userDao.UpdateUserById(u.Id, user)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	return
}

// UserAvatarUpload 更新头像
func (s *UserSrv) UserAvatarUpload(ctx context.Context, file multipart.File, fileSize int64, req *types.UserServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	var path string
	if conf.Config.System.UploadModel == consts.UploadModelLocal { // 兼容两种存储方式
		path, err = util.AvatarUploadToLocalStatic(file, uId, user.UserName)
	} else {
		path, err = util.UploadToQiNiu(file, fileSize)
	}
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	return
}

// SendEmail 发送邮件
func (s *UserSrv) SendEmail(ctx context.Context, req *types.SendEmailServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	var address string
	token, err := jwt.GenerateEmailToken(u.Id, req.OperationType, req.Email, req.Password)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	sender := email.NewEmailSender()
	address = conf.Config.Email.ValidEmail + token
	mailText := fmt.Sprintf(consts.EmailOperationMap[req.OperationType], address)
	if err = sender.Send(mailText, req.Email, "FanOneMall"); err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

// Valid 验证内容
func (s *UserSrv) Valid(ctx context.Context, req *types.ValidEmailServiceReq) (resp interface{}, err error) {
	var userId uint
	var email string
	var password string
	var operationType uint
	// 验证token
	if req.Token == "" {
		err = errors.New("token不存在")
		log.LogrusObj.Error(err)
		return
	}
	claims, err := jwt.ParseEmailToken(req.Token)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	} else {
		userId = claims.UserID
		email = claims.Email
		password = claims.Password
		operationType = claims.OperationType
	}

	// 获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		log.LogrusObj.Error(err)
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
			err = errors.New("密码加密错误")
			return
		}
	default:
		return nil, errors.New("操作不符合")
	}

	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	resp = &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	return
}

// UserInfoShow 用户信息展示
func (s *UserSrv) UserInfoShow(ctx context.Context, req *types.UserInfoShowReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	user, err := dao.NewUserDao(ctx).GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	resp = &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	return
}

func (s *UserSrv) UserFollow(ctx context.Context, req *types.UserFollowingReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewUserDao(ctx).FollowUser(u.Id, req.Id)

	return
}

func (s *UserSrv) UserUnFollow(ctx context.Context, req *types.UserUnFollowingReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewUserDao(ctx).UnFollowUser(u.Id, req.Id)

	return
}
