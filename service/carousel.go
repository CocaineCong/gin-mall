package service

import (
	"context"
	"sync"

	logging "github.com/sirupsen/logrus"

	"mall/pkg/e"
	"mall/repository/db/dao"
	"mall/serializer"
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
func (s *CarouselSrv) ListCarousel(ctx context.Context, req *types.ListCarouselsServiceReq) (serializer.Response, error) {
	code := e.SUCCESS
	carousels, err := dao.NewCarouselDao(ctx).ListCarousel()
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels))), nil
}
