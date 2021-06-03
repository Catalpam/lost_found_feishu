package webController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/common"
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
		var place1 []string
		json.Unmarshal([]byte(value.LostPlace1), &place1)
		var place2 []string
		json.Unmarshal([]byte(value.LostPlace1), &place2)
		var place3 []string
		json.Unmarshal([]byte(value.LostPlace1), &place3)
        var user dbModel.User
		db.Where("open_id=?",value.LosterOpenId).First(&user)
		temp := LostListModel{
			ID:              		value.ID,
			Name:            		user.Name,
			Student_Teacher_Id:     user.StudentId,
			Moblie:          		user.Mobile,
			Avatar:          		user.Avatar,
			SubType:         		value.TypeSubName,
			Place1:          		place1[0] + "-" + place1[1],
			Place2:          		place2[0] + "-" + place2[1],
			Place3:          		place3[0] + "-" + place3[1],
			LostDate:        		value.LostDate,
			LostTimeSession: 		value.LostTimeSession,
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
	subtype 	:= ctx.Query("subtype")
	openId 	:= ctx.Query("openid")
	date 		:= ctx.Query("date")
	timeSession := ctx.Query("time_session")
	//查找数据库
	db.Where(&dbModel.Lost{
		LosterOpenId:    openId,
		TypeSubName:     subtype,
		LostDate:        date,
		LostTimeSession: timeSession,
	}).Where("match_id=?",0).Order("lost_date ASC").Find(&losts)
}

