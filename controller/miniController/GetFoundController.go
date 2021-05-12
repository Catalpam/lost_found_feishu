package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strings"
)

func GetFoundList(ctx *gin.Context) {
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
	var FoundList []FoundListModel
	for _, value := range *founds {
		tempFound := FoundListModel{
			ID:        string(value.ID),
			SubType:   value.SubType,
			Campus:    value.Campus,
			Place:     value.Place,
			Image:	   value.ImageHome,
			FoundDate: value.FoundDate,
			FoundTime: value.FoundTime,
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
	ID string
	SubType string
	// Location
	Campus string
	Place string
	// Image
	Image string
	// Time
	FoundDate string
	FoundTime string
}