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
	var matches  []dbModel.Match
	OpenId := ctx.MustGet("open_id").(string)
	db.Where("open_id=?",OpenId).Find(&founds)
	db.Where("open_id=? AND match_id=?",OpenId,0).Find(&losts)
	db.Where("loster_open_id=?",OpenId).Find(&matches)
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
			ItemType:  common.TypeId2Name(value.TypeSmallId),
			Image:     value.Image,
		}
		myFounds = append(myFounds, tempFound)
	}
	// 查找未找到或已被自己找到
	for _, value := range losts {
		var isMatched bool
		imageUrl := ""
		if value.IsFoundBySelf == false{
			imageUrl = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=1618572354,3586828955&fm=26&gp=0.jpg"
		} else if value.IsFoundBySelf == true {
			isMatched = true
			lostCnt ++
			imageUrl = "https://www.fengzigeng.com/api/image?name=7d016e6add66f758c225c0454653797f.png"
		} else {
			continue
		}
		tempLost := myLost{
			LostID:    value.ID,
			IsMatched: isMatched,
			ItemType:  common.TypeId2Name(value.TypeSmallId),
			Image:     imageUrl,
		}
		myLosts = append(myLosts, tempLost)
	}
	// 查找已找到
	for _,value := range matches{
		println(value.TypeName)
		tempLost := myLost{
			LostID:    value.ID,
			IsMatched: true,
			ItemType:  common.TypeId2Name(value.TypeSmallId),
			Image:     value.Image,
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
