package serializer

import "FanOneMall/model"

type Admin struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

//序列化
func BuildAdmin(admin model.Admin) Admin {
	return Admin{
		ID:       admin.ID,
		UserName: admin.UserName,
		Avatar:   admin.AvatarURL(),
		CreateAt: admin.CreatedAt.Unix(),
	}
}
