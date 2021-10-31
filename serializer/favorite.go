package serializer

import "FanOneMall/model"

type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num 		  int 	 `json:"num"`
	OnSale 		  bool `json:"on_sale"`
}

//序列化收藏夹
func BuildFavorite(item1 model.Favorite, item2 model.Product, item3 model.User) Favorite {
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
		Num :item2.Num,
		OnSale :item2.OnSale,
	}
}

//收藏夹列表
func BuildFavorites(items []model.Favorite) (favorites []Favorite) {
	for _, item1 := range items {
		item2 := model.Product{}
		item3 := model.User{}
		err := model.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item1, item2, item3)
		favorites = append(favorites, favorite)
	}
	return favorites
}
