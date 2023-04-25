package model

type SkillProduct struct {
	Id         uint `gorm:"primarykey"`
	ProductId  uint `gorm:"not null"`
	BossId     uint `gorm:"not null"`
	Title      string
	Money      float64
	Num        int `gorm:"not null"`
	CustomId   uint
	CustomName string
}

type SkillProduct2MQ struct {
	SkillProductId uint    `json:"skill_good_id"`
	ProductId      uint    `json:"product_id"`
	BossId         uint    `json:"boss_id"`
	UserId         uint    `json:"user_id"`
	Money          float64 `json:"money"`
	AddressId      uint    `json:"address_id"`
	Key            string  `json:"key"`
}
