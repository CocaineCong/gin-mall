package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//初始化支付
func InitPay(c *gin.Context) {
	service := service.InitPayService{}
	if err:=c.ShouldBind(&service);err==nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

//接受FM支付回调接口
func ConfirmPay(c *gin.Context) {
	service := service.ConfirmPayService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Confirm()
		c.String(200, "success")
	} else {
		c.String(200, "success")
		logging.Info(err)
	}
}

func OrderPay(c *gin.Context) {
	service := service.OrderPay{}
	if err := c.ShouldBind(&service);err==nil {
		res := service.PayDowm()
		c.JSON(200,res)
	}else{
		logging.Info(err)
		c.JSON(200,ErrorResponse(err))
	}
}