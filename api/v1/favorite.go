package v1

import (
	"mall/consts"
	util "mall/pkg/utils"
	service2 "mall/service"

	"github.com/gin-gonic/gin"
)

// 创建收藏
func CreateFavorite(c *gin.Context) {
	service := service2.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service2.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	service := service2.FavoritesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
