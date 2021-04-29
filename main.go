package main

import (
	"lost_found/common"
	"lost_found/controller"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")

func main() {

	db := common.InitDB()
	defer db.Close()

	//feishu.SendCardMessage()
	//fmt.Print("Success DB Init")
	//feishu.UpdateStudentList()
	controller.ThingDbDefaultInit()
}
