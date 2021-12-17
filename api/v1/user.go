package v1

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	util "mall/pkg/utils"
	"mall/service"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	var userLoginService service.UserLoginService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserUpdateService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.Update(claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	uploadAvatarService := service.UploadAvatarService{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatarService); err == nil {
		res := uploadAvatarService.Post(chaim.ID, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmailService service.SendEmailService
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmailService); err == nil {
		res := sendEmailService.Send(chaim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}

func ValidEmail(c *gin.Context) {
	var vaildEmailService service.VaildEmailService
	if err := c.ShouldBind(vaildEmailService); err == nil {
		res := vaildEmailService.Valid(c.GetHeader("Authorization"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}