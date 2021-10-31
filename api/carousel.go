package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//创建轮播图
func CreateCarousel(c *gin.Context) {
	service := service.CreateCarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// 轮播图列表接口
func ListCarousels(c *gin.Context) {
	service := service.ListCarouselsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
