package general

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetTypes(ctx *gin.Context)  {
	db := common.GetDB()
	var itemTypes []dbModel.ItemType

	db.Order("type_id ASC").Find(&itemTypes)
	var typeBig []string
	var typeSmall []typeSmallIndex
	for _, itemType := range itemTypes{
		typeBig = append(typeBig, itemType.Name)
		var index typeSmallIndex
		json.Unmarshal([]byte(itemType.Subtypes), &index)
		typeSmall = append(typeSmall,index)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data": gin.H{
			"type1":typeBig,
			"type2":typeSmall,
		},
		"msg": "物品类型Types返回成功",
	})
}

type typeSmallIndex []string
