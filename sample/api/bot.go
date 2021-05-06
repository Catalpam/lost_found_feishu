package main

import (
	"context"
	"fmt"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/sample/configs"
	bot "lost_found/service/bot/v3"
)

// for redis store and logrus
// configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.TestConfig("https://open.feishu.cn")
var botService = bot.NewService(configs.TestConfig(constants.DomainFeiShu))

func main() {
	testBotGet()
}

func testBotGet() {
	coreCtx := core.WrapContext(context.Background())
	reqCall := botService.Bots.Get(coreCtx)
	result, err := reqCall.Do()
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(result.Bot))
}
