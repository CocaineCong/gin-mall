package model

//执行数据迁移

func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Product{}).
		AutoMigrate(&Carousel{}).
		AutoMigrate(&Category{}).
		AutoMigrate(&Favorite{}).
		AutoMigrate(&ProductImg{}).
		AutoMigrate(&ProductInfoImg{}).
		AutoMigrate(&ProductParamImg{}).
		AutoMigrate(&Order{}).
		AutoMigrate(&Cart{}).
		AutoMigrate(&Admin{}).
		AutoMigrate(&Address{})
	DB.Model(&Cart{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	DB.Model(&Order{}).AddForeignKey("user_id","User(id)","CASCADE","CASCADE")
	DB.Model(&Order{}).AddForeignKey("address_id","Address(id)","CASCADE","CASCADE")
	DB.Model(&Order{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	DB.Model(&Order{}).AddForeignKey("boss_id","User(id)","CASCADE","CASCADE")
	DB.Model(&Favorite{}).AddForeignKey("boss_id","User(id)","CASCADE","CASCADE")
	DB.Model(&Favorite{}).AddForeignKey("user_id","User(id)","CASCADE","CASCADE")
	DB.Model(&Favorite{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	DB.Model(&Product{}).AddForeignKey("category_id","Category(id)","CASCADE","CASCADE")
	DB.Model(&ProductImg{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	DB.Model(&ProductInfoImg{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	DB.Model(&ProductParamImg{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	DB.Model(&Address{}).AddForeignKey("user_id","User(id)","CASCADE","CASCADE")
}
