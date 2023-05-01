package types

type CartServiceReq struct {
	Id        uint `form:"id" json:"id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	UserId    uint `form:"user_id" json:"user_id"`
}

type CartCreateReq struct {
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
}

type CartDeleteReq struct {
	Id uint `form:"id" json:"id"`
}

type UpdateCartServiceReq struct {
	Id  uint `form:"id" json:"id"`
	Num uint `form:"num" json:"num"`
}

type CartListReq struct {
	BasePage
}

// 购物车
type CartResp struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check_        bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	Info          string `json:"info"`
}

//
// func BuildCart(cart *model2.Cart, product *model2.Product, boss *model2.User) *CartResp {
// 	c := &CartResp{
// 		ID:            cart.ID,
// 		UserID:        cart.UserID,
// 		ProductID:     cart.ProductID,
// 		CreateAt:      cart.CreatedAt.Unix(),
// 		Num:           cart.Num,
// 		MaxNum:        cart.MaxNum,
// 		Check:         cart.Check,
// 		Name:          product.Name,
// 		ImgPath:       conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + product.ImgPath,
// 		DiscountPrice: product.DiscountPrice,
// 		BossId:        boss.ID,
// 		BossName:      boss.UserName,
// 		Desc:          product.Info,
// 	}
// 	if conf.UploadModel == consts.UploadModelOss {
// 		c.ImgPath = product.ImgPath
// 	}
//
// 	return c
// }
//
// func BuildCarts(items []*model2.Cart) (carts []*CartResp) {
// 	for _, item1 := range items {
// 		product, err := dao2.NewProductDao(context.Background()).
// 			GetProductById(item1.ProductID)
// 		if err != nil {
// 			continue
// 		}
// 		boss, err := dao2.NewUserDao(context.Background()).
// 			GetUserById(item1.BossID)
// 		if err != nil {
// 			continue
// 		}
// 		cart := BuildCart(item1, product, boss)
// 		carts = append(carts, cart)
// 	}
// 	return carts
// }
