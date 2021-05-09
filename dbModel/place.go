package dbModel

type Campus struct {
	ID uint `gorm:"primary_key"`
	CampusId string `gorm:"type:varchar(10);not null"`
	Name string `gorm:"type:varchar(200);not null"`
}

type Place struct {
	ID       uint `gorm:"primary_key"`
	CampusId string `gorm:"type:varchar(10);not null"`
	PlaceId string `gorm:"type:varchar(10);not null"`
	Name     string `gorm:"type:varchar(200);not null"`
	Subareas string `gorm:"type:varchar(300);not null"`
}