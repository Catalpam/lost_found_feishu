package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"lost_found/dbModel"
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
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName,args)
	if err != nil {
		fmt.Print("fail to connect database, err: " + err.Error())
		panic("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&dbModel.User{})
	db.AutoMigrate(&dbModel.Type{})
	db.AutoMigrate(&dbModel.ItemClass{})
	db.AutoMigrate(&dbModel.Found{})
	db.AutoMigrate(&dbModel.Lost{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}