package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/cardMessage"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func AddLost(ctx *gin.Context) {
	//获取参数
	campusId, _ := ctx.GetPostForm("campus_id")
	LostDate := ctx.PostForm("lost_date")
	timeSession, _ := ctx.GetPostForm("time_session")

	//检查有无空参数
	{
		if campusId == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少校区参数：campus_id",
			})
			return
		}
		if timeSession == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：time_session",
			})
			return
		}

		switch timeSession {
		case "morning", "noon", "afternoon", "evening", "night":
			println("timeSession合法")
		default:
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "参数不合法：time_session应为五个参数中的一种",
			})
			return
		}
	}

	db := common.GetDB()
	if ctx.PostForm("place_1") == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 415,
			"data": "",
			"msg":  "Place不合法!",
		})
		return
	}
	if ctx.PostForm("place_1") == ctx.PostForm("place_2") || ctx.PostForm("place_1") == ctx.PostForm("place_3") || ctx.PostForm("place_2") == ctx.PostForm("place_3"){
		ctx.JSON(http.StatusOK, gin.H{
			"code": 415,
			"data": "",
			"msg":  "选择的位置有重复，若选择多个位置，请选择不同的位置!",
		})
		return
	}


	placeIndex1, _ := String2Index(ctx.PostForm("place_1"))
	placeIndex2, _ := String2Index(ctx.PostForm("place_2"))
	placeIndex3, _ := String2Index(ctx.PostForm("place_3"))
	//查找index对应的值是否存在于数据库中
	var typeBig   dbModel.TypeBig
	var typeSmall dbModel.TypeSmall
	typeIndex ,errType := String2Index(ctx.PostForm("type_index"))
	if errType == nil {
		db.Where("indexx=?", typeIndex[0]).First(&typeBig)
		db.Where("indexx=? AND big_id=?", typeIndex[1], typeBig.ID).First(&typeSmall)
	} else  {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 415,
			"data": "",
			"msg":  "物品类型格式不合法!",
		})
		return
	}

	if typeSmall.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的type_index不存在!",
		})
		return
	}

	var placeSmallIds = [3]uint{0,0,0}
	var placeSmallNames = [3]string{"","",""}
	{
		var placeBig dbModel.PlaceBig
		var placeSmall dbModel.PlaceSmall
		db.Where("indexx=? AND campus_id=?", placeIndex1[0], campusId).First(&placeBig)
		db.Where("indexx=? AND big_id=?", placeIndex1[1], placeBig.ID).First(&placeSmall)
		placeSmallIds[0] = placeSmall.ID
		placeSmallNames[0] = placeSmall.Name
	}
	{
		var placeBig dbModel.PlaceBig
		var placeSmall dbModel.PlaceSmall
		db.Where("indexx=? AND campus_id=?", placeIndex2[0], campusId).First(&placeBig)
		db.Where("indexx=? AND big_id=?", placeIndex2[1], placeBig.ID).First(&placeSmall)
		placeSmallIds[1] = placeSmall.ID
		placeSmallNames[1] = placeSmall.Name
	}
	{
		var placeBig dbModel.PlaceBig
		var placeSmall dbModel.PlaceSmall
		db.Where("indexx=? AND campus_id=?", placeIndex3[0], campusId).First(&placeBig)
		db.Where("indexx=? AND big_id=?", placeIndex3[1], placeBig.ID).First(&placeSmall)
		placeSmallIds[2] = placeSmall.ID
		placeSmallNames[2] = placeSmall.Name
	}


	//获取用户OpenId
	OpenId := ctx.MustGet("open_id").(string)
	var user dbModel.User
	db.Where("open_id=?",OpenId).First(&user)

	//将新的Lost对象添加至数据库中
	var newLost dbModel.Lost
	if ctx.PostForm("place_2") != "" && ctx.PostForm("place_3") != "" {
		newLost = dbModel.Lost{
			Validity:      true,
			IsFoundBySelf: false,
			OpenId:        OpenId,
			Name:          user.Name,
			MatchId:       0,
			TypeBigId:     typeSmall.BigId,
			TypeSmallId:   typeSmall.ID,
			PlaceSmallId1: placeSmallIds[0],
			PlaceSmallId2: placeSmallIds[1],
			PlaceSmallId3: placeSmallIds[2],
			Date:          LostDate,
			TimeSession:   Time2Session(),
		}
	} else if ctx.PostForm("place_2") != "" && ctx.PostForm("place_3") == "" {
		newLost = dbModel.Lost{
			Validity:      true,
			IsFoundBySelf: false,
			OpenId:        OpenId,
			Name:          user.Name,
			MatchId:       0,
			TypeBigId:     typeSmall.BigId,
			TypeSmallId:   typeSmall.ID,
			PlaceSmallId1: placeSmallIds[0],
			PlaceSmallId2: placeSmallIds[1],
			Date:          LostDate,
			TimeSession:   Time2Session(),
		}
	} else if ctx.PostForm("place_2") == "" && ctx.PostForm("place_3") == "" {
		newLost = dbModel.Lost{
			Validity:      true,
			IsFoundBySelf: false,
			OpenId:        OpenId,
			Name:          user.Name,
			MatchId:       0,
			TypeBigId:     typeSmall.BigId,
			TypeSmallId:   typeSmall.ID,
			PlaceSmallId1: placeSmallIds[0],
			Date:          LostDate,
			TimeSession:   Time2Session(),
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 415,
			"data": "",
			"msg":  "请按顺序将地点1,地点2,地点3填写!",
		})
		return
	}
	db.Create(&newLost)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"lost": newLost,
			"type":typeBig.Name+" "+typeSmall.Name,
			"place": placeSmallNames,
		},
		"msg":  "添加Lost成功，之后有疑似您的Found被建立后，我们将在第一时间通知您！",
	})

	// 消息卡片：
	timeSessionInChinese := ""
	switch newLost.TimeSession {
	case "morning": timeSessionInChinese = "上午（6：00-11：00）"
	case "noon": timeSessionInChinese = "中午（11：00-2：00）"
	case "afternoon": timeSessionInChinese = "下午（2：00-19：00）"
	case "evening": timeSessionInChinese = "晚上（19：00-22：00）"
	case "night": timeSessionInChinese = "夜间（00：00-6：00，22：00-24：00）"
	}
	cardMessage.SendCardMessage(
		OpenId,
		cardMessage.LostAddedCard(cardMessage.LostAdded{
			LostId: 	strconv.Itoa(int(newLost.ID)),
			ItemSubtype: typeBig.Name+" "+typeSmall.Name,
			LostDate:    newLost.Date +" "+timeSessionInChinese,
			LostPlace:   placeSmallNames[0]+" "+placeSmallNames[1]+" "+placeSmallNames[2],
		}),
	)
}