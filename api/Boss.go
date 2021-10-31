package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//店家注册接口
func BossRegister(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	BossID := session.Get("BossID")
	var service service.BossRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register(BossID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//店家登陆接口
func BossLogin(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	bossID := session.Get("BossID")
	var service service.BossLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(bossID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//店家修改信息
func BossUpdate(c *gin.Context) {
	var service service.BossUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
