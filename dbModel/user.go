package dbModel

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	StudentId string `gorm:"type:varchar(40);not null"`
	OpenId string `gorm:"type:varchar(40);not null"`
	//UserId string `gorm:"type:varchar(20);not null"`
	Mobile string `gorm:"type:char(11);not null"`
	//UnionId string `gorm:"type:varchar(20);not null"`
	DepartmentId string `gorm:"type:varchar(30);not null"`
	Avatar string `gorm:"type:varchar(40);not null"`
}