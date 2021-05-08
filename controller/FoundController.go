package controller

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func AddFound(ctx *gin.Context)  {
	DB := common.GetDB()

	//获取参数
	typeId,_ := ctx.GetPostForm("type_id")
	itemInfo, _ := ctx.GetPostForm("info")
	image, _ := ctx.GetPostForm("image")
	place, _ := ctx.GetPostForm("place")
	placeDetail,_ := ctx.GetPostForm("place_detail")
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
		TypeId:   	    typeId,
		ItemInfo:  	    itemInfo,
		Image:      	image,
		Place:       	place,
		PlaceDetail: 	placeDetail,
		LosterInfo:  	losterInfo,
		CurrentPlace:   currentPlace,
		AdditionalInfo: additionalInfo,
	}
	DB.Create(&newFound)
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
		TypeId:   	    typeId,
		ItemInfo:  	    itemInfo,
		Image:      	image,
		Place:       	place,
		PlaceDetail: 	placeDetail,
		LosterInfo:  	losterInfo,
		CurrentPlace:   currentPlace,
		AdditionalInfo: additionalInfo,
	}
	DB.Create(&newFound)
}