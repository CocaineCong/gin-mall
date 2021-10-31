package api

import (
	"FanOneMall/pkg/logging"
	service2 "FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//创建收藏
func CreateFavorite(c *gin.Context) {
	service := service2.CreateFavoritesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service2.ShowFavoritesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	service := service2.DeleteFavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
