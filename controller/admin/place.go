package admin

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetPlaces(ctx *gin.Context)  {
	db := common.GetDB()
	var bigs   []dbModel.PlaceBig
	var retOptions [2]retCampus
	retOptions[0] = retCampus{
		CampusId: "0",
		Campus:   "清水河校区",
	}
	retOptions[1] = retCampus{
		CampusId: "1",
		Campus:   "沙河校区",
	}
	db.Order("campus_id ASC").Order("id ASC").Find(&bigs)

	for _, big := range bigs {
		var smalls []dbModel.PlaceSmall
		if big.CampusId == "0" {
			db.Where("big_id=?", big.ID).Order("campus_id ASC").Order("indexx ASC").Find(&smalls)
			var retSmalls []retPlace
			for _,small := range smalls {
				retSmalls = append(retSmalls,retPlace{
					Place:    small.Name,
					Key:   small.ID,
				})
			}
			retOptions[0].Children = append(retOptions[0].Children,retPlaceBig{
				Place:    big.Name,
				BigKey:   big.ID,
				Children: retSmalls,
			})
		} else if big.CampusId == "1" {
			db.Where("big_id=?", big.ID).Order("campus_id ASC").Order("indexx ASC").Find(&smalls)
			var retSmalls []retPlace
			for _,small := range smalls {
				retSmalls = append(retSmalls,retPlace{
					Place:    small.Name,
					Key:   small.ID,
				})
			}
			retOptions[1].Children = append(retOptions[1].Children,retPlaceBig{
				Place:    big.Name,
				BigKey:   big.ID,
				Children: retSmalls,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": retOptions,
		"msg":  "BigPlaces获取成功",
	})
}

type retPlace struct {
	Place 		string `json:"label"`
	Key        	uint   `json:"value"`
}

type retPlaceBig struct {
	Place 		string `json:"label"`
	BigKey      uint   `json:"value"`
	Children 	[]retPlace `json:"children"`
}

type retCampus struct {
	CampusId string `json:"value"`
	Campus 	 string `json:"label"`
	Children []retPlaceBig `json:"children"`
}
