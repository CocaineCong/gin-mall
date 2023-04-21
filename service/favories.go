package service

import (
	"context"
	"sync"

	logging "github.com/sirupsen/logrus"

	"mall/pkg/e"
	"mall/repository/db/dao"
	"mall/repository/db/model"
	"mall/serializer"
	"mall/types"
)

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

type FavoriteSrv struct {
}

func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

// FavoriteList 商品收藏夹
func (s *FavoriteSrv) FavoriteList(ctx context.Context, uId uint, req *types.FavoritesServiceReq) serializer.Response {
	favoritesDao := dao.NewFavoritesDao(ctx)
	code := e.SUCCESS
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	favorites, total, err := favoritesDao.ListFavoriteByUserId(uId, req.PageSize, req.PageNum)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(total))
}

// FavoriteCreate 创建收藏夹
func (s *FavoriteSrv) FavoriteCreate(ctx context.Context, uId uint, req *types.FavoritesServiceReq) serializer.Response {
	code := e.SUCCESS
	favoriteDao := dao.NewFavoritesDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(req.ProductId, uId)
	if exist {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(req.BossId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ProductId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		UserID:    uId,
		User:      *user,
		ProductID: req.ProductId,
		Product:   *product,
		BossID:    req.BossId,
		Boss:      *boss,
	}
	favoriteDao = dao.NewFavoritesDaoByDB(favoriteDao.DB)
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// FavoriteDelete 删除收藏夹
func (s *FavoriteSrv) FavoriteDelete(ctx context.Context, req *types.FavoritesServiceReq) serializer.Response {
	code := e.SUCCESS

	favoriteDao := dao.NewFavoritesDao(ctx)
	err := favoriteDao.DeleteFavoriteById(req.FavoriteId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}
