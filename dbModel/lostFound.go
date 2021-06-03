package dbModel

import (
	"time"
)

type Found struct {
	//gorm Model
	ID uint `gorm:"primary_key"`
	Validity bool `gorm:"default:true;not null"`
	// 发现时间
	CreatedAt time.Time
	// Match时间
	UpdatedAt time.Time
	// 匹配的编号
	MatchId uint
	//发现日期&时间&时间段
	FoundDate        string `gorm:"type:varchar(20);not null;"`
	FoundTime        string `gorm:"type:varchar(20);not null;"`
	FoundTimeSession string `gorm:"type:varchar(20);not null"`
	//Founder信息
	FoundOpenId string `gorm:"type:varchar(50);not null;"`
	// 物品类型
	ItemType string `gorm:"type:varchar(20);not null;"`
	SubType  string `gorm:"type:varchar(20);not null;"`

	// 发现地点
	Campus   string `gorm:"type:varchar(20);not null"`
	Place    string `gorm:"type:varchar(20);not null"`
	SubPlace string `gorm:"type:varchar(20);not null"`

	// 物品信息
	ItemInfo string `gorm:"type:varchar(500);"`
	// 图片
	Image string `gorm:"type:varchar(200);"`
	ImageHome string `gorm:"type:varchar(200);"`
	ImageKey string `gorm:"type:varchar(200);"`
	// 发现的详细地点
	PlaceDetail string `gorm:"type:varchar(500);"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlace string `gorm:"type:char(1);not null"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlaceDetail string `gorm:"type:char(200);"`
	// 失主身份信息(填姓名）
	LosterInfo string `gorm:"type:varchar(500);"`
	// 补充信息
	AdditionalInfo string `gorm:"type:varchar(500)"`
	//失主留言
	LosterComment string `gorm:"type:varchar(500)"`
}

type Lost struct {
	//gorm Model
	ID        uint `gorm:"primary_key"`
	Validity bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	// Match时间
	UpdatedAt time.Time

	//Loster的OpenId
	LosterOpenId string `gorm:"type:char(50);not null"`
	//匹配的Found ID
	MatchId uint

	// 物品类型 为Thing表的NameID
	TypeSubName string `gorm:"type:varchar(20);not null;"`

	// 丢失地点
	LostPlace1 string `gorm:"type:char(100);not null"`
	LostPlace2 string `gorm:"type:char(100);"`
	LostPlace3 string `gorm:"type:char(100);"`

	LostDate string `gorm:"type:char(15);"`
	// 大致丢失的时间段
	//morning    上午（6：00-11：00）
	//noon       中午（11：00-2：00）
	//afternoon  下午（2：00-19：00）
	//evening    晚上（19：00-22：00）
	//night      夜间（00：00-6：00，22：00-24：00）
	LostTimeSession string `gorm:"type:varchar(20);"`
}

type LostPlace struct {
	ID       uint `gorm:"primary_key"`
	CampusName string `gorm:"type:varchar(100);not null"`
	PlaceName string `gorm:"type:varchar(100);not null"`
	SubPlaceName string `gorm:"type:varchar(100);not null"`
}

type LostPlaceList []LostPlace