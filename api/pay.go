package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//初始化支付
func InitPay(c *gin.Context) {
	initPayService := service.InitPayService{}
	if err:=c.ShouldBind(&initPayService);err==nil {
		res := initPayService.Init()
		c.JSON(200, res)
	} else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

//接受FM支付回调接口
func ConfirmPay(c *gin.Context) {
	confirmPayService := service.ConfirmPayService{}
	if err := c.ShouldBind(&confirmPayService); err == nil {
		confirmPayService.Confirm()
		c.String(200, "success")
	} else {
		c.String(200, "success")
		logging.Info(err)
	}
}

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	if err := c.ShouldBind(&orderPay);err==nil {
		res := orderPay.PayDowm()
		c.JSON(200,res)
	}else{
		logging.Info(err)
		c.JSON(200,ErrorResponse(err))
	}
}