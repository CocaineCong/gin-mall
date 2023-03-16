package serializer

import "mall/model"

type Category struct {
	ID           uint   `json:"id"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
