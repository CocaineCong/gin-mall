package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "mall/pkg/utils"
	"mall/service"
	"mall/types"
)

func CreateCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetCartSrv()
			resp, err := l.CartCreate(ctx.Request.Context(), userId, &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 购物车详细信息
func ListCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetCartSrv()
			resp, err := l.CartList(ctx.Request.Context(), userId)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// UpdateCartHandler 修改购物车信息
func UpdateCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			req.UserId = ctx.Keys["user_id"].(uint)
			l := service.GetCartSrv()
			resp, err := l.CartUpdate(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 删除购物车
func DeleteCartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CartServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			req.UserId = ctx.Keys["user_id"].(uint)
			l := service.GetCartSrv()
			resp, err := l.CartDelete(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
