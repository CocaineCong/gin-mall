package service

import (
	"sync"
)

var SkillProductSrvIns *SkillProductSrv
var SkillProductSrvOnce sync.Once

type SkillProductSrv struct {
}

func GetSkillProductSrv() *SkillProductSrv {
	SkillProductSrvOnce.Do(func() {
		SkillProductSrvIns = &SkillProductSrv{}
	})
	return SkillProductSrvIns
}

// SkillProductMQ2MySQL 从mq落库
// func SkillProductMQ2MySQL(ctx context.Context, req *story_types.LikeStoryReq) (err error) {
// 	storyDao := dao.NewStoryDao(ctx)
// 	usDao := dao.NewUserStoryDao(ctx)
// 	err = storyDao.UpdateStoryLikeOrStar(req.StoryId, 1, false)
// 	if err != nil {
// 		log.LogrusObj.Infoln(err)
// 		return
// 	}
//
// 	err = usDao.UserStoryUpsert(&user_story_types.UserStoryReq{
// 		UserId:        req.UserId,
// 		StoryId:       req.StoryId,
// 		OperationType: user_story_consts.UserStoryOperationTypeLike,
// 	})
// 	if err != nil {
// 		log.LogrusObj.Infoln(err)
// 		return
// 	}
//
// 	return
// }
