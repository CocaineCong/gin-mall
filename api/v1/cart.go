package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	"mall/service"
)

func CreateCart(c *gin.Context) {
	createCartService := service.CreateService{}
	claim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Param("id"),claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//购物车详细信息
func ShowCarts(c *gin.Context) {
	showCartsService := service.CreateService{}
	res := showCartsService.Show(c.Param("id"))
	c.JSON(200, res)
}

//修改购物车信息
func UpdateCart(c *gin.Context) {
	updateCartService := service.CreateService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//删除购物车
func DeleteCart(c *gin.Context) {
	deleteCartService := service.CreateService{}
	claim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Param("id"),claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
