package dbModel

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	StudentId string `gorm:"type:varchar(40);not null;unique"`
	OpenId string `gorm:"type:varchar(40);not null;unique"`
	//UserId string `gorm:"type:varchar(20);not null"`
	Mobile string `gorm:"type:char(14);not null"`
	//UnionId string `gorm:"type:varchar(20);not null"`
	DepartmentId string `gorm:"type:varchar(80);not null"`
	Avatar string `gorm:"type:varchar(500);not null"`
}