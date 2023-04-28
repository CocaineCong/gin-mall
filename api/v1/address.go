package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	util "mall/pkg/utils/log"
	"mall/service"
	"mall/types"
)

// CreateAddressHandler 新增收货地址
func CreateAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressCreateReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetAddressSrv()
			userId := ctx.Keys["user_id"].(uint)
			resp, err := l.AddressCreate(ctx.Request.Context(), req, userId)
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

// GetAddressHandler 展示某个收货地址
func GetAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.AddressServiceReq)

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetAddressSrv()
			resp, err := l.Show(ctx.Request.Context(), ctx.Param("id"))
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

// ListAddressHandler 展示收货地址
func ListAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.AddressServiceReq)

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetAddressSrv()
			resp, err := l.List(ctx.Request.Context(), cast.ToUint(ctx.Param("id")))
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

// UpdateAddressHandler 修改收货地址
func UpdateAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetAddressSrv()
			resp, err := l.Update(ctx.Request.Context(), &req, userId, cast.ToUint(ctx.Param("id")))
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

// DeleteAddressHandler 删除收获地址
func DeleteAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetAddressSrv()
			resp, err := l.Delete(ctx.Request.Context(), cast.ToUint(ctx.Param("id")), userId)
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
