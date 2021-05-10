package miniController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
	"strings"
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


	//将新的Found对象添加至数据库中
	newFound := dbModel.Found{
		ItemType:       itemType.Name,
		SubType:        SubTypeName,
		Campus:         campus.Name,
		Place:          place.Name,
		SubPlace:       subPlace,
		ItemInfo:       itemInfo,
		Image:          image,
		PlaceDetail:    placeDetail,
		CurrentPlace:   currentPlace,
		LosterInfo:     losterInfo,
		AdditionalInfo: additionalInfo,
	}
	db.Create(&newFound)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": newFound,
		"msg":  "添加Found成功，详见data",
	})
}

func GetFound(ctx *gin.Context) {
	db := common.GetDB()
	var founds []dbModel.Found
	db.Find(&founds)
	returnFounds(&founds,ctx)
}

func Get_Found(ctx *gin.Context) {
	db := common.GetDB()
	var founds []dbModel.Found

	typeIndex,_ := ctx.GetPostForm("type_index")
	//placeIndex,_ := ctx.GetPostForm("place_index")
	//timeSession,_ := ctx.GetPostForm("time_session")

	//获取物品大类
	//获取物品小类
	//获取校区
	//获取丢失地点大类
	//获取丢失地点小类
	//获取时段

	//没有输入TypeId时的返回
	if typeIndex == "" {
		db.Find(&founds)
	} else {
		//查找TypeId对应的类型属性存在于数据库中
		TypeId:= ""
		str_arr :=  strings.Split(typeIndex, `,`)
		str0 := strings.Split(str_arr[0], `{`)
		str1 := strings.Split(str_arr[1], `}`)
		for _, str := range str0 {
			TypeId = TypeId + str
		}
		for _, str := range str1 {
			TypeId = TypeId + str
		}
		println(TypeId)
		var thing dbModel.Type
		db.Where("type_id = ?", TypeId).Order("type_id ASC").First(&thing)
		if thing.ID == 0{
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": TypeId,
				"msg":  "参数不合法，type_id不存在",
			})
			return
		}
		println("----------------thing.Type获取成功！---------------------")
		println(thing.Type)
		db.Where("type_name = ?", thing.Type).Find(&founds)
	}
	returnFounds(&founds,ctx)
}

func returnFounds(founds *[]dbModel.Found, ctx *gin.Context)  {
	if len(*founds) == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "没有查询到符合条件的Founds",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": *founds,
		"msg":  "获取Found成功",
	})
}
