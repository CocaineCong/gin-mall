package ctl

import (
	"context"
)

type UserInfo struct {
	Id uint `json:"id"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	userId := ctx.Value("user_id").(uint)
	return &UserInfo{Id: userId}, nil
}
