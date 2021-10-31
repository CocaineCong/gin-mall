package serializer

import "FanOneMall/model"

//分类序列化器
type Category struct {
	ID           uint   `json:"id"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

//序列化分类
func BuildCategory(item model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

//序列化分类列表
func BuildCategories(items []model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
