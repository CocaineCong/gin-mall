package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	"mall/service"
)

func ImportSkillGoods(c *gin.Context) {
	var skillGoodsImport service.SkillGoodsImport
	file, _, _ := c.Request.FormFile("file")
	if err := c.ShouldBind(&skillGoodsImport); err == nil {
		res := skillGoodsImport.Import(file)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err, "ParentImport")
	}
}
