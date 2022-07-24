package cache

import (
	"fmt"
	"strconv"
)

const (
	//RankKey 每日排名
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
