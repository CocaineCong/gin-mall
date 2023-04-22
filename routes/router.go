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
		v1.GET("products", api.ListProductsHandler())
		v1.GET("product/:id", api.ShowProductHandler())
		v1.POST("products", api.SearchProductsHandler())
		v1.GET("imgs/:id", api.ListProductImgHandler()) // 商品图片
		v1.GET("categories", api.ListCategoryHandler()) // 商品分类
		v1.GET("carousels", api.ListCarouselsHandler()) // 轮播图

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{

			// 用户操作
			authed.PUT("user", api.UserUpdateHandler())
			authed.POST("user/sending-email", api.SendEmailHandler())
			authed.POST("user/valid-email", api.ValidEmailHandler())
			authed.POST("avatar", api.UploadAvatarHandler()) // 上传头像

			// 商品操作
			authed.POST("product", api.CreateProductHandler())
			authed.PUT("product/:id", api.UpdateProductHandler())
			authed.DELETE("product/:id", api.DeleteProductHandler())
			// 收藏夹
			authed.GET("favorites", api.ListFavoritesHandler())
			authed.POST("favorites", api.CreateFavoriteHandler())
			authed.DELETE("favorites/:id", api.DeleteFavoriteHandler())

			// 订单操作
			authed.POST("orders", api.CreateOrderHandler())
			authed.GET("orders", api.ListOrdersHandler())
			authed.GET("orders/:id", api.ShowOrderHandler())
			authed.DELETE("orders/:id", api.DeleteOrderHandler())

			// 购物车
			authed.POST("carts", api.CreateCartHandler())
			authed.GET("carts", api.ListCartHandler())
			authed.PUT("carts/:id", api.UpdateCartHandler()) // 购物车id
			authed.DELETE("carts/:id", api.DeleteCartHandler())

			// 收获地址操作
			authed.POST("addresses", api.CreateAddressHandler())
			authed.GET("addresses/:id", api.GetAddressHandler())
			authed.GET("addresses", api.ListAddressHandler())
			authed.PUT("addresses/:id", api.UpdateAddressHandler())
			authed.DELETE("addresses/:id", api.DeleteAddressHandler())

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
