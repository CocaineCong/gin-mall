package types

type CartServiceReq struct {
	Id        uint `form:"id" json:"id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	UserId    uint `form:"user_id" json:"user_id"`
}
