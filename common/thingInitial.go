package common

import (
	"fmt"
	"lost_found/dbModel"
)

func ThingDbDefaultInit() {
	DB := GetDB()
	for _, value := range DefaultThings {
		//
		var thing dbModel.Type
		type_id := value.TypeId
		DB.Where("type_id = ?", type_id).First(&thing)
		if thing.ID != 0 {
			fmt.Print(thing.Type + "  " + thing.TypeId + " already exits\n")
			continue
		}
		newThing := dbModel.Type{
			Type:    value.Type,
			TypeId:  value.TypeId,
			Class:   value.Class,
			ClassId: value.ClassId,
		}
		DB.Create(&newThing)
	}
}

var DefaultThings []dbModel.Type = []dbModel.Type{
	dbModel.Type{
		Type:    "手机",
		TypeId:  "1001",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "平板电脑",
		TypeId:  "1002",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "电纸书",
		TypeId:  "1003",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "笔记本电脑",
		TypeId:  "1004",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "耳机",
		TypeId:  "1005",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "充电器",
		TypeId:  "1006",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "充电宝",
		TypeId:  "1007",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "数据线",
		TypeId:  "1008",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "手表",
		TypeId:  "1009",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "U盘",
		TypeId:  "1010",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "Apple Pencil",
		TypeId:  "1011",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "耳机",
		TypeId:  "1012",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "鼠标",
		TypeId:  "1013",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "键盘",
		TypeId:  "1014",
		Class:   "电子设备",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "其他电子设备",
		TypeId:  "1000",
		Class:   "电子设备",
		ClassId: "1",
	},

	dbModel.Type{
		Type:    "一卡通",
		TypeId:  "2001",
		Class:   "证件或证书",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "身份证",
		TypeId:  "2002",
		Class:   "证件或证书",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "学生证",
		TypeId:  "2003",
		Class:   "证件或证书",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "健身卡",
		TypeId:  "2004",
		Class:   "证件或证书",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "各种证书",
		TypeId:  "2005",
		Class:   "证件或证书",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "其他证件或证书",
		TypeId:  "2000",
		Class:   "证件或证书",
		ClassId: "2",
	},
}
