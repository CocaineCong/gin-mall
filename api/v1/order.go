package v1

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	util "mall/pkg/utils"
	"mall/service"
)

func CreateOrder(c *gin.Context) {
	createOrderService := service.CreateOrderService{}
	claim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createOrderService);err==nil{
		res:= createOrderService.Create(claim.ID)
		c.JSON(200,res)
	}else {
		c.JSON(400,ErrorResponse(err))
		logging.Info(err)
	}
}

func ListOrders(c *gin.Context) {
	listOrdersService := service.ListOrdersService{}
	claim,_:=util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrdersService);err==nil{
		res:= listOrdersService.List(claim.ID)
		c.JSON(200,res)
	}else {
		c.JSON(400,ErrorResponse(err))
		logging.Info(err)
	}
}

//订单详情
func ShowOrder(c *gin.Context) {
	showOrderService := service.ShowOrderService{}
	if err:= c.ShouldBind(&showOrderService);err==nil{
		res:= showOrderService.Show(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		logging.Info(err)
	}
}

func DeleteOrder(c *gin.Context) {
	deleteOrderService := service.DeleteOrderService{}
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}

