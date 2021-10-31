package api

import (
	"FanOneMall/pkg/logging"
	"FanOneMall/service"
	"github.com/gin-gonic/gin"
)

//TODO

// 创建商品
func CreateProduct(c *gin.Context) {
	createProductService := service.CreateProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品列表
func ListProducts(c *gin.Context) {
	listProductsService := service.ListProductsService{}
	if err := c.ShouldBind(&listProductsService); err == nil {
		res := listProductsService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品详情
func ShowProduct(c *gin.Context) {
	showProductService := service.ShowProductService{}
	res := showProductService.Show(c.Param("id"))
	c.JSON(200, res)
}

//删除商品
func DeleteProduct(c *gin.Context) {
	deleteProductService := service.DeleteProductService{}
	res := deleteProductService.Delete(c.Param("id"))
	c.JSON(200, res)
}

//更新商品
func UpdateProduct(c *gin.Context) {
	updateProductService := service.UpdateProductService{}
	if err := c.ShouldBind(&updateProductService); err == nil {
		res := updateProductService.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func UpProduct(c *gin.Context) {
	upProductService := service.UpProductService{}
	if err := c.ShouldBind(&upProductService); err == nil {
		res := upProductService.UpProduct()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}


//搜索商品
func SearchProducts(c *gin.Context) {
	searchProductsService := service.SearchProductsService{}
	if err := c.ShouldBind(&searchProductsService); err == nil {
		res := searchProductsService.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
