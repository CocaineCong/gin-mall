package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	api "mall/api/v1"
	"mall/middleware"
)

// 路由配置
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
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("product/:id", api.ShowProduct)
		v1.POST("products", api.SearchProducts)
		v1.GET("imgs/:id", api.ListProductImg)          // 商品图片
		v1.GET("categories", api.ListCategoryHandler()) // 商品分类
		v1.GET("carousels", api.ListCarouselsHandler()) // 轮播图

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{

			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)
			authed.POST("avatar", api.UploadAvatar) // 上传头像

			// 商品操作
			authed.POST("product", api.CreateProduct)
			authed.PUT("product/:id", api.UpdateProduct)
			authed.DELETE("product/:id", api.DeleteProduct)
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
			authed.POST("paydown", api.OrderPay)

			// 显示金额
			authed.POST("money", api.ShowMoneyHandler())

			// 秒杀专场
			authed.POST("import_skill_goods", api.ImportSkillGoods)
			authed.POST("init_skill_goods", api.InitSkillGoods)
			authed.POST("skill_goods", api.SkillGoods)
		}
	}
	return r
}
