package miniController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/comander"
	"lost_found/common"
	"lost_found/dbModel"
	"lost_found/handler"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AddFound(ctx *gin.Context)  {
	db := common.GetDB()
	//获取参数
	typeIndex,_ := ctx.GetPostForm("type_index")
	itemInfo, _ := ctx.GetPostForm("info")
	image, _ := ctx.GetPostForm("image")
	campus_id, _ := ctx.GetPostForm("campus_id")
	placeIndex, _ := ctx.GetPostForm("place_index")
	placeDetail,_ := ctx.GetPostForm("place_detail")
	losterInfo, _ := ctx.GetPostForm("loster_info")
	currentPlace,_ := ctx.GetPostForm("current_place")
	currentPlaceDetail,_ := ctx.GetPostForm("current_place_detail")
	additionalInfo,_ := ctx.GetPostForm("additional_info")

	//检查有无空参数
	{
		if typeIndex == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少找到物品的类型：type_index",
			})
			return
		}
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

		if campus_id == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：campus_id",
			})
			return
		}

		//地点位置解析
		if placeIndex == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：place_index",
			})
			return
		}
		if placeDetail == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：place_detail",
			})
			return
		}
		//参数为0，1，2
		if currentPlace == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：current_place",
			})
			return
		}
		if currentPlace == "2" {
			if currentPlaceDetail == "" {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 413,
					"data": "",
					"msg":  "缺少参数：current_place_detail",
				})
				return
			}
		} else if (currentPlace != "0") && (currentPlace != "1") {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 413,
					"data": "",
					"msg":  "current_place不合法",
				})
				return
		}
	}

	//查找TypeId对应的类型属性是是否存在于数据库中
	index1 := ""
	index2 := ""
	str_arr := strings.Split(typeIndex, `,`)
	str0 := strings.Split(str_arr[0], `[`)
	str1 := strings.Split(str_arr[1], `]`)
	for _, str := range str0 {
		index1 = index1 + str
	}
	for _, str := range str1 {
		index2 = index2 + str
	}
	id_2, err2 := strconv.Atoi(index2)
	if err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "参数格式不合法!",
		})
		return
	}
	var itemType dbModel.ItemType
	var SubTypeName string
	db.Where("type_id = ?", index1).First(&itemType)
	var subtypes []string
	_ = json.Unmarshal([]byte(itemType.Subtypes), &subtypes)
	println("--------------" + subtypes[0] + "----------------------")
	if id_2 > (len(subtypes) - 1) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的type_index不存在!",
		})
		return
	} else {
		SubTypeName = subtypes[id_2]
	}


	//获取place信息
	var subPlace = ""
	index1 = ""
	index2 = ""
	str_arr2 := strings.Split(placeIndex, `,`)
	str0 = strings.Split(str_arr2[0], `[`)
	str1 = strings.Split(str_arr2[1], `]`)
	for _, str := range str0 {
		index1 = index1 + str
	}
	println("--------------" + "index1:"+ index1 + "----------------------")
	for _, str := range str1 {
		index2 = index2 + str
	}
	println("--------------" + "index2:"+ index2 + "----------------------")
	id_2, err2 = strconv.Atoi(index2)
	if err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的place_index不合法!",
		})
		return
	}
	var place dbModel.Place
	db.Where("place_id =? AND campus_id=?", index1,campus_id).First(&place)
	var subareas []string
	_ = json.Unmarshal([]byte(place.Subareas), &subareas)
	println("--------------" + "断点！！！" + "----------------------")
	if id_2 > (len(subareas) - 1) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求的校区中place_index不存在!",
		})
		return
	} else {
		subPlace = subareas[id_2]
	}

	var campus dbModel.Campus
	db.Where("campus_id=?",campus_id).First(&campus)

	//解析image数组
	//var imageList []string
	//imageListErr := json.Unmarshal([]byte(image), &imageList)
	//if imageListErr != nil {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code": 413,
	//		"data": "",
	//		"msg":  "Image List格式不合法!",
	//	})
	//	return
	//}
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
		FoundDate:          time.Now().Format("2006-01-02"),
		FoundTime:          time.Now().Format("15:04"),
		FoundTimeSession:   Time2Session(),
		FoundOpenId:        OpenId,
		ItemType:           itemType.Name,
		SubType:            SubTypeName,
		Campus:             campus.Name,
		Place:              place.Name,
		SubPlace:           subPlace,
		ItemInfo:           itemInfo,
		Image:              image,
		ImageHome:          image,
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