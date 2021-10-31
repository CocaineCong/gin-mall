package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	service := service.CreateCartService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//购物车详细信息
func ShowCarts(c *gin.Context) {
	service := service.ShowCartsService{}
	res := service.Show(c.Param("id"))
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
