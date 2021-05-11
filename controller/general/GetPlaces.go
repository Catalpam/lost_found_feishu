package general

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetPlaces(ctx *gin.Context) {
	campusId := ctx.PostForm("campus_id")
	if campusId == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 413,
			"data": "",
			"msg":  "campus_id为空！",
		})
	}
	db := common.GetDB()
	var text = "[["
	var place []dbModel.Place
	db.Where("campus_id = ?", campusId).Order("place_id ASC").Find(&place)
	println(place[0].Name)
	for _, itemClass := range place {
		text = text + itemClass.Name + ","
	}
	text = text + "],"

	text = text + "["
	for _, placeItem := range place {
		text = text + placeItem.Subareas + ","
	}
	text = text + "]"

	text = text + "]"
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": text,
		"msg":  "Places返回成功",
	})
}

func GetCampus(ctx *gin.Context) {
	db := common.GetDB()
	var text = "["
	var campuses []dbModel.Campus
	db.Order("campus_id ASC").Find(&campuses)
	println(campuses[0].Name)
	for _, campus := range campuses {
		text = text + campus.Name + ","
	}
	text = text + "]"
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": text,
		"msg":  "Campus返回成功",
	})
}
