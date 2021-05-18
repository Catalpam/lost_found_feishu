package miniController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/cardMessage"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
	"strings"
)

func AddLost(ctx *gin.Context) {
	db := common.GetDB()

	//获取参数
	typeIndex, _ := ctx.GetPostForm("type_index")
	campusId, _ := ctx.GetPostForm("campus_id")
	place1, _ := ctx.GetPostForm("place_1")
	place2, _ := ctx.GetPostForm("place_2")
	place3, _ := ctx.GetPostForm("place_3")
	LostDate := ctx.PostForm("lost_date")
	timeSession, _ := ctx.GetPostForm("time_session")

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

	var PlaceCom1 string = ""
	var PlaceCom2 string = ""
	var PlaceCom3 string = ""

	//查找地理位置是否合法
	var placeCnt int
	var placeHasErr = false
	if place1 == "" {
		placeCnt = 0
	} else if place2 == "" {
		placeCnt = 1
		PlaceComByte1, _ := json.Marshal(CheckPlace(place1, campusId, &placeHasErr))
		PlaceCom1 = string(PlaceComByte1)

	} else if place3 == "" {
		placeCnt = 2
		PlaceComByte1, _ := json.Marshal(CheckPlace(place1, campusId, &placeHasErr))
		PlaceComByte2, _ := json.Marshal(CheckPlace(place2, campusId, &placeHasErr))
		PlaceCom1 = string(PlaceComByte1)
		PlaceCom2 = string(PlaceComByte2)
		if place1 == place2 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 414,
				"data": "",
				"msg":  "Place有重复!",
			})
			return
		}
	} else {
		placeCnt = 3
		PlaceComByte1, _ := json.Marshal(CheckPlace(place1, campusId, &placeHasErr))
		PlaceComByte2, _ := json.Marshal(CheckPlace(place2, campusId, &placeHasErr))
		PlaceComByte3, _ := json.Marshal(CheckPlace(place3, campusId, &placeHasErr))
		PlaceCom1 = string(PlaceComByte1)
		PlaceCom2 = string(PlaceComByte2)
		PlaceCom3 = string(PlaceComByte3)
		if place1 == place2 || place1 == place3 || place2 == place3 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 414,
				"data": "",
				"msg":  "Place有重复!",
			})
			return
		}
	}
	if placeCnt == 0 || placeHasErr == true {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 415,
			"data": "",
			"msg":  "Place参数格式不合法!",
		})
		return
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
	SubTypeName := ""
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

	var campus dbModel.Campus
	db.Where("campus_id=?", campusId).First(&campus)

	//获取用户OpenId
	OpenId := ctx.MustGet("open_id").(string)
	//将新的Found对象添加至数据库中
	newLost := dbModel.Lost{
		LosterOpenId:    OpenId,
		TypeSubName:     SubTypeName,
		LostPlace1:      PlaceCom1,
		LostPlace2:      PlaceCom2,
		LostPlace3:      PlaceCom3,
		LostDate:        LostDate,
		LostTimeSession: timeSession,
	}
	db.Create(&newLost)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": newLost,
		"msg":  "添加Lost成功，之后有疑似您的Found被建立后，我们将在第一时间通知您！",
	})


	// 消息卡片：
	var temPlaceStr []string
	_ = json.Unmarshal([]byte(newLost.LostPlace1), &temPlaceStr)
	sendPlaceStr := temPlaceStr[1]
	if newLost.LostPlace2 != "" {
		var temPlaceStr []string
		_ = json.Unmarshal([]byte(newLost.LostPlace2), &temPlaceStr)
		sendPlaceStr = sendPlaceStr + "," + temPlaceStr[1]
	}
	if newLost.LostPlace3 != "" {
		var temPlaceStr []string
		_ = json.Unmarshal([]byte(newLost.LostPlace3), &temPlaceStr)
		sendPlaceStr = sendPlaceStr + "," + temPlaceStr[1]
	}
	timeSessionInChinese := ""
	switch newLost.LostTimeSession {
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
			ItemSubtype: newLost.TypeSubName,
			LostDate:    newLost.LostDate +" "+timeSessionInChinese,
			LostPlace:   sendPlaceStr,
		}),
	)
}

//获取Place信息
func CheckPlace(placeIndex string, campusId string, hasErr *bool) []string {
	db := common.GetDB()
	index1 := ""
	index2 := ""
	str_arr2 := strings.Split(placeIndex, `,`)
	str0 := strings.Split(str_arr2[0], `[`)
	str1 := strings.Split(str_arr2[1], `]`)
	for _, str := range str0 {
		index1 = index1 + str
	}
	println("--------------" + "index1:" + index1 + "----------------------")
	for _, str := range str1 {
		index2 = index2 + str
	}
	println("--------------" + "index2:" + index2 + "----------------------")
	id_2, err2 := strconv.Atoi(index2)
	if err2 != nil {
		*hasErr = true
		return []string{"", ""}
	}
	var place dbModel.Place
	db.Where("place_id =? AND campus_id=?", index1, campusId).First(&place)
	var subareas []string
	_ = json.Unmarshal([]byte(place.Subareas), &subareas)
	println("--------------" + "断点！！！" + "----------------------")
	if id_2 > (len(subareas) - 1) {
		*hasErr = true
		return []string{"", ""}
	} else {
		return []string{place.Name, subareas[id_2]}
	}
}
