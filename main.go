package main

import (
	"fmt"
	"lost_found/common"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")

func main() {

	db := common.InitDB()
	defer db.Close()

	//testSendMessage()
	//testSendCardMessage()
	//testDownloadFile()
	//feishu.SendCardMessage()
	fmt.Print("Success DB Init")
}
