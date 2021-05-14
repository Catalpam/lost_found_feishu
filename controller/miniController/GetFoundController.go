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

func GetFoundList(ctx *gin.Context) {
	db := common.GetDB()
	var founds []dbModel.Found
	// 获取Form中的参数 FoundId
	FoundIdStr := ctx.PostForm("id")
	// 查找参数
	if FoundIdStr != "" {
		FoundId, err := strconv.ParseUint(ctx.PostForm("id"), 10 ,32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": err,
				"msg":  "id格式不合法！",
			})
		}
		println(FoundId)
		db.Where("id=?",FoundId).Find(&founds)
	} else {
		SelectFound(&founds,ctx)
	}
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
	var FoundList []FoundListModel
	for _, value := range *founds {
		tempFound := FoundListModel{
			ID:        value.ID,
			SubType:   value.SubType,
			Campus:    value.Campus,
			Place:     value.Place,
			Image:	   value.ImageHome,
			FoundDate: value.FoundDate,
			FoundTime: value.FoundTime,
			ItemInfo: value.ItemInfo,
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
	SubType string
	// Location
	Campus string
	Place string
	// Image
	Image string
	// Time
	FoundDate string
	FoundTime string
	ItemInfo string
}

func SelectFound(founds *[]dbModel.Found, ctx *gin.Context)  {
	db := common.GetDB()
	//获取参数
	typeIndex,_ := ctx.GetPostForm("type_index")
	campus_id, _ := ctx.GetPostForm("campus_id")
	placeIndex, _ := ctx.GetPostForm("place_index")
	date, _ := ctx.GetPostForm("date")
	timeSession := ctx.PostForm("time_session")


	//初始化参数
	SubTypeName  := ""
	subPlace := ""

	//查找TypeId对应的类型属性是是否存在于数据库中
	if typeIndex != "" {
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
	}
	if campus_id == "" && placeIndex != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "请求带place_index时请带上campus_id!",
		})
		return
	}
	if campus_id != "" && placeIndex != "" {
		//获取place信息
		index1 := ""
		index2 := ""
		str_arr2 := strings.Split(placeIndex, `,`)
		str0 := strings.Split(str_arr2[0], `[`)
		str1 := strings.Split(str_arr2[1], `]`)
		for _, str := range str0 {
			index1 = index1 + str
		}
		println("--------------" + "index1:"+ index1 + "----------------------")
		for _, str := range str1 {
			index2 = index2 + str
		}
		println("--------------" + "index2:"+ index2 + "----------------------")
		id_2, err2 := strconv.Atoi(index2)
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
	}

	var campus dbModel.Campus
	db.Where("campus_id=?",campus_id).First(&campus)
	println("----------------准备开始查找符合条件的----------------------")
	//在数据库中查找Found对象
	//match_id=0
	println(campus.Name)
	db.Where(&dbModel.Found{
		SubType:            SubTypeName,
		Campus:             campus.Name,
		SubPlace:           subPlace,
		FoundDate:          date,
		FoundTimeSession:   timeSession,
		MatchId: 			0,
	}).Find(&founds)
}
