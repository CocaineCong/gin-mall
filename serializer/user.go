package serializer

import "FanOneMall/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	Monery   int `json:"monery"`
	CreateAt int64  `json:"create_at"`
}

//BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.Nickname,
		Type:     user.Type,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		Monery: user.Monery,
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
