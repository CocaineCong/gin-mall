package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)


func ListUsers(c *gin.Context) {
	listUsersService := service.ListUsersService{}
	if err := c.ShouldBind(&listUsersService); err == nil {
		res := listUsersService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
