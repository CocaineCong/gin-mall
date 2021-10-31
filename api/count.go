package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

func ShowCount(c *gin.Context) {
	service := service.ShowCountService{}
	res:= service.Show(c.Param("id"))
	c.JSON(200,res)
}