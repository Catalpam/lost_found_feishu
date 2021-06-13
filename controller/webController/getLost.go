package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/controller/general"
	"lost_found/dbModel"
	"net/http"
)


func GetLosts(ctx *gin.Context) {
	var losts []dbModel.Lost
	SelectLost(&losts,ctx)
	returnLosts(&losts,ctx)
}

func returnLosts(losts *[]dbModel.Lost, ctx *gin.Context)  {
	db := common.GetDB()
	var user dbModel.User
	if len(*losts) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data":"[]",
			"msg":  "没有查询到符合条件的Founds",
		})
		return
	}
	var LostList []LostListModel
	for _, value := range *losts {
		var placeSmall dbModel.PlaceSmall
		db.Where("id=?",value.PlaceSmallId1).First(&placeSmall)

		db.Where("open_id=?",value.OpenId).First(&user)
		temp := LostListModel{
			ID:                 value.ID,
			Name:               user.Name,
			Student_Teacher_Id: user.StudentId,
			Moblie:             user.Mobile,
			Avatar:             user.Avatar,
			Campus:             dbModel.CampusId2Str(placeSmall.CampusId),
			SubType:            common.TypeId2Name(value.TypeSmallId),
			Place1:             general.PlaceId2Name(value.PlaceSmallId1),
			Place2:             general.PlaceId2Name(value.PlaceSmallId1),
			Place3:             general.PlaceId2Name(value.PlaceSmallId1),
			LostDate:           value.Date,
			LostTimeSession:    common.TimeSessionToChinese(value.TimeSession),
		}
		LostList = append(LostList, temp)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": LostList,
		"msg":  "获取Found List成功",
	})
}

type LostListModel struct {
	ID uint
	// User
	Name string
	Student_Teacher_Id string
	Moblie string
	Avatar string
	Campus string
	// SubType
	SubType string
	// Location
	Place1 string
	Place2 string
	Place3 string
	// Time
	LostDate string
	LostTimeSession string
}

func SelectLost(losts *[]dbModel.Lost, ctx *gin.Context)  {
	db := common.GetDB()
	//获取参数
	//subtype 	:= ctx.Query("subtype")
	//openId 	:= ctx.Query("openid")
	//date 		:= ctx.Query("date")
	//timeSession := ctx.Query("time_session")
	//查找数据库
	db.Where(&dbModel.Lost{
		//LosterOpenId:    openId,
		//TypeSubName:     subtype,
		//LostDate:        date,
		//LostTimeSession: timeSession,
	}).Where("match_id=?",0).Order("date ASC").Find(&losts)
}
