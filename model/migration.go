package model

import (
	"fmt"
	"os"
)

//执行数据迁移

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{},
			&Product{},
			&Carousel{},
			&Category{},
			&Favorite{},
			&ProductImg{},
			&Order{},
			&Cart{},
			&Admin{},
			&Address{},
			&Notice{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
