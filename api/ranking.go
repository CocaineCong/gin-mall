package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//排行
func ListRanking(c *gin.Context) {
	service := service.ListRankingService{}
	res := service.List()
	c.JSON(200, res)
}

//家电排行
func ListElecRanking(c *gin.Context) {
	service := service.ListElecRankingService{}
	res := service.List()
	c.JSON(200, res)
}

//配件排行
func ListAcceRanking(c *gin.Context) {
	service := service.ListAcceRankingService{}
	res := service.List()
	c.JSON(200, res)
}
