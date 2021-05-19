package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/comander"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func ThanksMsg(ctx *gin.Context) {
	db := common.GetDB()
	var found dbModel.Found
	var lost  dbModel.Lost
	// 获取Form中的参数 FoundId
	FoundIdStr := ctx.PostForm("id")
	if FoundIdStr == "" {
		FoundIdStr = ctx.PostForm("found_id")
	}
	ThanksMsg := ctx.PostForm("thxmsg")
	//获取用户OpenId
	OpenId := ctx.MustGet("open_id").(string)

	// 查找参数
	if FoundIdStr == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "缺少参数id或found_id！",
		})
		return
	}
	// 查找参数
	if ThanksMsg == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "缺少参数thxmsg！",
		})
		return
	}

	FoundId, err := strconv.ParseUint(FoundIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": err,
			"msg":  "id不合法！",
		})
		return
	}
	db.Where("id=?", FoundId).First(&found)
	if found.MatchId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "Found还未被被认领！",
		})
		return
	}
	// 查找对应的Lost
	db.Where("id=?", found.MatchId).First(&lost)
	if lost.LosterOpenId != OpenId {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  "该Found不是被你认领！",
		})
		return
	}

	if found.LosterComment != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "你已经表达过感谢了！",
		})
		return
	}

	db.Model(&found).Update("loster_comment", ThanksMsg)
	go comander.SendThx(found.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data":found.LosterComment,
		"msg":  "感谢信息上传成功",
	})
	return
}


