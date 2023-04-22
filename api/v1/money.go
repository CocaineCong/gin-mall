package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "mall/pkg/utils"
	"mall/service"
	"mall/types"
)

func ShowMoneyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ShowMoneyServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetMoneySrv()
			resp, err := l.MoneyShow(ctx.Request.Context(), userId, &req)
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
