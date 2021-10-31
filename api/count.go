package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

func ShowCount(c *gin.Context) {
	showCountService := service.ShowCountService{}
	res:= showCountService.Show(c.Param("id"))
	c.JSON(200,res)
}