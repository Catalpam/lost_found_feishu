package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func DelFound(ctx *gin.Context)  {
	db := common.GetDB()
	var found dbModel.Found
	//获取参数
	FoundIdStr := ctx.Query("id")
	FoundId, err := strconv.ParseUint(FoundIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "Found Id 不合法",
		})
		return
	}
	//在数据库中查找Found对象
	if FoundId != 0 {
		db.Where("id=?",FoundId).Order("found_date ASC").Find(&found)
		if found.ID == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"msg":  "Found不存在",
			})
			return
		}
	}
	// 将其变为失效：
	db.Model(&found).Update("validity",false)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功将该Found变为失效！",
	})
}

func DelLost(ctx *gin.Context)  {
	db := common.GetDB()
	var lost dbModel.Lost
	//获取参数
	LostIdStr := ctx.Query("id")
	LostId, err := strconv.ParseUint(LostIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"msg":  "Lost Id不合法",
		})
		return
	}
	//在数据库中查找Found对象
	if LostId != 0 {
		db.Where("id=?",LostId).Order("found_date ASC").Find(&lost)
		if lost.ID == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"msg":  "Lost不存在",
			})
			return
		}
	}
	// 将其变为失效：
	db.Model(&lost).Update("validity",false)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功将该Lost变为失效！",
	})
}


