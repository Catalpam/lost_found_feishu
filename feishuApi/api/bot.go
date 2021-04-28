package feishu

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/constants"
	"github.com/larksuite/oapi-sdk-go/core/tools"
	"github.com/larksuite/oapi-sdk-go/sample/configs"
	bot "github.com/larksuite/oapi-sdk-go/service/bot/v3"
)

// for redis store and logrus
// configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.FeishuConfig("https://open.feishu.cn")
var botService = bot.NewService(configs.TestConfig(constants.DomainFeiShu))

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
