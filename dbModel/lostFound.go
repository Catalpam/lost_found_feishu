package dbModel

import (
	"time"
)

type Found struct {
	//gorm Model
	ID 				uint `gorm:"primary_key"`
	Validity 		bool `gorm:"default:true;not null"`
	//Founder信息
	OpenId 			string `gorm:"type:varchar(50);not null;"`
	// 匹配的编号
	MatchId 		uint
	//发现日期&时间&时间段
	Date  	    	string `gorm:"type:varchar(20);not null;"`
	Time        	string `gorm:"type:varchar(20);not null;"`
	TimeSession 	string `gorm:"type:varchar(20);not null"`
	// 物品类型
	TypeBigId 		uint
	TypeSmallId  	uint
	// 发现地点
	PlaceBigId  	uint
	PlaceSmallId 	uint

	// 物品信息
	ItemInfo 			string `gorm:"type:varchar(500);"`
	// 图片
	Image 				string `gorm:"type:varchar(200);"`
	ImageKey 			string `gorm:"type:varchar(200);"`
	// 发现的详细地点
	PlaceDetail 		string `gorm:"type:varchar(500);"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlace 		string `gorm:"type:char(1);not null"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlaceDetail 	string `gorm:"type:char(200);"`
	// 失主身份信息(填姓名）
	LosterInfo 			string `gorm:"type:varchar(500);"`
	// 补充信息
	AdditionalInfo 		string `gorm:"type:varchar(500)"`
}

type Lost struct {
	//gorm Model
	ID        		uint `gorm:"primary_key"`
	Validity 		bool `gorm:"default:true;not null"`
	IsFoundBySelf	bool
	CreatedAt 		time.Time
	UpdatedAt 		time.Time

	//Loster的OpenId
	OpenId 			string `gorm:"type:char(50);not null"`
	Name			string `gorm:"type:varchar(50);not null;"`
	//匹配的Found ID
	MatchId 		uint

	// 物品类型
	TypeBigId 		uint
	TypeSmallId  	uint
	// 发现地点
	PlaceSmallId1 	uint
	PlaceSmallId2 	uint
	PlaceSmallId3 	uint
	// 大致丢失的时间
	Date 			string `gorm:"type:char(15);"`
	TimeSession 	string `gorm:"type:varchar(20);"`
}

type Match struct {
	//gorm Model
	ID uint `gorm:"primary_key"`
	CreatedAt 	time.Time

	//发现日期&时间&时间段
	FoundDate   string `gorm:"type:varchar(20);not null;"`
	Time        string `gorm:"type:varchar(20);not null;"`
	TimeSession string `gorm:"type:varchar(20);not null"`
	//Founder信息
	LosterOpenId 			string `gorm:"type:varchar(50);not null;"`
	FoundOpenId 			string `gorm:"type:varchar(50);not null;"`

	// 物品类型
	TypeBigId 		uint
	TypeSmallId  	uint
	TypeName		string `gorm:"type:varchar(100);not null;"`
	// 发现地点
	Campus			string `gorm:"type:varchar(50);not null;"`
	PlaceBigId  	uint
	PlaceSmallId 	uint
	PlaceName		string `gorm:"type:varchar(100);not null;"`

	// 物品信息
	ItemInfo 			string `gorm:"type:varchar(500);"`
	// 图片
	Image 				string `gorm:"type:varchar(200);"`
	ImageKey 			string `gorm:"type:varchar(200);"`
	// 发现的详细地点
	PlaceDetail 		string `gorm:"type:varchar(500);"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlace 		string `gorm:"type:char(1);not null"`
	// 当前位置：0-留在原地 1-自己带走 2-教导失物招领处
	CurrentPlaceDetail 	string `gorm:"type:char(200);"`
	// 失主身份信息(填姓名）
	LosterInfo 			string `gorm:"type:varchar(500);"`
	// 补充信息
	AdditionalInfo 		string `gorm:"type:varchar(500)"`
	//失主感谢
	LosterComment 		string `gorm:"type:varchar(500)"`
}
//morning    上午（6：00-11：00）
//noon       中午（11：00-2：00）
//afternoon  下午（2：00-19：00）
//evening    晚上（19：00-22：00）
//night      夜间（00：00-6：00，22：00-24：00）
