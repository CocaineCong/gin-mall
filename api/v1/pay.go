package v1

import ("github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)
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
