package main

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/controller"
	"lost_found/routes"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")

func main() {

	db := common.InitDB()
	defer db.Close()

	//feishu.SendCardMessage()
	//fmt.Print("Success DB Init")
	controller.UpdateStudentList()
	controller.ThingDbDefaultInit()

	r := gin.Default()
	r = routes.CollectRoute(r)
	panic(r.Run(":1111"))

}
