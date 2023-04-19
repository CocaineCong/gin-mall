package serializer

import (
	"mall/conf"
	"mall/consts"
	"mall/repository/db/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        int    `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

// 序列化商品
func BuildProduct(item *model.Product) Product {
	p := Product{
		ID:            item.ID,
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          item.View(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		CreatedAt:     item.CreatedAt.Unix(),
		BossID:        int(item.BossID),
		BossName:      item.BossName,
		BossAvatar:    conf.PhotoHost + conf.HttpPort + conf.AvatarPath + item.BossAvatar,
	}

	if conf.UploadModel == consts.UploadModelOss {
		p.ImgPath = item.ImgPath
		p.BossAvatar = item.BossAvatar
	}

	return p

}

// 序列化商品列表
func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
