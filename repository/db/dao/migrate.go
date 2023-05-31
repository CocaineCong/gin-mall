package dao

import (
	"github.com/CocaineCong/gin-mall/repository/db/model"
)

func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Favorite{},
			&model.Order{}, &model.Admin{}, &model.Address{},
			&model.Cart{}, &model.Category{}, &model.Carousel{},
			&model.Notice{}, &model.Notice{}, &model.Product{},
			&model.ProductImg{}, &model.SkillProduct{},
			&model.SkillProduct2MQ{},
		)

	return
}
