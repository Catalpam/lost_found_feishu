package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetMatches(ctx *gin.Context) {
	var founds []dbModel.Found
	SelectMatches(&founds,ctx)
	returnMatches(&founds,ctx)
}

func returnMatches(losts *[]dbModel.Found, ctx *gin.Context)  {
	db := common.GetDB()
	if len(*losts) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data":"[]",
			"msg":  "没有查询到符合条件的Match",
		})
		return
	}
	var MatchList []MatchListModel
	for _, value := range *losts {
		var founder dbModel.User
		var loster 	dbModel.User
		var lost 	dbModel.Lost
		db.Where("open_id=?",value.FoundOpenId).First(&founder)
		print(value.MatchId)
		db.Where("id=?",value.MatchId).First(&lost)
		print(lost.ID)
		print(lost.LosterOpenId)
		db.Where("open_id=?",lost.LosterOpenId).First(&loster)
		println(loster.Name)
		temp := MatchListModel{
			FoundID:        value.ID,
			FounderName:    founder.Name,
			FounderId:      founder.StudentId,
			FounderMoblie:  founder.Mobile,
			FounderAvatar:  founder.Avatar,
			LosterName:     loster.Name,
			LosterId:       loster.StudentId,
			LosterMoblie:   loster.Mobile,
			LosterAvatar:   loster.Avatar,
			ItemType:       value.ItemType+"-"+value.SubType,
			Info:           value.ItemInfo,
			AdditionalInfo: value.AdditionalInfo,
			Campus:         value.Campus,
			Place:          value.Place,
			PlaceDetail:    value.PlaceDetail,
			Image:          value.ImageHome,
			FoundDate:      value.FoundDate,
			FoundTime:      value.FoundTime,
			ThxMsg:         value.LosterComment,
		}
		MatchList = append(MatchList, temp)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": MatchList,
		"msg":  "获取Match List成功",
	})
}

type MatchListModel struct {
	FoundID uint
	// User
	FounderName string
	FounderId string
	FounderMoblie string
	FounderAvatar string

	LosterName string
	LosterId string
	LosterMoblie string
	LosterAvatar string

	ItemType string
	Info string
	AdditionalInfo string

	// Location
	Campus string
	Place string
	PlaceDetail string
	// Image
	Image string
	// Time
	FoundDate string
	FoundTime string
	ThxMsg string
}

func SelectMatches(founds *[]dbModel.Found, ctx *gin.Context)  {
	db := common.GetDB()
	//获取参数
	subtype 	:= ctx.Query("subtype")
	openId 		:= ctx.Query("openid")
	date 		:= ctx.Query("date")
	timeSession := ctx.Query("time_session")
	//查找数据库
	db.Where(&dbModel.Found{
		FoundDate:          date,
		FoundTimeSession:   timeSession,
		FoundOpenId:        openId,
		SubType:            subtype,
	}).Where("match_id>=?",1).Find(&founds)
}
