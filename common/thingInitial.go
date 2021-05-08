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
			ClassId: value.ClassId,
		}
		DB.Create(&newThing)
	}
	for _, value := range DefaultClasses {
		//
		var itemClass dbModel.ItemClass
		class_id := value.ClassId
		DB.Where("class_id = ?", class_id).First(&itemClass)
		if itemClass.ID != 0 {
			fmt.Print(itemClass.ClassName + "  " + itemClass.ClassId + " already exits\n")
			continue
		}
		newItemClass := dbModel.ItemClass{
			ClassName:    value.ClassName,
			ClassId:      value.ClassId,
		}
		DB.Create(&newItemClass)
	}

}

var DefaultThings []dbModel.Type = []dbModel.Type{
	dbModel.Type{
		Type:    "手机",
		TypeId:  "1001",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "平板电脑",
		TypeId:  "1002",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "电纸书",
		TypeId:  "1003",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "笔记本电脑",
		TypeId:  "1004",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "耳机",
		TypeId:  "1005",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "充电器",
		TypeId:  "1006",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "充电宝",
		TypeId:  "1007",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "数据线",
		TypeId:  "1008",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "手表",
		TypeId:  "1009",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "U盘",
		TypeId:  "1010",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "Apple Pencil",
		TypeId:  "1011",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "耳机",
		TypeId:  "1012",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "鼠标",
		TypeId:  "1013",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "键盘",
		TypeId:  "1014",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "其他电子设备",
		TypeId:  "1000",
		ClassId: "1",
	},

	dbModel.Type{
		Type:    "一卡通",
		TypeId:  "2001",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "身份证",
		TypeId:  "2002",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "学生证",
		TypeId:  "2003",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "健身卡",
		TypeId:  "2004",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "各种证书",
		TypeId:  "2005",
		ClassId: "2",
	},
	dbModel.Type{
		Type:    "其他证件或证书",
		TypeId:  "2000",
		ClassId: "2",
	},
}

var DefaultClasses = []dbModel.ItemClass{
	dbModel.ItemClass{
		ClassName: "电子设备",
		ClassId: "1",
	},
	dbModel.ItemClass{
		ClassName: "证书或证件",
		ClassId: "2",
	},
}
