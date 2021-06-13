package dbModel

type PlaceSmall struct {
	ID       uint `gorm:"primary_key"`
	Indexx   uint
	Name     string `gorm:"type:varchar(100);not null"`
	BigId    uint
	BigName  string `gorm:"type:varchar(100);not null"`
	CampusId string `gorm:"type:varchar(10);not null"`
}

type PlaceBig struct {
	ID       uint `gorm:"primary_key"`
	Indexx   uint
	Name     string `gorm:"type:varchar(100);not null"`
	CampusId string `gorm:"type:varchar(10);not null"`
}

func CampusId2Str(id string) string {
	ret := ""
	if id == "0" {
		ret = "清水河校区"
	} else if id == "1" {
		ret = "沙河校区"
	}
	return ret
}
