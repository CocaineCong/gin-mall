package serializer

import "FanOneMall/model"

type ProductImg struct {
	ID         uint   `json:"id"`
	ProductID  uint   `json:"product_id"`
	ImgPath    string `json:"img_path"`
	CreateAt   int64  `json:"create_at"`
	BossID     int    `json:"boss_id"`
	BossName   string `json:"boss_name"`
	BossAvatar string `json:"boss_avatar"`
}

//序列化商品图片
func BuildImg(item model.ProductImg) ProductImg {
	return ProductImg{
		ID:         item.ID,
		ProductID:  item.ProductID,
		ImgPath:    item.ImgPath,
		CreateAt:   item.CreatedAt.Unix(),
		BossID:     item.BossID,
		BossName:   item.BossName,
		BossAvatar: item.BossAvatar,
	}
}

//序列化商品列表
func BuildImgs(items []model.ProductImg) (imgs []ProductImg) {
	for _, item := range items {
		img := BuildImg(item)
		imgs = append(imgs, img)
	}
	return imgs
}

//序列化商品详情图片
func BuildInfoImg(item model.ProductInfoImg) ProductImg {
	return ProductImg{
		ID:        item.ID,
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

//序列化商品详情图片列表
func BuildInfoImgs(items []model.ProductInfoImg) (imgs []ProductImg) {
	for _, item := range items {
		img := BuildInfoImg(item)
		imgs = append(imgs, img)
	}
	return imgs
}

//序列化商品参数图片
func BuildParamImg(item model.ProductParamImg) ProductImg {
	return ProductImg{
		ID:        item.ID,
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

//序列化商品参数图片列表
func BuildParamImgs(items []model.ProductParamImg) (imgs []ProductImg) {
	for _, item := range items {
		img := BuildParamImg(item)
		imgs = append(imgs, img)
	}
	return imgs
}
