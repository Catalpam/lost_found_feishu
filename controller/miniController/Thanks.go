package miniController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/comander"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func ThanksMsg(ctx *gin.Context) {
	db := common.GetDB()
	var match dbModel.Match
	// 获取Form中的参数 FoundId
	MatchIdStr := ctx.PostForm("id")
	if MatchIdStr == "" {
		MatchIdStr = ctx.PostForm("found_id")
	}
	ThanksMsg := ctx.PostForm("thxmsg")
	//获取用户OpenId
	OpenId := ctx.MustGet("open_id").(string)

	// 查找参数
	if MatchIdStr == "" {
		fmt.Printf("\t\t\t\"code\": 413,\n\t\t\t\"data\": \"\",\n\t\t\t\"msg\":  \"缺少参数id或found_id！\",\n")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "缺少参数id或found_id！",
		})
		return
	}
	// 查找参数
	if ThanksMsg == "" {
		fmt.Printf("\t\t\t\"code\": 413,\n\t\t\t\"data\": \"\",\n\t\t\t\"msg\":  \"诶嘿，你还没有填入感谢的话哦！\",\n")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "诶嘿，你还没有填入感谢的话哦！",
		})
		return
	}
	var found dbModel.Found
	db.Where("id=?",MatchIdStr).First(&found)
	db.Where("id=?", found.MatchId).First(&match)

	if match.LosterOpenId != OpenId{
		fmt.Printf("\t\t\t\"code\": 413,\n\t\t\t\"data\": match.TypeName,\n\t\t\t\"msg\":  \"这件物品好像不是你认领的-哦！\",\n")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": match.TypeName,
			"msg":  "这件物品好像不是你认领的-哦！",
		})
		return
	}

	if match.LosterComment != "" {
		fmt.Printf("\t\t\t\"code\": 400,\n\t\t\t\"data\": \"\",\n\t\t\t\"msg\":  \"你已经表达过感谢了！\",\n")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "你已经表达过感谢了！",
		})
		return
	}

	db.Model(&match).Update("loster_comment", ThanksMsg)
	go comander.SendThx(match.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data":	match.LosterComment,
		"msg":  "感谢信息上传成功",
	})
	return
}


