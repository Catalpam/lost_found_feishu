package main

import (
	"context"
	"fmt"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/sample/configs"
	optical_char_recognition "lost_found/service/optical_char_recognition/v1"
)

// for redis store and logrus
// configs.TestConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// configs.TestConfig("https://open.feishu.cn")
var opticalCharRecognitionService = optical_char_recognition.NewService(configs.TestConfig(constants.DomainFeiShu))

func main() {
	testImageBasicRecognize()
}

func testImageBasicRecognize() {
	coreCtx := core.WrapContext(context.Background())
	reqCall := opticalCharRecognitionService.Images.BasicRecognize(coreCtx, &optical_char_recognition.ImageBasicRecognizeReqBody{
		Image: "base64 image",
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
