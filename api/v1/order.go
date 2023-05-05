package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/consts"
	"mall/pkg/utils/log"
	"mall/service"
	"mall/types"
)

func CreateOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderCreateReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetOrderSrv()
			resp, err := l.OrderCreate(ctx.Request.Context(), &req)
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

func ListOrdersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderListReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.PageSize == 0 {
				req.PageSize = consts.BasePageSize
			}

			l := service.GetOrderSrv()
			resp, err := l.OrderList(ctx.Request.Context(), &req)
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

// 订单详情
func ShowOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderShowReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetOrderSrv()
			resp, err := l.OrderShow(ctx.Request.Context(), &req)
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

func DeleteOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderDeleteReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetOrderSrv()
			resp, err := l.OrderDelete(ctx.Request.Context(), &req)
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
