package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func AddTypeSmall(ctx *gin.Context) {

	bigkey := ctx.Query("big_key")
	subName := ctx.Query("name")
	db := common.GetDB()
	var typeBig dbModel.TypeBig
	var typeSmall dbModel.TypeSmall

	db.Where("id=? ",bigkey).First(&typeBig)
	if typeBig.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "没有这个父类型",
		})
		return
	}
	db.Where("big_id=?", typeBig.ID).Order("indexx DESC").Limit(1).Find(&typeSmall)
	db.Create(&dbModel.TypeSmall{
		Indexx:  typeSmall.Indexx +1,
		Name:    subName,
		BigId:   typeBig.ID,
		BigName: typeBig.Name,
	})

	println("------------添加子地点："+subName+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加子物品类型“"+subName+"”成功,请刷新页面后查看。",
	})
}

func AddTypeBig(ctx *gin.Context) {
	name := ctx.Query("name")
	var typeBig dbModel.TypeBig
	db := common.GetDB()
	db.Order("indexx DESC").Limit(1).Find(&typeBig)
	newPlaceBig := dbModel.TypeBig{
		Indexx:   typeBig.Indexx +1,
		Name:     name,
	}
	db.Create(&newPlaceBig)

	println("------------添加父Type成功："+newPlaceBig.Name+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加父类型“"+name+"”成功,请刷新页面后查看。",
	})
}

