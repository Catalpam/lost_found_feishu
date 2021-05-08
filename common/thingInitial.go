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
		TypeId:  "01",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "平板电脑",
		TypeId:  "02",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "电纸书",
		TypeId:  "03",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "笔记本电脑",
		TypeId:  "04",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "耳机",
		TypeId:  "05",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "充电器",
		TypeId:  "06",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "充电宝",
		TypeId:  "07",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "数据线",
		TypeId:  "08",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "手表",
		TypeId:  "09",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "U盘",
		TypeId:  "010",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "Apple Pencil",
		TypeId:  "011",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "耳机",
		TypeId:  "012",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "鼠标",
		TypeId:  "013",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "键盘",
		TypeId:  "014",
		ClassId: "0",
	},
	dbModel.Type{
		Type:    "其他电子设备",
		TypeId:  "00",
		ClassId: "0",
	},


	dbModel.Type{
		Type:    "一卡通",
		TypeId:  "11",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "身份证",
		TypeId:  "12",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "学生证",
		TypeId:  "13",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "健身卡",
		TypeId:  "14",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "各种证书",
		TypeId:  "15",
		ClassId: "1",
	},
	dbModel.Type{
		Type:    "其他证件或证书",
		TypeId:  "10",
		ClassId: "1",
	},
}

var DefaultClasses = []dbModel.ItemClass{
	dbModel.ItemClass{
		ClassName: "电子设备",
		ClassId: "0",
	},
	dbModel.ItemClass{
		ClassName: "证书或证件",
		ClassId: "1",
	},
}
