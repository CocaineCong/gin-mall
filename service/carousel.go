package service

import (
	"context"
	"sync"

	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/log"
	"mall/repository/db/dao"
	"mall/types"
)

var CarouselSrvIns *CarouselSrv
var CarouselSrvOnce sync.Once

type CarouselSrv struct {
}

func GetCarouselSrv() *CarouselSrv {
	CarouselSrvOnce.Do(func() {
		CarouselSrvIns = &CarouselSrv{}
	})
	return CarouselSrvIns
}

// ListCarousel 列表
func (s *CarouselSrv) ListCarousel(ctx context.Context, req *types.ListCarouselReq) (resp interface{}, err error) {
	carousels, err := dao.NewCarouselDao(ctx).ListCarousel()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(carousels, int64(len(carousels))), nil
}
