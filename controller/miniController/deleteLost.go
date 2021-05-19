package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
	"time"
)

func HasFoundBySelf(ctx *gin.Context) {
	var db = common.GetDB()
	var lost dbModel.Lost
	// 获取Form中的参数 FoundId
	LostIdStr := ctx.PostForm("LostId")
	OpenId := ctx.MustGet("open_id").(string)

	// 查找参数
	if LostIdStr == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "缺少参数LostId！",
		})
		return
	}
	LostId, err := strconv.ParseUint(LostIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": err,
			"msg":  "id不合法！",
		})
		return
	}
	db.Where("id=?", LostId).First(&lost)
	if lost.LosterOpenId != OpenId {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 403,
			"data": err,
			"msg":  "这不是你的Lost！",
		})
		return
	}
	db.Model(&lost).Update("match_id", 4294967294)
	db.Model(&lost).Update("lost_date", time.Now().Format("2006-01-02"))
	db.Model(&lost).Update("lost_time_session", time.Now().Format("15:04"))
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": err,
		"msg":  "自行找到成功！",
	})
	return
}
