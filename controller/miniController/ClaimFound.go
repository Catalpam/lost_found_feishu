package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func CliamFound(ctx *gin.Context) {
	db := common.GetDB()
	var found dbModel.Found
	// 获取Form中的参数 FoundId
	FoundIdStr := ctx.PostForm("id")
	LostIdStr := ctx.PostForm("lost")
	// 查找参数
	if FoundIdStr == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "缺少参数id！",
		})
		return
	}
	FoundId, err := strconv.ParseUint(ctx.PostForm("id"), 10, 32)
	LostId, errLostId := strconv.ParseUint(LostIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": err,
			"msg":  "id不合法！",
		})
		return
	}
	db.Where("id=?", FoundId).First(&found)
	if found.MatchId != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "Found已被认领！",
		})
		return
	}
	if errLostId != nil && LostId != 0 {
		db.Model(&found).Update("match_id", LostId)
	} else {
		newLost := dbModel.Lost{
			LosterOpenId:    ctx.MustGet("open_id").(string),
			TypeSubName:     found.SubType,
			LostPlace1:      "[\"" + found.Place + "\",\"" + found.SubPlace + "\"]",
			LostDate:        found.FoundDate,
			LostTimeSession: found.FoundTimeSession,
			MatchID:         found.ID,
		}
		db.Create(&newLost)
		println("----------newLost.ID----------------")
		println(newLost.ID)
		db.Model(&found).Update("match_id", newLost.ID)
	}
	returnFoundDeatil(&found, ctx)
}

func returnFoundDeatil(found *dbModel.Found, ctx *gin.Context) {
	if found.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "没有查询到符合条件的Founds",
		})
		return
	}

	FoundDetail := FoundDetailModel{
		ID:                 found.ID,
		SubType:            found.SubType,
		Campus:             found.Campus,
		Place:              found.Place,
		Image:              found.ImageHome,
		FoundDate:          found.FoundDate,
		FoundTime:          found.FoundTime,
		ItemInfo:           found.ItemInfo,
		CurrentPlace:       found.CurrentPlace,
		CurrentPlaceDetail: found.CurrentPlaceDetail,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": FoundDetail,
		"msg":  "认领成功！",
	})
	return
}

type FoundDetailModel struct {
	ID      uint
	SubType string
	// Location
	Campus string
	Place  string
	// Image
	Image string
	// Time
	FoundDate string
	FoundTime string
	ItemInfo  string
	// 当前位置：0-留在原地 1-自己带走 2-放在他处
	CurrentPlace string `gorm:"type:char(1);not null"`
	// 当前位置：0-留在原地 1-自己带走 2-放在他处
	CurrentPlaceDetail string `gorm:"type:char(200);"`
}
