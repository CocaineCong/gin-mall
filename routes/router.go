package routes

import (
	"FanOneMall/api"
	"FanOneMall/conf"
	"FanOneMall/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	middleware.HttpLogToFile(conf.AppMode)
	//r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/vaild-email", api.ValidEmail)

		//商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ShowProductImgs)     //商品图片
		v1.GET("info-imgs/:id", api.ShowInfoImgs)   //商品详情图片操作
		v1.GET("param-imgs/:id", api.ShowParamImgs) //商品参数图片操作
		v1.GET("categories", api.ListCategories)    //商品分类
		v1.GET("carousels", api.ListCarousels)      //轮播图

		//排行
		v1.GET("rankings", api.ListRanking)
		v1.GET("elec-rankings", api.ListElecRanking)
		v1.GET("acce-rankings", api.ListAcceRanking)

		//公告
		v1.GET("notices", api.ShowNotice)

		// 支付功能
		v1.POST("paydown",api.OrderPay)
		v1.GET("payments", api.ConfirmPay)

		//v1.POST("payments",api.InitPay)

		authed := v1.Group("/")            //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//验证token
			authed.GET("ping", api.CheckToken)

			// 用户操作
			authed.PUT("user", api.UserUpdate)                  //用户修改，更新
			authed.POST("user/sending-email", api.SendEmail)    //用户操作
			authed.POST("avatar",api.UploadToken) 		 		//上传头像

			// 商品操作
			authed.POST("products",api.CreateProduct)  		 //创建商品
			authed.DELETE("products/:id", api.DeleteProduct)
			authed.PUT("products_up", api.UpProduct)

			//收藏夹
			authed.GET("favorites/:id", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites", api.DeleteFavorite)

			//订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("user/:id/orders", api.ListOrders)
			authed.GET("orders/:num", api.ShowOrder)
			authed.DELETE("orders", api.DeleteOrder)

			//购物车
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.ShowCarts)
			authed.PUT("carts", api.UpdateCart)
			authed.DELETE("carts", api.DeleteCart)

			//收获地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ShowAddresses)
			authed.PUT("addresses", api.UpdateAddress)
			authed.DELETE("addresses", api.DeleteAddress)

			//数量操作
			authed.GET("counts/:id",api.ShowCount)

			//支付功能
			authed.POST("payments",api.InitPay)
			authed.POST("OrderPayment",api.OrderPay)
		}
	}
	v2 := r.Group("/api/v2")
	{
		//管理员登陆注册
		v2.POST("admin/register", api.AdminRegister)
		//登陆
		v2.POST("admin/login", api.AdminLogin)
		//商品操作
		v2.GET("products", api.ListProducts)
		v2.GET("products/:id", api.ShowProduct)
		v2.GET("imgs/:id", api.ShowProductImgs)		//商品图片
		//轮播图
		v2.GET("carousels", api.ListCarousels)
		//分类操作
		v2.GET("categories", api.ListCarousels)
		//用户操作
		v2.GET("users", api.ListUsers)
		authed2 := v2.Group("/")
		authed2.Use(middleware.JWTAdmin())
		{
			//商品操作
			authed2.POST("products", api.CreateProduct)
			//authed2.GET("users",api.ListUsers)
			authed2.DELETE("products/:id", api.DeleteProduct)
			authed2.PUT("products", api.UpdateProduct)
			//轮播图操作
			authed2.POST("carousels", api.CreateCarousel)
			//商品图片操作
			authed2.POST("imgs", api.CreateProductImg)
			//商品详情图片操作
			authed2.POST("info-imgs", api.CreateInfoImg)
			//商品参数图片操作
			authed2.POST("param-imgs", api.CreateParamImg)
			//分类操作
			authed2.POST("categories", api.CreateCategory)
			//公告操作
			authed2.POST("notices", api.CreateNotice)
			authed2.PUT("notice", api.UpdateNotice)
		}
	}
	return r
}
