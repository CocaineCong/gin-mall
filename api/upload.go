package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//上传授权
func UploadToken(c *gin.Context) {
	service := service.UploadAvatarService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
