package service

import (
	"context"
	"errors"
	"sync"

	"mall/conf"
	"mall/consts"
	"mall/pkg/utils/ctl"
	util "mall/pkg/utils/log"
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
func (s *FavoriteSrv) FavoriteList(ctx context.Context, req *types.FavoritesServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	favorites, total, err := dao.NewFavoritesDao(ctx).ListFavoriteByUserId(u.Id, req.PageSize, req.PageNum)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	for i := range favorites {
		if conf.UploadModel == consts.UploadModelLocal {
			favorites[i].ImgPath = conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + favorites[i].ImgPath
		}
	}

	return ctl.RespList(favorites, total), nil
}

// FavoriteCreate 创建收藏夹
func (s *FavoriteSrv) FavoriteCreate(ctx context.Context, req *types.FavoriteCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	fDao := dao.NewFavoritesDao(ctx)
	exist, _ := fDao.FavoriteExistOrNot(req.ProductId, u.Id)
	if exist {
		err = errors.New("已经存在了")
		util.LogrusObj.Error(err)
		return
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
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

	product, err := dao.NewProductDao(ctx).GetProductById(req.ProductId)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	favorite := &model.Favorite{
		UserID:    u.Id,
		User:      *user,
		ProductID: req.ProductId,
		Product:   *product,
		BossID:    req.BossId,
		Boss:      *boss,
	}
	err = fDao.CreateFavorite(favorite)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}

// FavoriteDelete 删除收藏夹
func (s *FavoriteSrv) FavoriteDelete(ctx context.Context, req *types.FavoriteDeleteReq) (resp interface{}, err error) {
	favoriteDao := dao.NewFavoritesDao(ctx)
	err = favoriteDao.DeleteFavoriteById(req.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return ctl.RespSuccess(), nil
}
