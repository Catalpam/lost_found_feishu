package dbModel

import (
	"time"
)

type Found struct {
	//gorm Model
	ID        uint `gorm:"primary_key"`
	// 发现时间
	CreatedAt time.Time
	UpdatedAt time.Time

	// 物品类型 为Thing表的NameID
	TypeName string `gorm:"type:varchar(20);not null;"`
	// 物品大类型 为Thing表的NameID
	ClassName string `gorm:"type:varchar(20);not null;"`
	// 物品信息
	ItemInfo string `gorm:"type:varchar(500);"`
	// 图片 直接存序列化好的数组[]string类型，以varchar保存
	Image string `gorm:"type:varchar(200);"`
	// 发现地点 为Location表的NameID
	Place string `gorm:"type:varchar(20);not null"`
	// 发现的详细地点
	PlaceDetail string `gorm:"type:varchar(500);"`
	// 失主身份信息(填姓名）
	LosterInfo string `gorm:"type:varchar(500);"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlace string `gorm:"type:char(1);not null"`
	// 补充信息
	AdditionalInfo string `gorm:"type:varchar(500)"`
}

type Lost struct {
	//gorm Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 物品类型 为Thing表的NameID
	TypeId string `gorm:"type:varchar(20);not null;"`
	// 物品信息
	ItemInfo string `gorm:"type:varchar(500);"`
	// 图片 直接保存为Image Key的Json
	Image string `gorm:"type:varchar(200);nor null;"`
	// 丢失地点 为Location表的NameID
	Place1 string `gorm:"type:char(12);not null"`
	Place2 string `gorm:"type:char(12);not null"`
	Place3 string `gorm:"type:char(12);not null"`
	// 丢失时间段 可以填
	//Morning    上午（6：00-11：00）
	//Noon       中午（11：00-2：00）
	//Afternoon  下午（2：00-19：00）
	//Evening    晚上（19：00-22：00）
	//Night      夜间（00：00-6：00，22：00-24：00）
	TimeSession string `gorm:"type:varchar(20);not null"`
}