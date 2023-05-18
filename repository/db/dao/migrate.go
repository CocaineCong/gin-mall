package dao

import (
	"github.com/CocaineCong/gin-mall/repository/db/model"
)

func migrate() {
	_db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
}
