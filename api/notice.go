package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//展示公告
func ShowNotice(c *gin.Context) {
	service := service.ShowNoticeService{}
	res := service.Show()
	c.JSON(200, res)
}

//创建公告
func CreateNotice(c *gin.Context) {
	service := service.CreateNoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//更新公告
func UpdateNotice(c *gin.Context) {
	service := service.UpdateNoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
