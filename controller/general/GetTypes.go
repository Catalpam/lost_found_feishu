package general

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetTypes(ctx *gin.Context)  {
	db := common.GetDB()
	var text = "[["
	var itemTypes []dbModel.ItemType

	db.Order("type_id ASC").Find(&itemTypes)
	println(itemTypes[0].Name)
	for _, itemType := range itemTypes{
		text = text + "\"" + itemType.Name + "\"" + ","
	}
	text = text + "],"

	text = text + "["
	for _, itemType := range itemTypes{
		text = text + itemType.Subtypes +","
	}
	text = text + "]"

	text = text + "]"
	ctx.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data": text,
		"msg": "物品类型Types返回成功",
	})
}