package api

import (
	service2 "FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//新增收货地址
func CreateAddress(c *gin.Context) {
	service := service2.CreateAddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//展示收货地址
func ShowAddresses(c *gin.Context) {
	service := service2.ShowAddressService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

//修改收货地址
func UpdateAddress(c *gin.Context) {
	service := service2.UpdateAddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//删除收获地址
func DeleteAddress(c *gin.Context) {
	service := service2.DeleteAddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
