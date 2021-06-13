package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetFounds(ctx *gin.Context) {
	var founds []dbModel.Found
	SelectFound(&founds,ctx)
	returnFounds(&founds,ctx)
}

func returnFounds(founds *[]dbModel.Found, ctx *gin.Context)  {
	db := common.GetDB()
	if len(*founds) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data":"[]",
			"msg":  "没有查询到符合条件的Founds",
		})
		return
	}
	var FoundList []FoundListModel
	for _, value := range *founds {
		var user dbModel.User
		var placeSmall dbModel.PlaceSmall
		db.Where("open_id=?",value.OpenId).First(&user)
		db.Where("id=?",value.PlaceSmallId).First(&placeSmall)
		tempFound := FoundListModel{
			ID:        				value.ID,
			Validity:				value.Validity,
			Name:            		user.Name,
			Student_Teacher_Id:     user.StudentId,
			Moblie:          		user.Mobile,
			Avatar:          		user.Avatar,
			SubType:   				common.TypeId2Name(value.TypeSmallId),
			Campus:    				dbModel.CampusId2Str(placeSmall.CampusId),
			Place:     				placeSmall.BigName+"-"+placeSmall.Name,
			PlaceDetail: 			value.PlaceDetail,
			Image:	   				value.Image,
			ImageList: 				value.Image,
			FoundDate: 		 		value.Date,
			FoundTime: 			 	value.Time,
			Info: 		 		    value.ItemInfo,
			AdditionalInfo : 		value.AdditionalInfo,
		}
		FoundList = append(FoundList, tempFound)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": FoundList,
		"msg":  "获取Found List成功",
	})
}

type FoundListModel struct {
	ID uint
	Validity bool
	SubType string
	// User
	Name string
	Student_Teacher_Id string
	Moblie string
	Avatar string
	// Location
	Campus string
	Place string
	PlaceDetail string
	// Image
	Image string
	ImageList string
	// Time
	FoundDate string
	FoundTime string
	Info string
	AdditionalInfo string
}

func SelectFound(founds *[]dbModel.Found, ctx *gin.Context)  {
	db := common.GetDB()
	//获取参数
	//typeIndex 	:= ctx.Query("subtype")
	//campus 	:= ctx.Query("campus")
	//date 		:= ctx.Query("date")
	//timeSession := ctx.Query("time_session")

	//在数据库中查找Found对象
	db.Where(&dbModel.Found{
		//SubType:            typeIndex,
		//Campus:             campus,
		//FoundDate:          date,
		//FoundTimeSession:   timeSession,
	}).Where("match_id=?",0).Order("date ASC").Find(&founds)
}