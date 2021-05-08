package controller

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strings"
)

func AddFound(ctx *gin.Context)  {
	db := common.GetDB()

	//获取参数
	typeIndex,_ := ctx.GetPostForm("type_index")
	itemInfo, _ := ctx.GetPostForm("info")
	image, _ := ctx.GetPostForm("image")
	place, _ := ctx.GetPostForm("place")
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

		//地点位置解析
		if place == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 413,
				"data": "",
				"msg":  "缺少参数：place",
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
	var tempClass dbModel.ItemClass
	db.Where("class_id = ?", thing.ClassId).First(&tempClass)
	ClassName :=  tempClass.ClassName
	TypeName := thing.Type

	//将新的Found对象添加至数据库中
	newFound := dbModel.Found{
		TypeName:       TypeName,
		ClassName:      ClassName,
		ItemInfo:       itemInfo,
		Image:          image,
		Place:          place,
		PlaceDetail:    placeDetail,
		LosterInfo:     losterInfo,
		CurrentPlace:   currentPlace,
		AdditionalInfo: additionalInfo,
	}
	db.Create(&newFound)
}

func GetFound(ctx *gin.Context) {
	DB := common.GetDB()

	//获取参数
	typeId,_ := ctx.GetPostForm("type_id")
	itemInfo, _ := ctx.GetPostForm("info")
	image, _ := ctx.GetPostForm("image")
	place, _ := ctx.GetPostForm("place")
	placeDetail, _ := ctx.GetPostForm("place_detail")
	losterInfo, _ := ctx.GetPostForm("loster_info")
	currentPlace,_ := ctx.GetPostForm("current_place")
	additionalInfo,_ := ctx.GetPostForm("additional_info")

	//检查有无空参数
	if typeId == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少找到物品的类型：type_name_id",
		})
		return
	}
	if itemInfo == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：info",
		})
		return
	}
	if image == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：image",
		})
		return
	}
	if place == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：place",
		})
		return
	}
	if placeDetail == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：place_detail",
		})
		return
	}
	if losterInfo == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：loster_info",
		})
		return
	}
	if currentPlace == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：current_place",
		})
		return
	}
	if additionalInfo == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "缺少参数：additional_info",
		})
		return
	}


	//查找类型ID对应的类型属性是否存在
	var thing dbModel.Type
	DB.Where("type_id = ?", typeId).First(&thing)
	if thing.ID == 0{
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg": "参数不合法，type_id不存在",
		})
		return
	}

	//将新的Found对象添加至数据库中
	newFound := dbModel.Found{
		TypeName:       typeId,
		ItemInfo:       itemInfo,
		Image:          image,
		Place:          place,
		PlaceDetail:    placeDetail,
		LosterInfo:     losterInfo,
		CurrentPlace:   currentPlace,
		AdditionalInfo: additionalInfo,
	}
	DB.Create(&newFound)
}

func GetTypes(ctx *gin.Context)  {
	db := common.GetDB()
	var text = "[["
	var types []dbModel.Type
	var classes []dbModel.ItemClass
	db.Order("class_id ASC").Find(&classes)
	println(classes[0].ClassName)
	for _, itemClass := range classes{
			text = text + itemClass.ClassName + ","
	}
	text = text + "],"

	text = text + "["
	for _, itemClass := range classes{
		text = text + "["
		db.Where("class_id = ?",itemClass.ClassId).Order("type_id ASC").Find(&types)
		for _, value := range types{
				text =  text + value.Type + ","
		}
		text = text + "],"
	}
	text = text + "],"

	text = text + "]"
	ctx.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data": text,
		"msg": "物品类型Types返回成功",
	})
}