package dbModel

type Privilege  struct {
	ID       uint `gorm:"primary_key"`
	OpenId string `gorm:"type:varchar(40);not null;unique"`
	// 0 无权限，1-查看权限，2-管理员权限，3-超级管理员权限
	Permission string `gorm:"type:char(1);not null;"`
	PlaceSmallId uint
}
