package dbModel

type TypeSmall struct {
	ID       	uint `gorm:"primary_key"`
	Indexx    	uint
	Name     	string `gorm:"type:varchar(100);not null"`
	BigName		string `gorm:"type:varchar(100);not null"`
	BigId		uint
}

type TypeBig struct {
	ID     uint `gorm:"primary_key"`
	Indexx uint
	Name   string `gorm:"type:varchar(100);not null"`
}

