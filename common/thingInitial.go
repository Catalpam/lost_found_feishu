package common

import (
	"encoding/json"
	"lost_found/dbModel"
)
func ItemTypeInitial() {
	db := GetDB()
	var newItemType dbModel.ItemType
	db.Where("name=? ","电子设备").First(&newItemType)
	if newItemType.ID == 0 {
		subareas := []string{"手机","平板电脑","电纸书","笔记本电脑","耳机","充电器","充电宝","数据线",
			"智能手表","U盘","Apple Pencil","鼠标","键盘","其他电子设备"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newItemType = dbModel.ItemType{
			Name:     "电子设备",
			TypeId:   "0",
			Subtypes: string(subareasJson),
		}
		db.Create(&newItemType)
	}
	newItemType.ID = 0
	db.Where("name=? ","证书或证件").First(&newItemType)
	if newItemType.ID == 0 {
		subareas := []string{"一卡通","身份证","学生证","健身卡","各种证书","其他证件或证书"}
		subareasJson,err := json.Marshal(subareas)
		if err != nil {
			println("--------Marshal error:------------")
			println(err)
		}
		newItemType = dbModel.ItemType{
			Name:     "证件或证书",
			TypeId:   "1",
			Subtypes: string(subareasJson),
		}
		db.Create(&newItemType)
	}
}
