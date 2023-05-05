package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/pkg/utils/log"
	"mall/service"
	"mall/types"
)

func OrderPaymentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.PaymentDownReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetPaymentSrv()
			resp, err := l.PayDown(ctx.Request.Context(), &req)
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
