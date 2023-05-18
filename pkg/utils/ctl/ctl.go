package ctl

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/gin-mall/consts"
	"github.com/CocaineCong/gin-mall/pkg/e"
)

// Response 基础序列化器
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Error   string      `json:"error"`
	TrackId string      `json:"track_id"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackId string `json:"track_id"`
}

// RespSuccess 带data成功返回
func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	trackId, _ := getTrackIdFromCtx(ctx)
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	if data == nil {
		data = "操作成功"
	}

	r := &Response{
		Status:  status,
		Data:    data,
		Msg:     e.GetMsg(status),
		TrackId: trackId,
	}

	return r
}

// RespError 错误返回
func RespError(ctx *gin.Context, err error, data string, code ...int) *TrackedErrorResponse {
	trackId, _ := getTrackIdFromCtx(ctx)
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
		TrackId: trackId,
	}

	return r
}

func getTrackIdFromCtx(ctx *gin.Context) (trackId string, err error) {
	spanCtxInterface, _ := ctx.Get(consts.SpanCTX)
	str := fmt.Sprintf("%v", spanCtxInterface)
	re := regexp.MustCompile(`([0-9a-fA-F]{16})`)

	match := re.FindStringSubmatch(str)
	if len(match) > 0 {
		return match[1], nil
	}
	return "", errors.New("获取 track id 错误")
}
