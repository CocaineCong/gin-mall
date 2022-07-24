package serializer

import (
	"context"
	"mall/dao"
	"mall/model"
)

type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

//序列化收藏夹
func BuildFavorite(item1 *model.Favorite, item2 *model.Product, item3 *model.User) Favorite {
	return Favorite{
		UserID:        item1.UserID,
		ProductID:     item1.ProductID,
		CreatedAt:     item1.CreatedAt.Unix(),
		Name:          item2.Name,
		CategoryID:    item2.CategoryID,
		Title:         item2.Title,
		Info:          item2.Info,
		ImgPath:       item2.ImgPath,
		Price:         item2.Price,
		DiscountPrice: item2.DiscountPrice,
		BossID:        item3.ID,
		Num:           item2.Num,
		OnSale:        item2.OnSale,
	}
}

// 收藏夹列表
func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []Favorite) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)

	for _, fav := range items {
		product, err := productDao.GetProductById(fav.ProductID)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(fav.UserID)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(fav, product, boss)
		favorites = append(favorites, favorite)
	}
	return favorites
}
