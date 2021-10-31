package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//注册接口
func AdminRegister(c *gin.Context) {
	var service service.AdminRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//登陆接口
func AdminLogin(c *gin.Context) {
	var service service.AdminLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
