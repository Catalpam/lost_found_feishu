package miniController

//}//import (
////	"encoding/json"
////	"github.com/gin-gonic/gin"
////	"lost_found/common"
////	"lost_found/dbModel"
////	"net/http"
////	"strconv"
////	"strings"
////)
////
////func GetFoundList(ctx *gin.Context) {
////	var lost dbModel.Found
////	// 获取Form中的参数 FoundId LostId
////	FoundIdStr := ctx.PostForm("FoundId")
////	LostIdStr := ctx.PostForm("LostId")
////
////	// 查找参数
////	if FoundIdStr != "" {
////		FoundId, err := strconv.ParseUint(ctx.PostForm("id"), 10 ,32)
////		if err != nil {
////			ctx.JSON(http.StatusOK, gin.H{
////				"code": 413,
////				"data": err,
////				"msg":  "FoundId格式不合法！",
////			})
////		}
////		returnMeFound(FoundId,ctx)
////	} else if {
////		LostId, err := strconv.ParseUint(ctx.PostForm("id"), 10 ,32)
////		if err != nil {
////			ctx.JSON(http.StatusOK, gin.H{
////				"code": 413,
////				"data": err,
////				"msg":  "LostId格式不合法！",
////			})
////		}
////	} else {
////		ctx.JSON(http.StatusOK, gin.H{
////			"code": 413,
////			"data": err,
////			"msg":  "FoundId、LostId中至少输入一种Id！",
////		})
////	}
////}
////
////func returnMeFound(FoundId uint64, ctx *gin.Context)  {
////	db := common.GetDB()
////	var found dbModel.Found
////	db.Where("FoundId=?",FoundId).First(&found)
////	if found.ID == 0{
////
////	}
////	tempFound := MyFoundDetail{
////			ID:        			found.ID,
////			SubType:   			found.SubType,
////			Campus:    			found.Campus,
////			Place:     			found.Place+"-"+found.SubPlace,
////			PlaceDetail: 		found.PlaceDetail,
////			Image:	   			found.ImageHome,
////			ImageList: 			found.Image,
////			FoundDate: 		 	found.FoundDate,
////			FoundTime: 		 	found.FoundTime,
////			ItemInfo: 		 	found.ItemInfo,
////			AdditionalInfo : 	found.AdditionalInfo,
////		}
////	}
////
////	ctx.JSON(http.StatusOK, gin.H{
////		"code": 200,
////		"data": tempFound,
////		"msg":  "获取Found List成功",
////	})
////}
////
////func returnMeLost(founds *dbModel.Lost, ctx *gin.Context)  {
////	if len(*founds) == 0{
////		ctx.JSON(http.StatusOK, gin.H{
////			"code": 404,
////			"msg":  "没有查询到符合条件的Founds",
////		})
////		return
////	}
////	tempFound := MyFoundDetail{
////		ID:        			value.ID,
////		SubType:   			value.SubType,
////		Campus:    			value.Campus,
////		Place:     			value.Place+"-"+value.SubPlace,
////		PlaceDetail: 		value.PlaceDetail,
////		Image:	   			value.ImageHome,
////		ImageList: 			value.Image,
////		FoundDate: 		 	value.FoundDate,
////		FoundTime: 		 	value.FoundTime,
////		ItemInfo: 		 	value.ItemInfo,
////		AdditionalInfo : 	value.AdditionalInfo,
////	}
////}
////
////func ReturnFalse(ctx *gin.Context)  {
////	ctx.JSON(http.StatusOK, gin.H{
////		"code": 404,
////		"msg":  "没有查询到符合条件的Founds",
////	})
//
//
//
//
//type MyFoundDetail struct {
//	ID uint
//	SubType string
//	// Location
//	Campus string
//	Place string
//	PlaceDetail string
//	// Image
//	Image string
//	ImageList string
//	// Time
//	FoundDate string
//	FoundTime string
//	ItemInfo string
//	AdditionalInfo string
//}