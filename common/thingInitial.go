package common

import (
	"encoding/json"
	"lost_found/dbModel"
)

func ItemTypeInitial() {
	db := GetDB()
	{
		var newItemType dbModel.ItemType
		db.Where("name=?", "电子设备").First(&newItemType)
		if newItemType.ID == 0 {
			subareas := []string{"手机", "平板电脑", "电纸书", "笔记本电脑", "耳机", "充电器", "充电宝", "数据线",
				"智能手表", "U盘", "Apple Pencil", "鼠标", "键盘", "其他电子设备"}
			subareasJson, err := json.Marshal(subareas)
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
	}
	{
		var newItemType dbModel.ItemType
		db.Where("name=?", "证件或证书").First(&newItemType)
		if newItemType.ID == 0 {
			subareas := []string{"一卡通", "身份证", "学生证", "健身卡","考试通过证书","比赛奖状","其他证件或证书"}
			subareasJson, err := json.Marshal(subareas)
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
	{
		var newItemType dbModel.ItemType
		db.Where("name=?", "个人物品").First(&newItemType)
		if newItemType.ID == 0 {
			subareas := []string{"衣物（含帽子手套围巾等）", "包", "眼镜", "钱包", "钥匙", "水杯", "化妆品", "首饰", "笔袋", "未领走快递", "雨伞", "运动器械相关","其它个人物品"}
			subareasJson, err := json.Marshal(subareas)
			if err != nil {
				println("--------Marshal error:------------")
				println(err)
			}
			newItemType = dbModel.ItemType{
				Name:     "个人物品",
				TypeId:   "2",
				Subtypes: string(subareasJson),
			}
			db.Create(&newItemType)
		}
	}
	{
		var newItemType dbModel.ItemType
		db.Where("name=?", "学习用品").First(&newItemType)
		if newItemType.ID == 0 {
			subareas := []string{"笔袋", "书籍", "纸质笔记本"}
			subareasJson, err := json.Marshal(subareas)
			if err != nil {
				println("--------Marshal error:------------")
				println(err)
			}
			newItemType = dbModel.ItemType{
				Name:     "学习用品",
				TypeId:   "3",
				Subtypes: string(subareasJson),
			}
			db.Create(&newItemType)
		}
	}
}
