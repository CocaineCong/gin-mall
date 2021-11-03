package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//创建商品图片
func CreateProductImg(c *gin.Context) {
	createImgServe := service.CreateImgServe{}
	if err := c.ShouldBind(&createImgServe); err == nil {
		res := createImgServe.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品详情接口
func ShowProductImgs(c *gin.Context) {
	showProductService := service.ShowProductService{}
	res := showProductService.Show(c.Param("id"))
	c.JSON(200, res)
}

//创建商品详情图片接口
func CreateInfoImg(c *gin.Context) {
	createInfoImgService := service.CreateInfoImgService{}
	if err := c.ShouldBind(&createInfoImgService); err == nil {
		res := createInfoImgService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品详情图片接口
func ShowInfoImgs(c *gin.Context) {
	showInfoImgsService := service.ShowInfoImgsService{}
	res := showInfoImgsService.Show(c.Param("id"))
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
