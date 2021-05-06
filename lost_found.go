package main

import (
	"github.com/gin-gonic/gin"
	"lost_found/common"
	"lost_found/routes"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")

func main() {

	db := common.InitDB()
	defer db.Close()
	//fmt.Print("Success DB Init")
	//controller.UpdateStudentList()
	//controller.ThingDbDefaultInit()
	//controller.GetAccessToken()
	r := gin.Default()
	r = routes.CollectRoute(r)
	print("hello")
	panic(r.Run(":1111"))

}
