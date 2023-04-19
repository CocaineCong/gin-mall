package v1

import (
	"mall/consts"
	util "mall/pkg/utils"
	"mall/service"

	"github.com/gin-gonic/gin"
)

func ImportSkillGoods(c *gin.Context) {
	var skillGoodsImport service.SkillGoodsImport
	file, _, _ := c.Request.FormFile("file")
	if err := c.ShouldBind(&skillGoodsImport); err == nil {
		res := skillGoodsImport.Import(c.Request.Context(), file)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err, "ImportSkillGoods")
	}
}

func InitSkillGoods(c *gin.Context) {
	var skillGoods service.SkillGoodsService
	if err := c.ShouldBind(&skillGoods); err == nil {
		res := skillGoods.InitSkillGoods(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err, "InitSkillGoods")
	}
}

func SkillGoods(c *gin.Context) {
	var skillGoods service.SkillGoodsService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&skillGoods); err == nil {
		res := skillGoods.SkillGoods(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err, "SkillGoods")
	}
}
