package cache

import (
	"fmt"
	"strconv"
)

const (
	// RankKey 每日排名
	RankKey             = "rank"
	SkillProductKey     = "skill:product:%d"
	SkillProductListKey = "skill:product_list"
	SkillProductUserKey = "skill:user:%s"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
