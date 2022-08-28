package service

import (
	"context"
	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	logging "github.com/sirupsen/logrus"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"mime/multipart"
	"strconv"
)

type SkillGoodsImport struct {
}

func (service *SkillGoodsImport) Import(ctx context.Context, file multipart.File) serializer.Response {
	xlFile, err := xlsx.OpenReader(file)
	if err != nil {
		logging.Info(err)
	}
	code := e.SUCCESS
	rows := xlFile.GetRows("Sheet1")
	length := len(rows[1:])
	skillGoods := make([]*model.SkillGoods, length, length)
	for index, colCell := range rows[1:] {
		pId, _ := strconv.Atoi(colCell[0])
		bId, _ := strconv.Atoi(colCell[1])
		num, _ := strconv.Atoi(colCell[3])
		money, _ := strconv.ParseFloat(colCell[4], 64)
		skillGood := &model.SkillGoods{
			ProductId: uint(pId),
			BossId:    uint(bId),
			Title:     colCell[2],
			Money:     money,
			Num:       num,
		}
		skillGoods[index-1] = skillGood
	}
	err = dao.NewSkillGoodsDao(ctx).CreateByList(skillGoods)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "上传失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
