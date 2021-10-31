package serializer

import "FanOneMall/model"

type Boss struct {
	ID           uint      `json:"id"`
	BossName     string    `json:"boss_name"`
	BossNickName string    `json:"boss_nick_name"`
	Email        string    `json:"email"`
	Status       string    `json:"status"`
	Avatar       string    `json:"avatar"`
	CreateAt     int64     `json:"create_at"`
	Product      []Product `json:"product"`
}

func BuildBoss(boss model.Boss) Boss {
	return Boss{
		ID:           boss.ID,
		BossName:     boss.UserName,
		BossNickName: boss.Nickname,
		Email:        boss.Email,
		Status:       boss.Status,
		Avatar:       boss.Avatar,
		CreateAt:     boss.CreatedAt.Unix(),
		Product:      nil,
	}
}
