package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	"mall/service"
)

func CreateCategory(c *gin.Context) {
        service := service.CreateCategoryService{}
        if err := c.ShouldBind(&service); err == nil {
                res := service.Create(c.Request.Context())
                c.JSON(200, res)
        } else {
                c.JSON(200, ErrorResponse(err))
                util.LogrusObj.Infoln(err)
        }
}

func ListCategories(c *gin.Context) {
	listCategoriesService := service.ListCategoriesService{}
	if err := c.ShouldBind(&listCategoriesService); err == nil {
		res := listCategoriesService.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
