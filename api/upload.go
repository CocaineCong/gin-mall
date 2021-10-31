package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//上传授权
func UploadToken(c *gin.Context) {
	uploadAvatarService := service.UploadAvatarService{}
	if err := c.ShouldBind(&uploadAvatarService); err == nil {
		res := uploadAvatarService.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
