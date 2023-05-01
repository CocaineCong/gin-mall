package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	api "mall/api/v1"
	"mall/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		// 商品操作
		v1.GET("product_list", api.ListProductsHandler())
		v1.GET("product_show", api.ShowProductHandler())
		v1.POST("product_search", api.SearchProductsHandler())
		v1.GET("imgs_list", api.ListProductImgHandler()) // 商品图片
		v1.GET("categories", api.ListCategoryHandler())  // 商品分类
		v1.GET("carousels", api.ListCarouselsHandler())  // 轮播图

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.AuthMiddleware())
		{

			// 用户操作
			authed.POST("user_update", api.UserUpdateHandler())
			authed.POST("user/send_email", api.SendEmailHandler())
			authed.GET("user/valid_email", api.ValidEmailHandler())
			authed.POST("avatar", api.UploadAvatarHandler()) // 上传头像

			// 商品操作
			authed.POST("product_create", api.CreateProductHandler())
			authed.POST("product_update", api.UpdateProductHandler())
			authed.POST("product_delete", api.DeleteProductHandler())
			// 收藏夹
			authed.GET("favorites_list", api.ListFavoritesHandler())
			authed.POST("favorites_create", api.CreateFavoriteHandler())
			authed.POST("favorites_delete", api.DeleteFavoriteHandler())

			// 订单操作
			authed.POST("orders_create", api.CreateOrderHandler())
			authed.GET("orders_list", api.ListOrdersHandler())
			authed.GET("orders_show", api.ShowOrderHandler())
			authed.POST("orders_delete", api.DeleteOrderHandler())

			// 购物车
			authed.POST("carts_create", api.CreateCartHandler())
			authed.GET("carts_list", api.ListCartHandler())
			authed.POST("carts_update", api.UpdateCartHandler()) // 购物车id
			authed.POST("carts_delete", api.DeleteCartHandler())

			// 收获地址操作
			authed.POST("addresses_create", api.CreateAddressHandler())
			authed.GET("addresses_show", api.ShowAddressHandler())
			authed.GET("addresses_list", api.ListAddressHandler())
			authed.POST("addresses_update", api.UpdateAddressHandler())
			authed.POST("addresses_delete", api.DeleteAddressHandler())

			// 支付功能
			authed.POST("paydown", api.OrderPaymentHandler())

			// 显示金额
			authed.POST("money", api.ShowMoneyHandler())

			// 秒杀专场
			authed.POST("import_skill_goods", api.ImportSkillProductHandler())
			authed.POST("init_skill_goods", api.InitSkillProductHandler())
			authed.POST("skill_goods", api.SkillProductHandler())
		}
	}
	return r
}
