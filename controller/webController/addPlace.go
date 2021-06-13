package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func AddSubPlace(ctx *gin.Context) {

	bigkey := ctx.Query("big_key")
	subName := ctx.Query("name")

	db := common.GetDB()
	var placeBig dbModel.PlaceBig
	var placeSmall dbModel.PlaceSmall
	println("big_key")
	println(ctx.Query("big_key"))
	println("name")
	println(ctx.Query("name"))
	db.Where("id=? ",bigkey).First(&placeBig)
	if placeBig.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "没有这个父地点",
		})
		return
	}
	db.Where("big_id=?",placeBig.ID).Order("indexx DESC").Limit(1).Find(&placeSmall)
	db.Create(&dbModel.PlaceSmall{
		Indexx:   placeSmall.Indexx +1,
		Name:     subName,
		BigId:    placeBig.ID,
		BigName:  placeBig.Name,
		CampusId: placeBig.CampusId,
	})

	println("------------添加子地点："+subName+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加子地点“"+subName+"”成功",
	})
}

func AddPlace(ctx *gin.Context) {
	campusId := ctx.Query("campus_id")
	name := ctx.Query("name")

	var placeBig dbModel.PlaceBig
	db := common.GetDB()
	if campusId == "" || name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "参数格式错误！",
		})
		return
	}
	db.Where("campus_id=?",campusId).Order("indexx DESC").Limit(1).Find(&placeBig)
	newPlaceBig := dbModel.PlaceBig{
		Indexx:   placeBig.Indexx +1,
		Name:     name,
		CampusId: campusId,
	}
	db.Create(&newPlaceBig)

	println("------------添加父地点成功："+newPlaceBig.Name+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加父地点“"+name+"”成功,请刷新页面后查看。",
	})
}

