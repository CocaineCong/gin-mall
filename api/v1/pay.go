package v1

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"mall/service"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}
