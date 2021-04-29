package dbModel

import "github.com/jinzhu/gorm"


type Thing struct {
	gorm.Model
	Name string `gorm:"type:varchar(200);not null"`
	NameId string `gorm:"type:varchar(200);nor null;unique"`
	Type string `gorm:"type:varchar(200);not null"`
	TypeId string `gorm:"type:varchar(200);not null"`
}


