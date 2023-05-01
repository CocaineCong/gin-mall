package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/consts"
	"mall/pkg/utils/log"
	"mall/service"
	"mall/types"
)

func CreateCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartCreateReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetCartSrv()
			resp, err := l.CartCreate(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 购物车详细信息
func ListCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartListReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.PageSize == 0 {
				req.PageSize = consts.BasePageSize
			}
			l := service.GetCartSrv()
			resp, err := l.CartList(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// UpdateCartHandler 修改购物车信息
func UpdateCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateCartServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetCartSrv()
			resp, err := l.CartUpdate(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 删除购物车
func DeleteCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartDeleteReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetCartSrv()
			resp, err := l.CartDelete(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
