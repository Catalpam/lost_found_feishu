package common

import (
	"github.com/gin-gonic/gin"
	"lost_found/dbModel"
	"net/http"
)

func ItemTypeInitial(ctx *gin.Context) {
	db := GetDB()
	db.Exec("drop table type_bigs")
	db.Exec("drop table type_smalls")
	db.AutoMigrate(&dbModel.TypeSmall{})
	db.AutoMigrate(&dbModel.TypeBig{})

	var subareasList0 [][]string
	var areas0 = [4]string{"电子设备","证件或证书","个人物品","学习用品"}
	subareasList0 = append(subareasList0,[]string{"手机", "平板电脑", "电纸书", "笔记本电脑", "耳机", "充电器", "充电宝", "数据线","智能手表", "U盘", "Apple Pencil", "鼠标", "键盘", "其他电子设备"})
	subareasList0 = append(subareasList0,[]string{"一卡通", "身份证", "学生证", "健身卡","考试通过证书","比赛奖状","其他证件或证书"})
	subareasList0 = append(subareasList0,[]string{"衣物（含帽子手套围巾等）", "包", "眼镜", "钱包", "钥匙", "水杯", "化妆品", "首饰", "笔袋", "未领走快递", "雨伞", "运动器械相关","其它个人物品"})
	subareasList0 = append(subareasList0,[]string{"笔袋", "书籍", "纸质笔记本","其它学习物品"})
	for key,value := range areas0 {
		var placeBig dbModel.TypeBig
		var newPlaceBig dbModel.TypeBig
		db.Order("indexx DESC").Limit(1).Find(&placeBig)
		db.Where("name=?",value).First(&newPlaceBig)
		var index uint = 0
		if placeBig.ID != 0 {
			index = placeBig.Indexx + 1
		}
		if newPlaceBig.ID == 0 {
			newPlaceBig = dbModel.TypeBig{
				Indexx: index,
				Name:   value,
			}
			db.Create(&newPlaceBig)
		}
		for _, subvalue := range subareasList0[key] {
			var placeSmall dbModel.TypeSmall
			var newPlaceSmall dbModel.TypeSmall
			db.Where("big_id=?",newPlaceBig.ID).Order("indexx DESC").Limit(1).Find(&placeSmall)
			db.Where("big_id=? AND name=?",newPlaceBig.ID,subvalue).First(&newPlaceSmall)
			if newPlaceSmall.ID == 0 {
				var indexx uint = 0
				if placeSmall.ID != 0 {
					indexx = placeSmall.Indexx + 1
				}
				db.Create(&dbModel.TypeSmall{
					Indexx:   indexx,
					Name:     subvalue,
					BigId:    newPlaceBig.ID,
					BigName:  newPlaceBig.Name,
				})
			}
		}
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code" : 200,
		"msg"  : "物品类型成功恢复为默认值。",
	})
}