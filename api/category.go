package api

import (
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	createCategoryService := service.CreateCategoryService{}
	if err := c.ShouldBind(&createCategoryService); err == nil {
		res := createCategoryService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ListCategories 分类列表接口
func ListCategories(c *gin.Context) {
	listCategoriesService := service.ListCategoriesService{}
	if err := c.ShouldBind(&listCategoriesService); err == nil {
		res := listCategoriesService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
