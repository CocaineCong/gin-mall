package serializer

import (
	"mall/conf"
	"mall/consts"
	"mall/repository/db/model"
)

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	u := &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.PhotoHost + conf.HttpPort + conf.AvatarPath + user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	if conf.UploadModel == consts.UploadModelOss {
		u.Avatar = user.Avatar
	}

	return u
}

func BuildUsers(items []*model.User) (users []*User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
