package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "mall/pkg/utils"
	"mall/service"
	"mall/types"
)

func UserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.Register(ctx.Request.Context(), &req)
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

// UserLogin 用户登陆接口
func UserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.Login(ctx.Request.Context(), &req)
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

func UserUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetUserSrv()
			resp, err := l.Update(ctx.Request.Context(), userId, &req)
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

func UploadAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			file, fileHeader, _ := ctx.Request.FormFile("file")
			fileSize := fileHeader.Size
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetUserSrv()
			resp, err := l.Post(ctx.Request.Context(), userId, file, fileSize, &req)
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

func SendEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SendEmailServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetUserSrv()
			resp, err := l.Send(ctx.Request.Context(), userId, &req)
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

func ValidEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ValidEmailServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()
			resp, err := l.Valid(ctx.Request.Context(), ctx.Param("token"), &req)
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
