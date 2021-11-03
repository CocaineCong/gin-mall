package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateCart(c *gin.Context) {
	createCartService := service.CreateCartService{}
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//购物车详细信息
func ShowCarts(c *gin.Context) {
	showCartsService := service.ShowCartsService{}
	res := showCartsService.Show(c.Param("id"))
	c.JSON(200, res)
}

//修改购物车信息
func UpdateCart(c *gin.Context) {
	service := service.UpdateCartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//删除购物车
func DeleteCart(c *gin.Context) {
	service := service.DeleteCartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
