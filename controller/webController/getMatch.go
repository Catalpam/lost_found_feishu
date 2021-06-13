package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/controller/general"
	"lost_found/dbModel"
	"net/http"
)

func GetMatches(ctx *gin.Context) {
	var matches []dbModel.Match
	SelectMatches(&matches,ctx)
	returnMatches(&matches,ctx)
}

func returnMatches(matches *[]dbModel.Match, ctx *gin.Context)  {
	db := common.GetDB()
	if len(*matches) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data":"[]",
			"msg":  "没有查询到符合条件的Match",
		})
		return
	}
	var MatchList []MatchListModel
	for _, value := range *matches {
		var founder dbModel.User
		var loster 	dbModel.User
		var lost 	dbModel.Lost
		db.Where("open_id=?",value.FoundOpenId).First(&founder)
		print(value.ID)
		print(lost.ID)
		print(lost.OpenId)
		db.Where("open_id=?",value.LosterOpenId).First(&loster)
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
			ItemType:       common.TypeId2Name(value.TypeSmallId),
			Info:           value.ItemInfo,
			AdditionalInfo: value.AdditionalInfo,
			Campus:         value.Campus,
			Place:          general.PlaceId2Name(value.PlaceSmallId),
			PlaceDetail:    value.PlaceDetail,
			Image:          value.Image,
			FoundDate:      value.FoundDate,
			FoundTime:      value.Time,
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

func SelectMatches(matches *[]dbModel.Match, ctx *gin.Context)  {
	db := common.GetDB()
	//获取参数
	//subtype 	:= ctx.Query("subtype")
	//openId 		:= ctx.Query("openid")
	//date 		:= ctx.Query("date")
	//timeSession := ctx.Query("time_session")
	//查找数据库
	db.Where(&dbModel.Found{
		//FoundDate:          date,
		//FoundTimeSession:   common.TimeSessionToChinese(),
		//FoundOpenId:        openId,
		//SubType:            subtype,
	}).Find(&matches)
}
