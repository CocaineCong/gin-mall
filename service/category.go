package service

import (
	"context"
	"sync"

	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/repository/db/dao"
	"mall/types"
)

var CategorySrvIns *CategorySrv
var CategorySrvOnce sync.Once

type CategorySrv struct {
}

func GetCategorySrv() *CategorySrv {
	CategorySrvOnce.Do(func() {
		CategorySrvIns = &CategorySrv{}
	})
	return CategorySrvIns
}

// ListCategory 列举分类
func (s *CategorySrv) ListCategory(ctx context.Context, req *types.ListCategoryServiceReq) (types.Response, error) {
	code := e.SUCCESS
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Error(err)
		code = e.ErrorDatabase
		return types.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}, err
	}
	return types.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   types.BuildCategories(categories),
	}, nil
}
