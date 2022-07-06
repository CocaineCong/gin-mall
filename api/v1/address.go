package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	service2 "mall/service"
)

// CreateAddress 新增收货地址
func CreateAddress(c *gin.Context) {
	service := service2.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowAddresses 展示收货地址
func ShowAddresses(c *gin.Context) {
	service := service2.AddressService{}
	res := service.Show(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

// UpdateAddress 修改收货地址
func UpdateAddress(c *gin.Context) {
	service := service2.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteAddress 删除收获地址
func DeleteAddress(c *gin.Context) {
	service := service2.AddressService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
