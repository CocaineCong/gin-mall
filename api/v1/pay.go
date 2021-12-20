package v1

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	util "mall/pkg/utils"
	"mall/service"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	claim,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(claim.ID)
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}
