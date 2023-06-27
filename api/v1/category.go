package v1

import (
	"net/http"

	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/CocaineCong/gin-mall/service"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/types"
)

func ListCategoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListCategoryReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetCategorySrv()
		resp, err := l.CategoryList(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
