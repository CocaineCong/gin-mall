package types

type SkillProductImportReq struct {
}

type SkillProductServiceReq struct {
	SkillProductId uint   `json:"skill_product_id" form:"skill_product_id"`
	ProductId      uint   `json:"product_id" form:"product_id"`
	BossId         uint   `json:"boss_id" form:"boss_id"`
	AddressId      uint   `json:"address_id" form:"address_id"`
	Key            string `json:"key" form:"key"`
}
