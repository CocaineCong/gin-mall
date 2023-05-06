package ctl

import (
	"context"

	"github.com/gin-gonic/gin"

	"mall/middleware"
	"mall/pkg/e"
)

// Response 基础序列化器
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Error   string      `json:"error"`
	TrackId string      `json:"track_id"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackId string `json:"track_id"`
}

// RespSuccess 成功返回
func RespSuccess(ctx *gin.Context, code ...int) *Response {
	spanCtxInterface, _ := ctx.Get(middleware.SpanCTX)
	var spanCtx context.Context
	spanCtx = spanCtxInterface.(context.Context)
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status:  status,
		Data:    "操作成功",
		Msg:     e.GetMsg(status),
		TrackId: spanCtx.Value(middleware.SpanCTX).(string),
	}

	return r
}

// RespSuccessWithData 带data成功返回
func RespSuccessWithData(ctx *gin.Context, data interface{}, code ...int) *Response {
	spanCtxInterface, _ := ctx.Get(middleware.SpanCTX)
	var spanCtx context.Context
	spanCtx = spanCtxInterface.(context.Context)
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status:  status,
		Data:    data,
		Msg:     e.GetMsg(status),
		TrackId: spanCtx.Value(middleware.SpanCTX).(string),
	}

	return r
}

// RespError 错误返回
func RespError(ctx *gin.Context, err error, data string, code ...int) *TrackedErrorResponse {
	spanCtxInterface, _ := ctx.Get(middleware.SpanCTX)
	var spanCtx context.Context
	spanCtx = spanCtxInterface.(context.Context)
	status := e.ERROR
	if code != nil {
		status = code[0]
	}

	r := &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Msg:    e.GetMsg(status),
			Data:   data,
			Error:  err.Error(),
		},
		TrackId: spanCtx.Value(middleware.SpanCTX).(string),
	}

	return r
}
