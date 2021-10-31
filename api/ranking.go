package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//排行
func ListRanking(c *gin.Context) {
	listRankingService := service.ListRankingService{}
	res := listRankingService.List()
	c.JSON(200, res)
}

//家电排行
func ListElecRanking(c *gin.Context) {
	listElecRankingService := service.ListElecRankingService{}
	res := listElecRankingService.List()
	c.JSON(200, res)
}

//配件排行
func ListAcceRanking(c *gin.Context) {
	listAcceRankingService := service.ListAcceRankingService{}
	res := listAcceRankingService.List()
	c.JSON(200, res)
}
