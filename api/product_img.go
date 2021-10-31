package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//创建商品图片
func CreateProductImg(c *gin.Context) {
	service := service.CreateImgServe{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品详情接口
func ShowProductImgs(c *gin.Context) {
	service := service.ShowProductService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

//创建商品详情图片接口
func CreateInfoImg(c *gin.Context) {
	service := service.CreateInfoImgService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品详情图片接口
func ShowInfoImgs(c *gin.Context) {
	service := service.ShowInfoImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

//商品参数图片
func CreateParamImg(c *gin.Context) {
	service := service.CreateParamImgService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品参数图片接口
func ShowParamImgs(c *gin.Context) {
	service := service.ShowParamImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
