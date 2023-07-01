package types

type SkillProductImportReq struct {
}

type SkillProductReq struct {
	SkillProductId uint   `json:"skill_product_id" form:"skill_product_id"`
	ProductId      uint   `json:"product_id" form:"product_id"`
	BossId         uint   `json:"boss_id" form:"boss_id"`
	AddressId      uint   `json:"address_id" form:"address_id"`
	Key            string `json:"key" form:"key"`
}

type ListSkillProductReq struct {
	PageSize int64 `json:"page_size" form:"page_size"`
	PageNum  int64 `json:"page_num" form:"page_num"`
}

type GetSkillProductReq struct {
	ProductId uint `json:"product_id" form:"product_id"`
}
