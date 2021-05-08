package dbModel

import "github.com/jinzhu/gorm"


type Type struct {
	gorm.Model
	//物品小分类：名称
	Type   string `gorm:"type:varchar(200);not null"`
	TypeId string `gorm:"type:varchar(200);nor null;unique"`
	//物品大分类：类型
	ClassId string `gorm:"type:varchar(200);not null"`
}

type ItemClass struct {
	gorm.Model
	ClassName   string `gorm:"type:varchar(200);not null"`
	ClassId string `gorm:"type:varchar(200);not null"`
}

//type ItemClass struct {
//	Name string
//	Types []string
//}
//
//var ItemClasses []ItemClass