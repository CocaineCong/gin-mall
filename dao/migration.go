package dao

import (
	"fmt"
	"mall/model"
	"os"
)

// Migration 执行数据迁移
func Migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{},
			&model.Product{},
			&model.Carousel{},
			&model.Category{},
			&model.Favorite{},
			&model.ProductImg{},
			&model.Order{},
			&model.Cart{},
			&model.Admin{},
			&model.Address{},
			&model.Notice{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
