package main

import (
	"context"
	"fmt"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/sample/configs"
	translation "lost_found/service/translation/v1"
)

// for redis store and logrus
// configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.TestConfig("https://open.feishu.cn")
var translationService = translation.NewService(configs.TestConfig(constants.DomainFeiShu))

func main() {
	testTextDetect()
}

func testTextDetect() {
	coreCtx := core.WrapContext(context.Background())
	reqCall := translationService.Texts.Translate(coreCtx, &translation.TextTranslateReqBody{
		SourceLanguage: "zh",
		Text:           "测试",
		TargetLanguage: "en",
		Glossary: []*translation.Term{
			{
				From: "测",
				To:   "test",
			},
		},
	})
	result, err := reqCall.Do()
	fmt.Printf("request_id:%s\n", coreCtx.GetRequestID())
	fmt.Printf("http status code:%d", coreCtx.GetHTTPStatusCode())
	if err != nil {
		e := err.(*response.Error)
		fmt.Printf(tools.Prettify(e))
		return
	}
	fmt.Printf("reault:%s", tools.Prettify(result))
}
