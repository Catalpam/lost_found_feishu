package miniController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/controller/general"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func GetMeDetail(ctx *gin.Context) {
	// 获取Form中的参数 FoundId LostId
	FoundIdStr := ctx.PostForm("FoundId")
	LostIdStr := ctx.PostForm("LostId")
	MatchIdStr := ctx.PostForm("MatchId")


	// 查找参数
	fmt.Println("foundId:"+FoundIdStr)
	fmt.Println("lost ID"+LostIdStr)
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
	} else if MatchIdStr != ""{
		MatchId, err := strconv.ParseUint(MatchIdStr, 10 ,32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": err,
				"msg":  "LostId格式不合法！",
			})
		}
		returnMeMatch(uint(MatchId),ctx)
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
	var placeSmall dbModel.PlaceSmall
	var typeSmall  dbModel.TypeSmall
	var match dbModel.Match
	var tempFound  MyFoundDetail
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
	db.Where("id=?",found.PlaceSmallId).First(&placeSmall)
	db.Where("id=?",found.TypeSmallId).First(&typeSmall)
	if isMatched == true {
		db.Where("id=?",found.MatchId).First(&match)
		tempFound = MyFoundDetail{
			ID:             found.ID,
			IsMatched:      isMatched,
			SubType:        typeSmall.BigName+" "+typeSmall.Name,
			FoundDate:      found.Date,
			FoundTime:      found.Time,
			Campus:         dbModel.CampusId2Str(placeSmall.CampusId),
			Place:          placeSmall.BigName + "-" + placeSmall.Name,
			PlaceDetail:    found.PlaceDetail,
			Image:          found.Image,
			ItemInfo:       found.ItemInfo,
			AdditionalInfo: found.AdditionalInfo,
			LosterComment:  match.LosterComment,
		}
	} else {
		tempFound = MyFoundDetail{
			ID:             found.ID,
			IsMatched:      isMatched,
			SubType:        typeSmall.BigName+" "+typeSmall.Name,
			FoundDate:      found.Date,
			FoundTime:      found.Time,
			Campus:         dbModel.CampusId2Str(placeSmall.CampusId),
			Place:          placeSmall.BigName + "-" + placeSmall.Name,
			PlaceDetail:    found.PlaceDetail,
			Image:          found.Image,
			ItemInfo:       found.ItemInfo,
			AdditionalInfo: found.AdditionalInfo,
			LosterComment:  "还没有失主认领哦～～",
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": tempFound,
		"msg":  "Found Detail 返回成功！",
	})
}

func returnMeLost(LostId uint, ctx *gin.Context)  {
	db := common.GetDB()
	var lost  dbModel.Lost
	db.Where("id=?",LostId).First(&lost)
	if lost.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 403,
			"data": nil,
			"msg":  "没有找到对应的信息！",
		})
		return
	}
	if lost.MatchId == 0 && lost.IsFoundBySelf == false{
		tempLost := MyLostFalseDetail{
			ID:              lost.ID,
			IsMatched:       false,
			LosterOpenId:    lost.OpenId,
			TypeSubName:     common.TypeId2Name(lost.TypeSmallId),
			LostPlace1:      general.PlaceId2Name(lost.PlaceSmallId1),
			LostPlace2:      general.PlaceId2Name(lost.PlaceSmallId2),
			LostPlace3:      general.PlaceId2Name(lost.PlaceSmallId3),
			LostDate:        lost.Date,
			LostTimeSession: lost.TimeSession,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": tempLost,
			"msg":  "未找到的Lost Detail返回成功！",
		})
		return
	}
}

func returnMeMatch(matchId uint, ctx *gin.Context)  {
	db := common.GetDB()
	var match dbModel.Match
	db.Where("id=?",matchId).First(&match)
	tempFound := MyLostTrueDetail{
		ID:             match.ID,
		IsMatched:      true,
		SubType:        common.TypeId2Name(match.TypeSmallId),
		FoundDate:      match.FoundDate,
		FoundTime:      match.Time,
		Campus:         match.Campus,
		Place:          general.PlaceId2Name(match.PlaceSmallId),
		PlaceDetail:    match.PlaceDetail,
		Image:          match.Image,
		ItemInfo:       match.ItemInfo,
		AdditionalInfo: match.AdditionalInfo,
		LosterComment:  match.LosterComment,
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

