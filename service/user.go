package service

import (
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	logging "github.com/sirupsen/logrus"
	"FanOneMall/pkg/util"
	"FanOneMall/serializer"
	"github.com/jinzhu/gorm"
	"mime/multipart"
)

//UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname  string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

type UploadAvatarService struct {
}

//valid 验证表单 验证用户是否存在
func (service *UserRegisterService) Valid(userId, status interface{}) *serializer.Response {
	var code int
	count := 0
	err := model.DB.Model(&model.User{}).Where("nickname=?", service.Nickname).Count(&count).Error
	if err != nil {
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ErrorExistNick
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	count = 0
	err = model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).Count(&count).Error
	if err != nil {
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return nil
}

//Register 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}
	code := e.SUCCESS
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//UserLoginService 管理用户登陆的服务
type UserLoginService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Challenge string `form:"challenge" json:"challenge"`
	Validate  string `form:"validate" json:"validate"`
	Seccode   string `form:"seccode" json:"seccode"`
}

//Login 用户登陆函数
func (service *UserLoginService) Login(userID, status interface{}) serializer.Response {
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
	token, err := util.GenerateToken(service.UserName, service.Password, 0)
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

//用户修改信息的服务
type UserUpdateService struct {
	ID       uint   `form:"id" json:"id"`
	NickName string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Avatar   string `form:"avatar" json:"avatar"`
}

//Update 用户修改信息
func (service *UserUpdateService) Update() serializer.Response {
	var user model.User
	code := e.SUCCESS
	//找到用户
	err := model.DB.First(&user, service.ID).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Nickname = service.NickName
	user.UserName = service.UserName
	if service.Avatar != "" {
		user.Avatar = service.Avatar
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

func (service *UploadAvatarService) Post(id string, file multipart.File,fileSize int64) serializer.Response {
	var user model.User
	code := e.SUCCESS
	status , info := util.UploadToQiNiu(file,fileSize)
	if status != 200 {
		return serializer.Response{
			Status:  status  ,
			Data:      e.GetMsg(status),
			Error:info,
		}
	}
	model.DB.Where("id=?",id).First(&user)
	user.Avatar = info
	model.DB.Save(&user)
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}