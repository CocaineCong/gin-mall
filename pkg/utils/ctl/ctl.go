package ctl

import (
	"mall/pkg/e"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
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

// RespList 带有总数的列表构建器
func RespList(items interface{}, total int64) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

// RespSuccess 成功返回
func RespSuccess(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}

	return r
}

// RespSuccessWithData 带data成功返回
func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}

	return r
}

// RespError 错误返回
func RespError(err error, data string, code ...int) *TrackedErrorResponse {
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
		// TrackId:  // TODO 加上track id
	}

	return r
}
