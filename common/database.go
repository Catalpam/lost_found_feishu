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
	database := "feishu_lost_found"
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
	db.AutoMigrate(&dbModel.Campus{})
	db.AutoMigrate(&dbModel.Place{})
	db.AutoMigrate(&dbModel.ItemType{})
	db.AutoMigrate(&dbModel.Found{})
	db.AutoMigrate(&dbModel.Lost{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}