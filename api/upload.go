package api

import (
	"FanOneMall/pkg/util"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

//上传授权
func UploadToken(c *gin.Context) {
	file , fileHeader ,_ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	uploadAvatarService := service.UploadAvatarService{}
	chaim,_:= util.ParseToken("Authorization")
	if err := c.ShouldBind(&uploadAvatarService); err == nil {
		res := uploadAvatarService.Post(chaim.Id,file,fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
