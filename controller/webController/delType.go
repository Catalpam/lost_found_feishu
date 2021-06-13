package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func DelTypeBig(ctx *gin.Context) {

	key := ctx.Query("big_key")
	var typeBig dbModel.TypeBig
	var typeSmalls []dbModel.TypeSmall

	db := common.GetDB()
	db.Where("id=?",key).First(&typeBig)
	if typeBig.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg": "父物品类型不存在！",
		})
		return
	}

	willDel := typeBig.Name
	var delIndexx uint = typeBig.Indexx
	for true{
		var big dbModel.TypeBig
		db.Where("indexx=?",delIndexx+1).First(&big)
		if big.ID == 0 {
			break
		}
		db.Model(&big).Update("indexx",delIndexx)
		println("序号提前："+big.Name)
		delIndexx ++
	}
	db.Delete(&typeBig)
	db.Where("big_id=?", typeBig.ID).Find(&typeSmalls).Delete(&dbModel.TypeSmall{})

	println("------------删除父Type："+willDel+"------------------")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除父物品类型“"+willDel+"”成功",
	})
}

func DelTypeSmall(ctx *gin.Context) {
	key := ctx.Query("key")
	var typeSmall dbModel.TypeSmall
	db := common.GetDB()
	db.Where("id=?",key).First(&typeSmall)

	if typeSmall.ID == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "key格式错误！",
		})
		return
	}
	willDel := typeSmall.Name
	var delIndexx uint = typeSmall.Indexx
	for true{
		var small dbModel.TypeSmall
		db.Where("indexx=? AND big_id=?",delIndexx+1, typeSmall.BigId).First(&small)
		if small.ID == 0 {
			break
		}
		db.Model(&small).Update("indexx",delIndexx)
		println("序号提前："+small.Name)
		delIndexx ++
	}
	db.Delete(&typeSmall)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除子物品类型“"+willDel+"”成功",
	})
}