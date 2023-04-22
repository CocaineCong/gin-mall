package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	util "mall/pkg/utils"
	"mall/service"
	"mall/types"
)

func ListCarouselsHandler() gin.HandlerFunc {
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
