package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/comander"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func CliamFound(ctx *gin.Context) {
	db := common.GetDB()
	OpenId := ctx.MustGet("open_id").(string)
	var found dbModel.Found
	// 获取Form中的参数 FoundId
	FoundIdStr := ctx.PostForm("id")
	if FoundIdStr == "" {
		FoundIdStr = ctx.PostForm("found_id")
	}
	LostIdStr := ctx.PostForm("lost")

	// 查找参数
	if FoundIdStr == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "缺少参数id或found_id！",
		})
		return
	}
	FoundId, err := strconv.ParseUint(FoundIdStr, 10, 32)
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
	if found.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "没有查询到符合条件的Founds",
		})
		return
	}
	if found.OpenId == OpenId {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 405,
			"msg":  "这是你自己发布的Found！",
		})
		return
	}


	var typeSmall	dbModel.TypeSmall
	var placeSmall 	dbModel.PlaceSmall
	db.Where("id=?",found.TypeSmallId).First(&typeSmall)
	db.Where("id=?",found.PlaceSmallId).First(&placeSmall)

	newMatch := dbModel.Match{
		FoundDate:    found.Date,
		Time:         found.Time,
		TimeSession:  found.TimeSession,
		LosterOpenId: OpenId,
		FoundOpenId:  found.OpenId,
		TypeBigId:    found.TypeBigId,
		TypeSmallId:  found.TypeSmallId,
		TypeName:     typeSmall.BigName+" "+typeSmall.Name,
		PlaceBigId:   found.PlaceBigId,
		PlaceSmallId: found.PlaceSmallId,
		PlaceName:    placeSmall.BigName+" "+placeSmall.Name,
		ItemInfo:           found.ItemInfo,
		Image:              found.Image,
		ImageKey:           found.ImageKey,
		PlaceDetail:        found.PlaceDetail,
		CurrentPlace:       found.CurrentPlace,
		CurrentPlaceDetail: found.CurrentPlaceDetail,
		LosterInfo:         found.LosterInfo,
		AdditionalInfo:     found.AdditionalInfo,
	}
	db.Create(&newMatch)
	println("----------newMatch.ID----------------")
	println(newMatch.ID)

	if errLostId != nil && LostId != 0 {
		db.Model(&dbModel.Lost{}).Where("id=?",LostId).Update("match_id", newMatch.ID)
	}

	db.Model(&found).Update("match_id", newMatch.ID)
	// 若为自己带走，给两边发送信息
	if found.CurrentPlace == "1" {
		go comander.SendUesrToBoth(newMatch.ID)
	}
	// 给小程序返回详情
	returnMatchDeatil(&newMatch, ctx)
}

func returnMatchDeatil(match *dbModel.Match, ctx *gin.Context) {

	MatchDetail := FoundDetailModel{
		ID:                 match.ID,
		SubType:            match.TypeName,
		Campus:             match.Campus,
		Place:              match.PlaceName,
		Image:              match.Image,
		FoundDate:          match.FoundDate,
		FoundTime:          match.Time,
		ItemInfo:           match.ItemInfo,
		CurrentPlace:       match.CurrentPlace,
		CurrentPlaceDetail: match.CurrentPlaceDetail,
		AdditionalInfo:     match.AdditionalInfo,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": MatchDetail,
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
	AdditionalInfo string `gorm:"type:varchar(500)"`
}
