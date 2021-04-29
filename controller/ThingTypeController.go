package controller

import (
	"fmt"
	"lost_found/common"
	"lost_found/dbModel"
)

func ThingDbDefaultInit() {
	DB := common.GetDB()
	DefaultThings := []dbModel.Thing {
		dbModel.Thing{
			Name : "手机",
			NameId : "1001",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "平板电脑",
			NameId : "1002",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "电纸书",
			NameId : "1003",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "笔记本电脑",
			NameId : "1004",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "耳机",
			NameId : "1005",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "充电器",
			NameId : "1006",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "充电宝",
			NameId : "1007",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "数据线",
			NameId : "1008",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "手表",
			NameId : "1009",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "U盘",
			NameId : "1010",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "Apple Pencil",
			NameId : "1011",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "耳机",
			NameId : "1012",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "鼠标",
			NameId : "1013",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "键盘",
			NameId : "1014",
			Type : "电子设备",
			TypeId :"1",
		},
		dbModel.Thing{
			Name : "其他电子设备",
			NameId : "1000",
			Type : "电子设备",
			TypeId :"1",
		},


		dbModel.Thing{
			Name : "一卡通",
			NameId : "2001",
			Type: "证件或证书",
			TypeId :"2",
		},
		dbModel.Thing{
			Name : "身份证",
			NameId : "2002",
			Type : "证件或证书",
			TypeId :"2",
		},
		dbModel.Thing{
			Name : "学生证",
			NameId : "2003",
			Type : "证件或证书",
			TypeId :"2",
		},
		dbModel.Thing{
			Name : "健身卡",
			NameId : "2004",
			Type : "证件或证书",
			TypeId :"2",
		},
		dbModel.Thing{
			Name : "各种证书",
			NameId : "2005",
			Type : "证件或证书",
			TypeId :"2",
		},
		dbModel.Thing{
			Name : "其他证件或证书",
			NameId : "2000",
			Type : "证件或证书",
			TypeId :"2",
		},
	}

	for _,value := range DefaultThings {
		//
		var thing dbModel.Thing
		name_id := value.NameId
		DB.Where("name_id = ?", name_id).First(&thing)
		if thing.ID != 0{
			fmt.Print(thing.Name+"  "+thing.NameId+" already exits\n")
			continue
		}
		//
		newThing := dbModel.Thing{
			Name: value.Name,
			NameId: value.NameId,
			Type: value.Type,
			TypeId: value.TypeId,
		}
		DB.Create(&newThing)
	}
}
