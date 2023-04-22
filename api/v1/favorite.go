package v1

import (
	"net/http"

	util "mall/pkg/utils"
	"mall/service"
	"mall/types"

	"github.com/gin-gonic/gin"
)

// 创建收藏
func CreateFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetFavoriteSrv()
			userId := ctx.Keys["user_id"].(uint)
			resp, err := l.FavoriteCreate(ctx.Request.Context(), userId, &req)
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

// ListFavoritesHandler 收藏夹详情接口
func ListFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetFavoriteSrv()
			userId := ctx.Keys["user_id"].(uint)
			resp, err := l.FavoriteList(ctx.Request.Context(), userId, &req)
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

// DeleteFavoriteHandler 删除收藏夹
func DeleteFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetFavoriteSrv()
			resp, err := l.FavoriteDelete(ctx.Request.Context(), &req)
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
