package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	service := service.ListUserService{}
	res := service.List()
	c.JSON(200, res)
}

func ListUsers(c *gin.Context) {
	service := service.ListUsersService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
