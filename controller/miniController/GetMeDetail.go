package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
	 "strings"
)

func GetMeDetail(ctx *gin.Context) {
	// 获取Form中的参数 FoundId LostId
	FoundIdStr := ctx.PostForm("FoundId")
	LostIdStr := ctx.PostForm("LostId")

	// 查找参数
	if FoundIdStr != "" {
		FoundId, err := strconv.ParseUint(FoundIdStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": err,
				"msg":  "FoundId格式不合法！",
			})
		}
		returnMeFound(uint(FoundId),ctx)
	} else if LostIdStr != ""{
		LostId, err := strconv.ParseUint(LostIdStr, 10 ,32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": err,
				"msg":  "LostId格式不合法！",
			})
		}
		returnMeLost(uint(LostId),ctx)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "FoundId、LostId中至少输入一种Id！",
		})
	}
}

func returnMeFound(FoundId uint, ctx *gin.Context)  {
	db := common.GetDB()
	var found dbModel.Found
	db.Where("id=?",FoundId).First(&found)
	if found.ID == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "没有找到对应的Found！",
		})
		return
	}
	var isMatched bool
	if found.MatchId == 0{
		isMatched = false
	} else {
		isMatched = true
	}
	tempFound := MyFoundDetail{
		ID:             found.ID,
		IsMatched:      isMatched,
		SubType:        found.SubType,
		FoundDate:      found.FoundDate,
		FoundTime:      found.FoundTime,
		Campus:         found.Campus,
		Place:          found.Place + "-" + found.SubPlace,
		PlaceDetail:    found.PlaceDetail,
		Image:          found.ImageHome,
		ItemInfo:       found.ItemInfo,
		AdditionalInfo: found.AdditionalInfo,
		LosterComment:  found.LosterComment,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": tempFound,
		"msg":  "Found Detail 返回成功！",
	})
}

func returnMeLost(LostId uint, ctx *gin.Context)  {
	db := common.GetDB()
	var found dbModel.Found
	var lost  dbModel.Lost
	db.Where("id=?",LostId).First(&lost)
	if lost.MatchId == 0{
		lostPlaceStr1 := strings.Replace(lost.LostPlace1,"[\"","",-1)
		lostPlaceStr1 = strings.Replace(lostPlaceStr1,"\"]","",-1)
		lostPlaceStr1 = strings.Replace(lostPlaceStr1,"\",\"","-",-1)

		lostPlaceStr2 := strings.Replace(lost.LostPlace2,"[\"","",-1)
		lostPlaceStr2 = strings.Replace(lostPlaceStr2,"\"]","",-1)
		lostPlaceStr2 = strings.Replace(lostPlaceStr2,"\",\"","-",-1)

		lostPlaceStr3 := strings.Replace(lost.LostPlace3,"[\"","",-1)
		lostPlaceStr3 = strings.Replace(lostPlaceStr3,"\"]","",-1)
		lostPlaceStr3 = strings.Replace(lostPlaceStr3,"\",\"","-",-1)
		tempLost := MyLostFalseDetail{
			ID:              lost.ID,
			IsMatched:       false,
			LosterOpenId:    lost.LosterOpenId,
			TypeSubName:     lost.TypeSubName,
			LostPlace1:      lostPlaceStr1,
			LostPlace2:      lostPlaceStr2,
			LostPlace3:      lostPlaceStr3,
			LostDate:        lost.LostDate,
			LostTimeSession: lost.LostTimeSession,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": tempLost,
			"msg":  "未找到的Lost Detail返回成功！",
		})
		return
	}
	if lost.MatchId == 4294967294 {
		tempFound := MyLostTrueDetail{
			ID:             lost.ID,
			IsMatched:      true,
			SubType:        lost.TypeSubName,
			FoundDate:      lost.LostDate,
			FoundTime:      lost.LostTimeSession,
			Campus:         "",
			Place:          "自行找到",
			PlaceDetail:    "自行找到",
			Image:          "https://www.fengzigeng.com/api/image?name=7d016e6add66f758c225c0454653797f.png",
			ItemInfo:       "",
			AdditionalInfo: "",
			LosterComment:  "",
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": tempFound,
			"msg":  "获取自行找到的Lost Detail成功",
		})
		return
	}
	db.Where("id=?",lost.MatchId).First(&found)
	tempFound := MyLostTrueDetail{
		ID:             lost.ID,
		IsMatched:      true,
		SubType:        found.SubType,
		FoundDate:      found.FoundDate,
		FoundTime:      found.FoundTime,
		Campus:         found.Campus,
		Place:          found.Place + "-" + found.SubPlace,
		PlaceDetail:    found.PlaceDetail,
		Image:          found.ImageHome,
		ItemInfo:       found.ItemInfo,
		AdditionalInfo: found.AdditionalInfo,
		LosterComment:  found.LosterComment,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": tempFound,
		"msg":  "获取已找到的Lost Detail成功",
	})
}


type MyFoundDetail struct {
	ID      uint
	IsMatched      bool
	SubType string
	// Time
	FoundDate string
	FoundTime string
	// Location
	Campus      string
	Place       string
	PlaceDetail string
	// Image
	Image          string
	ItemInfo       string
	AdditionalInfo string
	LosterComment  string
}
type MyLostFalseDetail struct {
	ID        uint `gorm:"primary_key"`
	IsMatched bool
	//Loster的OpenId
	LosterOpenId string `gorm:"type:char(50);not null"`
	// 物品类型 为Thing表的NameID
	TypeSubName string `gorm:"type:varchar(20);not null;"`
	// 丢失地点
	LostPlace1 string `gorm:"type:char(100);not null"`
	LostPlace2 string `gorm:"type:char(100);"`
	LostPlace3 string `gorm:"type:char(100);"`
	LostDate string `gorm:"type:char(15);"`
	LostTimeSession string `gorm:"type:varchar(20);"`
}
type MyLostTrueDetail struct {
	ID uint
	IsMatched bool
	SubType string
	// Time
	FoundDate string
	FoundTime string
	LostDate string
	LostPlace string
	// Location
	Campus string
	Place string
	PlaceDetail string
	// Image
	Image string
	ItemInfo string
	AdditionalInfo string
	LosterComment string
}

