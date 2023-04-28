package service

import (
	"context"
	"errors"
	"sync"

	util "mall/pkg/utils"
	"mall/pkg/utils/ctl"
	"mall/repository/db/dao"
	"mall/repository/db/model"
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
func (s *FavoriteSrv) FavoriteList(ctx context.Context, uId uint, req *types.FavoritesServiceReq) (resp interface{}, err error) {
	favorites, total, err := dao.NewFavoritesDao(ctx).ListFavoriteByUserId(uId, req.PageSize, req.PageNum)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return ctl.RespList(favorites, total), nil
}

// FavoriteCreate 创建收藏夹
func (s *FavoriteSrv) FavoriteCreate(ctx context.Context, uId uint, req *types.FavoritesServiceReq) (resp interface{}, err error) {
	favoriteDao := dao.NewFavoritesDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(req.ProductId, uId)
	if exist {
		err = errors.New("已经存在了")
		util.LogrusObj.Error(err)
		return
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(req.BossId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ProductId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
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
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}

// FavoriteDelete 删除收藏夹
func (s *FavoriteSrv) FavoriteDelete(ctx context.Context, req *types.FavoritesServiceReq) (resp interface{}, err error) {
	favoriteDao := dao.NewFavoritesDao(ctx)
	err = favoriteDao.DeleteFavoriteById(req.FavoriteId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}
