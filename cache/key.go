package cache

import (
	"fmt"
	"strconv"
)

const (
	//RankKey 每日排名
	RankKey = "rank"
	//ElectricalRank 家电排名
	ElectricalRank = "elecRank"
	//AccessoryRank 配件排名
	AccessoryRank = "acceRank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
