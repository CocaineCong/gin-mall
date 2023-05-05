package v1

import (
	"net/http"

	"mall/consts"
	"mall/pkg/utils/log"
	"mall/service"
	"mall/types"

	"github.com/gin-gonic/gin"
)

// 创建收藏
func CreateFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteCreateReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetFavoriteSrv()
			resp, err := l.FavoriteCreate(ctx.Request.Context(), &req)
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

// ListFavoritesHandler 收藏夹详情接口
func ListFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.PageSize == 0 {
				req.PageSize = consts.BasePageSize
			}
			l := service.GetFavoriteSrv()
			resp, err := l.FavoriteList(ctx.Request.Context(), &req)
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

// DeleteFavoriteHandler 删除收藏夹
func DeleteFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteDeleteReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetFavoriteSrv()
			resp, err := l.FavoriteDelete(ctx.Request.Context(), &req)
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
