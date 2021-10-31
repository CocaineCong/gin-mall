package serializer

import "FanOneMall/model"

//公告序列化
type Notice struct {
	ID       uint   `json:"id"`
	Text     string `json:"text"`
	CreateAt int64  `json:"create_at"`
}

//序列化
func BuildNotice(item model.Notice) Notice {
	return Notice{
		ID:       item.ID,
		Text:     item.Text,
		CreateAt: item.CreatedAt.Unix(),
	}
}
