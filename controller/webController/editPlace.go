package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func EditPlaceSmall(ctx *gin.Context) {
	id := ctx.Query("key")
	name := ctx.Query("name")

	db := common.GetDB()
	if id == "" || name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "啊哦，你还没有填入修改的名称～",
		})
		return
	}
	if id == ""{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "啊哦，你还没有填入修改的名称～",
		})
		return
	}
	db.Model(&dbModel.PlaceSmall{}).Where("id=?",id).Update("name",name)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已成功将名称更改为“"+name+"”，请刷新页面后查看。",
	})
}

func EditPlaceBig(ctx *gin.Context) {
	id := ctx.Query("big_key")
	name := ctx.Query("name")

	db := common.GetDB()
	if id == "" || name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "啊哦，你还没有填入修改的名称～",
		})
		return
	}
	if id == ""{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "啊哦，你还没有填入修改的名称～",
		})
		return
	}
	db.Model(&dbModel.PlaceBig{}).Where("id=?",id).Update("name",name)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "已成功将名称更改为“"+name+"”，请刷新页面后查看。",
	})
}