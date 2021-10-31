package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	service := service.CreateOrderService{}
	if err := c.ShouldBind(&service);err==nil{
		res:=service.Create()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

func ListOrders(c *gin.Context) {
	service := service.ListOrdersService{}
	if err := c.ShouldBind(&service);err==nil{
		res:=service.List(c.Param("id"))
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

//订单详情
func ShowOrder(c *gin.Context) {
	service := service.ShowOrderService{}
	if err:= c.ShouldBind(&service);err==nil{
		res:=service.Show(c.Param("num"))
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

func DeleteOrder(c *gin.Context) {
	service := service.DeleteOrderService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

