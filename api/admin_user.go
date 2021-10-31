package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
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
