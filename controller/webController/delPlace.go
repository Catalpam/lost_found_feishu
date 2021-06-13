package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func DelPlace(ctx *gin.Context) {

	key := ctx.Query("big_key")
	var placeBig dbModel.PlaceBig
	var placeSmalls []dbModel.PlaceSmall

	db := common.GetDB()
	db.Where("id=?",key).First(&placeBig)
	if placeBig.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg": "大地点不存在！",
		})
		return
	}

	willDel := placeBig.Name
	var delIndexx uint = placeBig.Indexx
	for true{
		var big dbModel.PlaceBig
		db.Where("indexx=?",delIndexx+1).First(&big)
		if big.ID == 0 {
			break
		}
		db.Model(&big).Update("indexx",delIndexx)
		println("序号提前："+big.Name)
		delIndexx ++
	}
	db.Delete(&placeBig)
	db.Where("big_id=?",placeBig.ID).Find(&placeSmalls).Delete(&dbModel.PlaceSmall{})

	println("------------删除父地点："+willDel+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除父地点“"+willDel+"”成功",
	})
}

func DelSubPlace(ctx *gin.Context) {
	key := ctx.Query("key")
	var placeSmall dbModel.PlaceSmall
	db := common.GetDB()
	db.Where("id=?",key).First(&placeSmall)

	if placeSmall.ID == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "key格式错误！",
		})
		return
	}
	willDel := placeSmall.Name
	var delIndexx uint = placeSmall.Indexx
	for true{
		var small dbModel.PlaceSmall
		db.Where("indexx=? AND big_id=?",delIndexx+1,placeSmall.BigId).First(&small)
		if small.ID == 0 {
			break
		}
		db.Model(&small).Update("indexx",delIndexx)
		println("序号提前："+small.Name)
		delIndexx ++
	}
	db.Delete(&placeSmall)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除子地点“"+willDel+"”成功",
	})
}