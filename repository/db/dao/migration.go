package dao

import (
	"fmt"
	"os"

	model2 "mall/repository/db/model"
)

// Migration 执行数据迁移
func Migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model2.User{},
			&model2.Product{},
			&model2.Carousel{},
			&model2.Category{},
			&model2.Favorite{},
			&model2.ProductImg{},
			&model2.Order{},
			&model2.Cart{},
			&model2.Admin{},
			&model2.Address{},
			&model2.Notice{},
			&model2.SkillGoods{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
