package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/consts"
	"mall/pkg/e"
	"mall/pkg/utils/log"
	"mall/service"
	"mall/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserRegister(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
		}
	}
}

// UserLoginHandler 用户登陆接口
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserLogin(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
		}
	}
}

func UserUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoUpdateReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.UserInfoUpdate(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
		}
	}
}

func UploadAvatarHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			file, fileHeader, _ := ctx.Request.FormFile("file")
			if fileHeader == nil {
				err := errors.New(e.GetMsg(e.ErrorUploadFile))
				ctx.JSON(consts.IlleageRequest, ErrorResponse(ctx, err))
				log.LogrusObj.Infoln(err)
				return
			}
			fileSize := fileHeader.Size
			l := service.GetUserSrv()
			resp, err := l.UserAvatarUpload(ctx.Request.Context(), file, fileSize, &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
		}
	}
}

func SendEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SendEmailServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.SendEmail(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
		}
	}
}

func ValidEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ValidEmailServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.Valid(ctx.Request.Context(), &req)
			if err != nil {
				log.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
		}
	}
}
