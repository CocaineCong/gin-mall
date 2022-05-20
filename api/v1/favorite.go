package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	service2 "mall/service"
)

//创建收藏
func CreateFavorite(c *gin.Context) {
	service := service2.FavoritesService{}
	claim , _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service2.FavoritesService{}
	claim , _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	service := service2.FavoritesService{}
	claim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(claim.ID,c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
