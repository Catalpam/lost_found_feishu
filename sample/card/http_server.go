package main

import (
	"fmt"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	card2 "lost_found/handler/card"
	"lost_found/handler/card/http/native"
	model2 "lost_found/handler/card/model"
	"lost_found/sample/configs"
	"net/http"
)

func main() {

	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")
	var conf = configs.TestConfig(constants.DomainFeiShu)

	card2.SetHandler(conf, func(ctx *core.Context, card *model2.Card) (interface{}, error) {
		fmt.Println(ctx.GetRequestID())
		fmt.Println(tools.Prettify(card))
		return "{\"config\":{\"wide_screen_mode\":true},\"i18n_elements\":{\"zh_cn\":[{\"tag\":\"div\",\"text\":{\"tag\":\"lark_md\",\"content\":\"[飞书golang](https://www.feishu.cn)整合即时沟通、日历、音视频会议、云文档、云盘、工作台等功能于一体，成就组织和个人，更高效、更愉悦。\"}}]}}", nil
	})

	native.Register("/webhook/card", conf)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		panic(err)
	}

}
