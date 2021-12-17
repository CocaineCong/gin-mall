package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	util "mall/pkg/utils"
	"mall/service"
)


// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	fmt.Println("c.Request.MultipartForm",form)
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.CreateProductService{}
	//c.SaveUploadedFile()
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
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
		c.JSON(400, ErrorResponse(err))
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
		res := updateProductService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}



//搜索商品
func SearchProducts(c *gin.Context) {
	searchProductsService := service.SearchProductsService{}
	if err := c.ShouldBind(&searchProductsService); err == nil {
		res := searchProductsService.Search()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}

func ListProductImg(c *gin.Context) {
	var listProductImgService service.ListProductImgService
	if err := c.ShouldBind(&listProductImgService); err == nil {
		res := listProductImgService.List(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}
