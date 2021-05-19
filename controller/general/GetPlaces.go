package general

import (
	"encoding/json"
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
	var place []dbModel.Place
	db.Where("campus_id = ?", campusId).Order("place_id ASC").Find(&place)
	println(place[0].Name)
	var placeBig []string
	var placeSmall []placeSmallIndex
	for _, itemClass := range place {
		placeBig = append(placeBig, itemClass.Name)
		var index placeSmallIndex
		json.Unmarshal([]byte(itemClass.Subareas), &index)
		placeSmall = append(placeSmall,index)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"place1": placeBig,
			"place2":placeSmall,
		},
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
		text = text + `"` + campus.Name + `"` + ","
	}
	text = text + "]"
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": json.RawMessage(text),
		"msg":  "Campus返回成功",
	})
}

type placeSmallIndex []string
