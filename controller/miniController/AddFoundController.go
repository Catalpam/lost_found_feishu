package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/comander"
	"lost_found/common"
	"lost_found/dbModel"
	"lost_found/handler"
	"net/http"
	"strings"
	"time"
)

func AddFound(ctx *gin.Context)  {
	println("Add Found")
	db := common.GetDB()
	//获取参数
	typeIndex, errType := String2Index(ctx.PostForm("type_index"))
	placeIndex, errPlace := String2Index(ctx.PostForm("place_index"))
	if errType != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "物品类型参数格式不合法!",
		})
		return
	}
	if errPlace != nil{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "地点参数格式不合法!",
		})
		return
	}

	itemInfo, _ := ctx.GetPostForm("info")
	image, _ := ctx.GetPostForm("image")
	campusId, _ := ctx.GetPostForm("campus_id")
	placeDetail,_ := ctx.GetPostForm("place_detail")
	currentPlace := ctx.PostForm("current_place")
	println("currentPlace"+currentPlace)
	currentPlaceDetail := ctx.PostForm("current_place_detail")
	losterInfo, _ := ctx.GetPostForm("loster_info")
	additionalInfo,_ := ctx.GetPostForm("additional_info")

	//检查有无空参数
	{
		if itemInfo == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：info",
			})
			return
		}
		if image == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：image",
			})
			return
		}

		if campusId == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：campus_id",
			})
			return
		}

	}

	//查找index对应的值是否存在于数据库中
	var typeBig   dbModel.TypeBig
	var typeSmall dbModel.TypeSmall
	db.Where("indexx=?", typeIndex[0]).First(&typeBig)
	db.Where("indexx=? AND big_id=?", typeIndex[1], typeBig.ID).First(&typeSmall)

	if typeSmall.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的type_index不存在!",
		})
		return
	}

	var placeBig   dbModel.PlaceBig
	var placeSmall dbModel.PlaceSmall
	db.Where("indexx=? AND campus_id=?", placeIndex[0],campusId).First(&placeBig)
	db.Where("indexx=? AND big_id=?", placeIndex[1], placeBig.ID).First(&placeSmall)

	if typeSmall.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的place_index不存在!",
		})
		return
	}

	imageNameList := strings.Split(image,`image?name=`)
	imageKey, errUpload := handler.Uploadimage2Feishu(imageNameList[1])
	if errUpload != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": image,
			"msg":  "ImageUrl必须是之前上传过的图片!",
		})
		println(errUpload)
		return
	}

	//获取用户OpenId
	OpenId := ctx.MustGet("open_id").(string)
	//将新的Found对象添加至数据库中
	newFound := dbModel.Found{
		Date:          		time.Now().Format("2006-01-02"),
		Time:          		time.Now().Format("15:04"),
		TimeSession:   		Time2Session(),
		OpenId:        		OpenId,
		TypeBigId:          typeBig.ID,
		TypeSmallId:        typeSmall.ID,
		PlaceBigId:         placeBig.ID,
		PlaceSmallId:       placeSmall.ID,
		ItemInfo:           itemInfo,
		Image:              image,
		ImageKey: 			imageKey,
		PlaceDetail:        placeDetail,
		CurrentPlace:       currentPlace,
		CurrentPlaceDetail: currentPlaceDetail,
		LosterInfo:         losterInfo,
		AdditionalInfo:     additionalInfo,
	}

	db.Create(&newFound)
	// 新建一个匹配进程，若有符合匹配将会使用机器人发送
	go comander.CheckNewFoundIsMatchedLost(newFound.ID)
	// 此进程继续返回创立Found成功的Response
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": newFound,
		"msg":  "添加Found成功，详见data",
	})
}

func Time2Session() string  {
	currentTime := time.Now().Hour()
	var currentSession string

	switch currentTime {
	case 6,7,8,9,10:
		currentSession = "morning"
	case 11,12,13:
		currentSession = "noon"
	case 14,15,16,17,18:
		currentSession = "afternoon"
	case 19,20,21:
		currentSession = "evening"
	case 22,23,1,2,3,4,5:
		currentSession = "night"
	}
	return currentSession
}