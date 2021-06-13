package common

import (
	"github.com/gin-gonic/gin"
	"lost_found/dbModel"
	"net/http"
)

func PlaceInitial(ctx *gin.Context) {
	db := GetDB()
	db.Exec("drop table place_bigs")
	db.Exec("drop table place_smalls")
	db.AutoMigrate(&dbModel.PlaceSmall{})
	db.AutoMigrate(&dbModel.PlaceBig{})

	var subareasList0 [][]string
	var subareasList1 [][]string
	var areas0 = [7]string{"教学楼","硕丰苑","学知苑","博翰苑","留学生宿舍区","餐厅","其他"}
	var areas1 = [5]string{"教学楼","宿舍区","食堂","运动场所","其它"}
	subareasList0 = append(subareasList0,[]string{"研究院大楼","四号科研楼A区","四号科研楼B区","四号科研楼C区","五号科研楼","综合楼","磁共振脑成像中心","科研楼","基础实验楼","品学楼A区","品学楼B区","品学楼C区","立人楼"} )
	subareasList0 = append(subareasList0,[]string{"硕丰苑1栋","硕丰苑2栋","硕丰苑3栋","硕丰苑4栋","硕丰苑5栋","硕丰苑6栋","硕丰苑7栋","硕丰苑8栋","硕丰苑9栋","硕丰苑10栋","硕丰苑11栋","硕丰苑12栋","硕丰苑13栋","硕丰苑14栋","硕丰苑15栋","硕丰苑16栋","硕丰苑17栋","硕丰苑18栋","硕丰苑19栋","硕丰苑20栋","硕丰苑21栋","硕丰苑22栋","硕丰苑23栋","硕丰苑24栋","硕丰苑25栋","硕丰苑26栋","硕丰苑27栋","硕丰苑28栋"})
	subareasList0 = append(subareasList0,[]string{"学知苑1栋","学知苑2栋","学知苑3栋","学知苑4栋","学知苑5栋","学知苑6栋","学知苑7栋","学知苑8栋","学知苑9栋","学知苑10栋","学知苑11栋","学知苑12栋","学知苑13栋","学知苑14栋","学知苑15栋","学知苑16栋","学知苑17栋","学知苑18栋","学知苑19栋","学知苑20栋","学知苑21栋","学知苑22栋","学知苑23栋","学知苑24栋","学知苑25栋","学知苑26栋","学知苑27栋","学知苑28栋"})
	subareasList0 = append(subareasList0,[]string{"博翰苑1栋","博翰苑2栋","博翰苑3栋","博翰苑4栋","博翰苑5栋","博翰苑6栋","博翰苑7栋","博翰苑8栋","博翰苑9栋"})
	subareasList0 = append(subareasList0,[]string{"留学生2栋","留学生3栋","留学生4栋"})
	subareasList0 = append(subareasList0,[]string{"清真餐厅","学生二食堂","春晖苑（三食堂）"})
	subareasList0 = append(subareasList0,[]string{"体育馆","游泳馆","体育场","众创空间","主楼A1区","主楼A2区","主楼B1区","主楼B2区","主楼B3区","主楼C2区","主楼C1区","图书馆","校医院","学生活动中心","成电会堂","商业街","综合训练馆","游泳中心"})

	subareasList1 = append(subareasList1,[]string{"第二教学楼","逸夫楼","微固学院","光电学院","医学院","第三教学楼","第四教学楼","信软实验楼","通信大楼","211大楼"})
	subareasList1 = append(subareasList1,[]string{"校内16栋","校内17栋","校内18栋","校内12栋","校内15栋","校内14栋","校内13栋","欣苑"})
	subareasList1 = append(subareasList1,[]string{"阳光餐厅","万友餐厅","风华餐厅","桂苑餐厅","西北美食坊"})
	subareasList1 = append(subareasList1,[]string{"体育馆","羽毛球馆","健身房","篮球场","操场","乒乓球场","足球场","保卫处","游泳馆","网球场"})
	subareasList1 = append(subareasList1,[]string{"图书馆","主楼","学术交流中心","校医院住院部","校医院","沙河停车场"} )

	for key,value := range areas0 {
		var placeBig dbModel.PlaceBig
		var newPlaceBig dbModel.PlaceBig
		db.Where("campus_id=?","0").Order("indexx DESC").Limit(1).Find(&placeBig)
		db.Where("campus_id=? AND name=?","0",value).First(&newPlaceBig)
		var index uint = 0
		if placeBig.ID != 0 {
			index = placeBig.Indexx + 1
		}
		if newPlaceBig.ID == 0 {
			newPlaceBig = dbModel.PlaceBig {
				Indexx:   index,
				Name:     value,
				CampusId: "0",
			}
			db.Create(&newPlaceBig)
		}
		for _, subvalue := range subareasList0[key] {
			var placeSmall dbModel.PlaceSmall
			var newPlaceSmall dbModel.PlaceSmall
			db.Where("big_id=?",newPlaceBig.ID).Order("indexx DESC").Limit(1).Find(&placeSmall)
			db.Where("big_id=? AND name=?",newPlaceBig.ID,subvalue).First(&newPlaceSmall)
			if newPlaceSmall.ID == 0 {
				var indexx uint = 0
				if placeSmall.ID != 0 {
					indexx = placeSmall.Indexx + 1
				}
				db.Create(&dbModel.PlaceSmall{
					Indexx:   indexx,
					Name:     subvalue,
					BigId:    newPlaceBig.ID,
					BigName:  newPlaceBig.Name,
					CampusId: newPlaceBig.CampusId,
				})
			}
		}
	}
	for key,value := range areas1 {
		var placeBig dbModel.PlaceBig
		var newPlaceBig dbModel.PlaceBig
		db.Where("campus_id=?","1").Order("indexx DESC").Limit(1).Find(&placeBig)
		db.Where("campus_id=? AND name=?","1",value).First(&newPlaceBig)
		var index uint = 0
		if placeBig.ID != 0 {
			index = placeBig.Indexx + 1
		}
		if newPlaceBig.ID == 0 {
			newPlaceBig = dbModel.PlaceBig {
				Indexx:   index,
				Name:     value,
				CampusId: "1",
			}
			db.Create(&newPlaceBig)
		}
		for _, subvalue := range subareasList1[key] {
			var placeSmall dbModel.PlaceSmall
			var newPlaceSmall dbModel.PlaceSmall
			db.Where("big_id=?",newPlaceBig.ID).Order("indexx DESC").Limit(1).Find(&placeSmall)
			db.Where("big_id=? AND name=?",newPlaceBig.ID,subvalue).First(&newPlaceSmall)
			var indexx uint = 0
			if placeSmall.ID != 0 {
				indexx = placeSmall.Indexx + 1
			}
			if newPlaceSmall.ID == 0 {
				db.Create(&dbModel.PlaceSmall{
					Indexx:   indexx,
					Name:     subvalue,
					BigId:    newPlaceBig.ID,
					BigName:  newPlaceBig.Name,
					CampusId: newPlaceBig.CampusId,
				})
			}
		}
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code" : 200,
		"msg"  : "地点信息成功恢复为默认值。",
	})
}