package main

import (
	feishu "lost_found/feishuApi/api"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")

func main() {
	//testSendMessage()
	//testSendCardMessage()
	//testDownloadFile()
	feishu.SendCardMessage()
}
