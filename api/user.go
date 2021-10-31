package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
	"FanOneMall/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UserRegister 用户注册
func UserRegister(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	userID := session.Get("userId")
	var service service.UserRegisterService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register(userID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	userID := session.Get("userId")
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(userID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UserUpdate 用户修改信息
func UserUpdate(c *gin.Context) {
	var service service.UserUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
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
	var service service.SendEmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Send()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//VaildEmail  绑定和解绑邮箱接口
func VaildEmail(c *gin.Context) {
	var service service.VaildEmailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
