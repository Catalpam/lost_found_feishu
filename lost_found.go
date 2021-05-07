package main

import (
	"lost_found/common"
	"lost_found/handler"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")

func main() {
	db := common.InitDB()
	defer db.Close()
	handler.SendUser()
	//r := gin.Default()
	//r = routes.CollectRoute(r)
	//print("hello")
	//
	//panic(r.Run(":1111"))
}