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
	db.Where("name=? AND campus_id=?","食堂","1").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"阳光餐厅","万友餐厅","风华餐厅","桂苑餐厅","西北美食坊"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "1",
			PlaceId : "2",
			Name:     "食堂",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","运动场所","1").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"体育馆","羽毛球馆","健身房","篮球场","操场","乒乓球场","足球场","保卫处","游泳馆","网球场"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "1",
			PlaceId : "3",
			Name:     "运动场所",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","其它","1").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"图书馆","主楼","学术交流中心","校医院住院部","校医院","沙河停车场"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "1",
			PlaceId : "4",
			Name:     "其它",
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
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","学知苑","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"学知苑1栋","学知苑2栋","学知苑3栋","学知苑4栋","学知苑5栋","学知苑6栋","学知苑7栋","学知苑8栋","学知苑9栋","学知苑10栋","学知苑11栋","学知苑12栋","学知苑13栋","学知苑14栋","学知苑15栋","学知苑16栋","学知苑17栋","学知苑18栋","学知苑19栋","学知苑20栋","学知苑21栋","学知苑22栋","学知苑23栋","学知苑24栋","学知苑25栋","学知苑26栋","学知苑27栋","学知苑28栋"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "2",
			Name:     "学知苑",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	db.Where("name=? AND campus_id=?","博翰苑","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"博翰苑1栋","博翰苑2栋","博翰苑3栋","博翰苑4栋","博翰苑5栋","博翰苑6栋","博翰苑7栋","博翰苑8栋","博翰苑9栋"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "3",
			Name:     "博翰苑",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","留学生宿舍区","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"留学生2栋","留学生3栋","留学生4栋"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "4",
			Name:     "留学生宿舍区",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","餐厅","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"清真餐厅","学生二食堂","春晖苑（三食堂）"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "5",
			Name:     "餐厅",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
	newPlace.ID = 0
	db.Where("name=? AND campus_id=?","其他","0").First(&newPlace)
	if newPlace.ID == 0 {
		subareas := []string{"体育馆","游泳馆","体育场","众创空间","主楼A1区","主楼A2区","主楼B1区","主楼B2区","主楼B3区","主楼C2区","主楼C1区","图书馆","校医院","学生活动中心","成电会堂","商业街","综合训练馆","游泳中心"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newPlace = dbModel.Place{
			CampusId: "0",
			PlaceId : "6",
			Name:     "其他",
			Subareas: string(subareasJson),
		}
		db.Create(&newPlace)
	}
}