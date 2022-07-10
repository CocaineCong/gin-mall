package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	"mall/service"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	//c.SaveUploadedFile()
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//商品列表
func ListProducts(c *gin.Context) {
	listProductsService := service.ProductService{}
	if err := c.ShouldBind(&listProductsService); err == nil {
		res := listProductsService.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//商品详情
func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}
	res := showProductService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

//删除商品
func DeleteProduct(c *gin.Context) {
	deleteProductService := service.ProductService{}
	res := deleteProductService.Delete(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

//更新商品
func UpdateProduct(c *gin.Context) {
	updateProductService := service.ProductService{}
	if err := c.ShouldBind(&updateProductService); err == nil {
		res := updateProductService.Update(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//搜索商品
func SearchProducts(c *gin.Context) {
	searchProductsService := service.ProductService{}
	if err := c.ShouldBind(&searchProductsService); err == nil {
		res := searchProductsService.Search(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ListProductImg(c *gin.Context) {
	var listProductImgService service.ListProductImgService
	if err := c.ShouldBind(&listProductImgService); err == nil {
		res := listProductImgService.List(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
