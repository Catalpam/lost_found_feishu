package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"lost_found/dbModel"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "lost_found"
	username := "root"
	password := "zhzhzhzh"
	charset := "utf8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(driverName,args)
	if err != nil {
		fmt.Print("fail to connect database, err: " + err.Error())
		panic("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&dbModel.User{})
	db.AutoMigrate(&dbModel.TypeSmall{})
	db.AutoMigrate(&dbModel.TypeBig{})
	db.AutoMigrate(&dbModel.PlaceSmall{})
	db.AutoMigrate(&dbModel.PlaceBig{})
	db.AutoMigrate(&dbModel.Found{})
	db.AutoMigrate(&dbModel.Lost{})
	db.AutoMigrate(&dbModel.Match{})
	db.AutoMigrate(&dbModel.Privilege{})
	println("数据库加载完成")

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}


func TypeId2Name(id uint) string  {
	db := GetDB()
	var typeSmall dbModel.TypeSmall
	db.Where("id=?",id).First(&typeSmall)
	println(typeSmall.ID)
	println(typeSmall.Name)

	return typeSmall.BigName+" "+typeSmall.Name
}

func PlaceId2Name(id uint) string  {
	db := GetDB()
	var placeSmall dbModel.PlaceSmall
	println(id)
	db.Where("id=?",id).Find(&placeSmall)
	println(placeSmall.ID)
	println(placeSmall.Name)
	if placeSmall.ID == 0 {
		return ""
	} else {
		return placeSmall.BigName+" "+placeSmall.Name
	}
}