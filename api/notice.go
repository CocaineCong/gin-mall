package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//展示公告
func ShowNotice(c *gin.Context) {
	showNoticeService := service.ShowNoticeService{}
	res := showNoticeService.Show()
	c.JSON(200, res)
}

//创建公告
func CreateNotice(c *gin.Context) {
	createNoticeService := service.CreateNoticeService{}
	if err := c.ShouldBind(&createNoticeService); err == nil {
		res := createNoticeService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//更新公告
func UpdateNotice(c *gin.Context) {
	updateNoticeService := service.UpdateNoticeService{}
	if err := c.ShouldBind(&updateNoticeService); err == nil {
		res := updateNoticeService.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
