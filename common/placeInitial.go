package common

import (
	"encoding/json"
	"lost_found/dbModel"
)

func PlaceInitial() {
	db := GetDB()
	var searchCampus dbModel.Campus
	db.Where("campus_id=?","0").First(&searchCampus)
	if searchCampus.ID == 0 {
		var newCampus = dbModel.Campus{
			CampusId : "0",
			Name : "清水河校区",
		}
		db.Create(&newCampus)
	}
	searchCampus.ID = 0
	db.Where("campus_id=?","1").First(&searchCampus)
	if searchCampus.ID == 0 {
		var newCampus = dbModel.Campus{
			CampusId : "1",
			Name : "沙河校区",
		}
		db.Create(&newCampus)
	}

	var newPlace dbModel.Place
	db.Where("name=? AND campus_id=?","宿舍区","1").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"校内16栋","校内17栋","校内18栋","校内12栋","校内15栋","校内14栋","校内13栋","欣苑"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "1",
			PlaceId : "0",
			Name:     "宿舍区",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","教学楼","1").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"第二教学楼","逸夫楼","微固学院","光电学院","医学院","第三教学楼","第四教学楼","信软实验楼","通信大楼","211大楼"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "1",
			PlaceId : "1",
			Name:     "教学楼",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","教学楼","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"研究院大楼","四号科研楼A区","四号科研楼B区","四号科研楼C区","五号科研楼","综合楼","磁共振脑成像中心","科研楼","基础实验楼","品学楼A区","品学楼B区","品学楼C区","立人楼"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "0",
			Name:     "教学楼",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","硕丰苑","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"硕丰苑1栋","硕丰苑2栋","硕丰苑3栋","硕丰苑4栋","硕丰苑5栋","硕丰苑6栋","硕丰苑7栋","硕丰苑8栋","硕丰苑9栋","硕丰苑10栋","硕丰苑11栋","硕丰苑12栋","硕丰苑13栋","硕丰苑14栋","硕丰苑15栋","硕丰苑16栋","硕丰苑17栋","硕丰苑18栋","硕丰苑19栋","硕丰苑20栋","硕丰苑21栋","硕丰苑22栋","硕丰苑23栋","硕丰苑24栋","硕丰苑25栋","硕丰苑26栋","硕丰苑27栋","硕丰苑28栋"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "1",
			Name:     "硕丰苑",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
}
