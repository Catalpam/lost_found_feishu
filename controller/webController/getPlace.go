package webController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
	"strconv"
)

func GetPlaces(ctx *gin.Context) {
	campusId := ctx.Query("campus_id")
	db := common.GetDB()
	var places []dbModel.Place
	var retPlace []RetPlace

	if campusId != "" {
		db.Where("campus_id = ?", campusId).Order("place_id ASC").Find(&places)
	} else {
		db.Find(&places)
	}

	for _, itemClass := range places {
		var campus string = ""
		var subPlaces []string
		json.Unmarshal([]byte(itemClass.Subareas), &subPlaces)

		if itemClass.CampusId == "0" {
			campus = "清水河校区"
		} else if itemClass.CampusId == "1" {
			campus = "沙河校区"
		}
		println(subPlaces)
		for subKey, subName := range subPlaces {
			var keyIndex []int
			PlaceIndexInt, _ := strconv.Atoi(itemClass.PlaceId)
			CampusIdInt, _ := strconv.Atoi(itemClass.CampusId)
			keyIndex = append(keyIndex,CampusIdInt)
			keyIndex = append(keyIndex,PlaceIndexInt)
			keyIndex = append(keyIndex,subKey)
			retPlace = append(retPlace, RetPlace{
				Campus:   campus,
				Place:    itemClass.Name,
				SubPlace: subName,
				Key:      keyIndex,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": retPlace,
		"msg":  "Place获取成功",
	})
}

type RetPlace struct {
	Campus 		string
	Place 		string
	SubPlace 	string
	Key        	[]int
}
