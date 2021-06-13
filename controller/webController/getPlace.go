package webController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetPlaces(ctx *gin.Context) {
	campusId := ctx.Query("campus_id")
	db := common.GetDB()
	var placeSmalls []dbModel.PlaceSmall
	var retPlace []RetPlace
	db.Where("campus_id = ?", campusId).Order("campus_id ASC").Order("indexx ASC").Find(&placeSmalls)

	if campusId != "" {
	} else {
		db.Find(&placeSmalls)
	}

	for _, value := range placeSmalls {
		var campus string = ""
		if value.CampusId == "0" {
			campus = "清水河校区"
		} else if value.CampusId == "1" {
			campus = "沙河校区"
		}
		retPlace = append(retPlace, RetPlace{
			Campus:   campus,
			Place:    value.BigName,
			SubPlace: value.Name,
			Key:      value.ID,
			BigKey:	  value.BigId,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": retPlace,
		"msg":  "Place获取成功",
	})
}

func GetPlacesBig(ctx *gin.Context) {
	db := common.GetDB()
	var placeBigs []dbModel.PlaceBig
	var retOptions [2]AddSmallOptionSingle
	retOptions[0] = AddSmallOptionSingle{
		CampusId: "0",
		Campus:   "清水河校区",
	}
	retOptions[1] = AddSmallOptionSingle{
		CampusId: "1",
		Campus:   "沙河校区",
	}
	db.Order("campus_id ASC").Order("indexx ASC").Find(&placeBigs)

	for _, value := range placeBigs {
		if value.CampusId == "0" {
			retOptions[0].Children = append(retOptions[0].Children,RetPlaceBig{
				Place:  value.Name,
				BigKey: value.ID,
			})
		} else if value.CampusId == "1" {
			retOptions[1].Children = append(retOptions[1].Children,RetPlaceBig{
				Place:  value.Name,
				BigKey: value.ID,
			})
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": retOptions,
		"msg":  "BigPlaces获取成功",
	})
}


type RetPlace struct {
	Campus 		string
	Place 		string
	SubPlace 	string
	Key        	uint
	BigKey		uint
}

type RetPlaceBig struct {
	Place 		string `json:"label"`
	BigKey      uint   `json:"value"`
}

type AddSmallOptionSingle struct {
	CampusId string `json:"value"`
	Campus 	 string `json:"label"`
	Children []RetPlaceBig `json:"children"`
}
