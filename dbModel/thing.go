package dbModel

import (
	"github.com/jinzhu/gorm"
)

type Thing struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Type string `gorm:"type:varchar(20);not null"`
}

