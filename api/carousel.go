package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//创建轮播图
func CreateCarousel(c *gin.Context) {
	createCarouselService := service.CreateCarouselService{}
	if err := c.ShouldBind(&createCarouselService); err == nil {
		res := createCarouselService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// 轮播图列表接口
func ListCarousels(c *gin.Context) {
	listCarouselsService := service.ListCarouselsService{}
	if err := c.ShouldBind(&listCarouselsService); err == nil {
		res := listCarouselsService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
