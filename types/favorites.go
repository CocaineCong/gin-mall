package types

type FavoritesServiceReq struct {
	ProductId  uint `form:"product_id" json:"product_id"`
	BossId     uint `form:"boss_id" json:"boss_id"`
	FavoriteId uint `form:"favorite_id" json:"favorite_id"`
	PageNum    int  `form:"pageNum"`
	PageSize   int  `form:"pageSize"`
}

type FavoriteCreateReq struct {
	ProductId  uint `form:"product_id" json:"product_id"`
	BossId     uint `form:"boss_id" json:"boss_id"`
	FavoriteId uint `form:"favorite_id" json:"favorite_id"`
	PageNum    int  `form:"pageNum"`
	PageSize   int  `form:"pageSize"`
}

type FavoriteDeleteReq struct {
	Id uint `form:"id" json:"id"`
}

type FavoriteListResp struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	CategoryName  string `json:"category_name"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

//
// // 序列化收藏夹
// func BuildFavorite(item1 *model2.Favorite, item2 *model2.Product, item3 *model2.User) Favorite {
// 	return Favorite{
// 		UserID:        item1.UserID,
// 		ProductID:     item1.ProductID,
// 		CreatedAt:     item1.CreatedAt.Unix(),
// 		Name:          item2.Name,
// 		CategoryID:    item2.CategoryID,
// 		Title:         item2.Title,
// 		Info:          item2.Info,
// 		ImgPath:       item2.ImgPath,
// 		Price:         item2.Price,
// 		DiscountPrice: item2.DiscountPrice,
// 		BossID:        item3.ID,
// 		Num:           item2.Num,
// 		OnSale:        item2.OnSale,
// 	}
// }
//
// // 收藏夹列表
// func BuildFavorites(ctx context.Context, items []*model2.Favorite) (favorites []Favorite) {
// 	productDao := dao2.NewProductDao(ctx)
// 	bossDao := dao2.NewUserDao(ctx)
//
// 	for _, fav := range items {
// 		product, err := productDao.GetProductById(fav.ProductID)
// 		if err != nil {
// 			continue
// 		}
// 		boss, err := bossDao.GetUserById(fav.UserID)
// 		if err != nil {
// 			continue
// 		}
// 		favorite := BuildFavorite(fav, product, boss)
// 		favorites = append(favorites, favorite)
// 	}
// 	return favorites
// }
