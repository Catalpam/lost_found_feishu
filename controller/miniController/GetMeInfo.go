package miniController

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/dbModel"
	"net/http"
)

func GetMeInfo(ctx *gin.Context)  {
	db := common.GetDB()
	var founds []dbModel.Found
	var losts  []dbModel.Lost
	OpenId := ctx.MustGet("open_id").(string)
	db.Where("found_open_id=?",OpenId).Find(&founds)
	db.Where("loster_open_id=?",OpenId).Find(&losts)
	var myFounds []myFound
	var myLosts 	 []myLost
	// 统计数量
	helpedCnt := 0
	lostCnt := 0
	for _, value := range founds {
		var isMatched bool
		if value.MatchId == 0{
			isMatched = false
		} else {
			isMatched = true
			helpedCnt ++
		}
		tempFound := myFound{
			FoundID:        value.ID,
			IsMatched: isMatched,
			ItemType:  value.SubType,
			Image:     value.Image,
		}
		myFounds = append(myFounds, tempFound)
	}
	for _, value := range losts {
		var isMatched bool
		var tempfound dbModel.Found
		if value.MatchId == 0{
			isMatched = false
		} else {
			isMatched = true
			lostCnt ++
			db.Where("match_id=?",value.ID).First(&tempfound)
		}
		tempLost := myLost{
			LostID:        value.ID,
			IsMatched: isMatched,
			ItemType:  value.TypeSubName,
			Image:     tempfound.Image,
		}
		myLosts = append(myLosts, tempLost)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"HasHelped":helpedCnt,
			"HasFound":lostCnt,
			"MyLosts":myLosts,
			"MyFounds":myFounds,
		},
		"msg": "返回成功",
	})
	return
}

type myFound struct {
	//gorm Model
	FoundID uint
	IsMatched bool
	ItemType string
	Image string
}
type myLost struct {
	//gorm Model
	LostID uint `gorm:"primary_key"`
	IsMatched bool
	ItemType string `gorm:"type:varchar(20);not null;"`
	Image string `gorm:"type:varchar(200);"`
}


//data: {
//	Lost:[{
//		Id:
//		ItemType:
//		Image:
//		isFound:
//	}],
//	Found:[{
//		Id:
//		ItemType:
//		Image:
//		isClaimed:
//	}],
//}
