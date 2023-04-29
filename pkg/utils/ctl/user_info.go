package ctl

import (
	"context"
	"errors"
	"fmt"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

func InitUserInfo(ctx context.Context) {
	fmt.Println("ctx", ctx)
}
