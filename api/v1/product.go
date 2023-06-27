package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/gin-mall/consts"
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/service"
	"github.com/CocaineCong/gin-mall/types"
)

// CreateProductHandler 创建商品
func CreateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		form, _ := ctx.MultipartForm()
		files := form.File["image"]
		l := service.GetProductSrv()
		resp, err := l.ProductCreate(ctx.Request.Context(), files, &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ListProductsHandler 商品列表
func ListProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductListReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = consts.BaseProductPageSize
		}

		l := service.GetProductSrv()
		resp, err := l.ProductList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// ShowProductHandler 商品详情
func ShowProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductShow(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// DeleteProductHandler 删除商品
func DeleteProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductDelete(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// UpdateProductHandler 更新商品
func UpdateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductUpdateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductUpdate(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// SearchProductsHandler 搜索商品
func SearchProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductSearchReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = consts.BasePageSize
		}

		l := service.GetProductSrv()
		resp, err := l.ProductSearch(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func ListProductImgHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListProductImgReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.ID == 0 {
			err := errors.New("参数错误,id不能为空")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetProductSrv()
		resp, err := l.ProductImgList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
