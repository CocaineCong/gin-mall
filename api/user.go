package api

import (
	"FanOneMall/serializer"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	var userLoginService service.UserLoginService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UserUpdate 用户修改信息
func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserUpdateService
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//CheckToken 用户详情
func CheckToken(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}

//SendEmail 发送邮件接口
func SendEmail(c *gin.Context) {
	//var sendEmailService service.SendEmailService
	//if err := c.ShouldBind(&sendEmailService); err == nil {
	//	res := sendEmailService.Send()
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(200, ErrorResponse(err))
	//	logging.Info(err)
	//}
}

//ValidEmail  绑定和解绑邮箱接口
func ValidEmail(c *gin.Context) {
	//var vaildEmailService service.VaildEmailService
	//if err := c.ShouldBind(&vaildEmailService); err == nil {
	//	res := vaildEmailService.Vaild()
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(200, ErrorResponse(err))
	//	logging.Info(err)
	//}
}
