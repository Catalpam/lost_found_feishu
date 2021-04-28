package feishu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/constants"
	"github.com/larksuite/oapi-sdk-go/core/tools"
	"github.com/larksuite/oapi-sdk-go/sample/configs"
	"io/ioutil"
	"os"
)

// for redis store and logrus
// var conf = configs.ConfigWithLogrusAndRedisStore(constants.DomainFeiShu)
// var conf = configs.FeishuConfig("https://open.feishu.cn")
var conf = configs.TestConfig(constants.DomainFeiShu)


// send card message
func SendCardMessage() {
	coreCtx := core.WrapContext(context.Background())
	cardContent := "{\"config\":{\"wide_screen_mode\":true},\"i18n_elements\":{\"zh_cn\":[{\"tag\":\"div\",\"text\":{\"tag\":\"lark_md\",\"content\":\"太强啦！！！\"}},{\"tag\":\"action\",\"actions\":[{\"tag\":\"button\",\"text\":{\"tag\":\"plain_text\",\"content\":\"给邓佬点赞\"},\"type\":\"primary\",\"value\":{\"key\":\"primary\"}},{\"tag\":\"button\",\"text\":{\"tag\":\"plain_text\",\"content\":\"给邓佬投币\"},\"type\":\"default\",\"value\":{\"key\":\"default\"}}]}]}}"
	card := map[string]interface{}{}
	err := json.Unmarshal([]byte(cardContent), &card)
	if err != nil {
		panic(err)
	}
	body := map[string]interface{}{
		"open_id":  "ou_273dbf68377bc685de3dd11c6102f879",
		"msg_type": "interactive",
		"card":     card,
	}
	ret := make(map[string]interface{})
	req := request.NewRequestWithNative("message/v4/send", "POST",
		request.AccessTokenTypeTenant, body, &ret,
	)
	err = api.Send(coreCtx, conf, req)
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(ret))
}

// send message
func SendMessage() {
	coreCtx := core.WrapContext(context.Background())
	body := map[string]interface{}{
		"open_id":  "ou_273dbf68377bc685de3dd11c6102f879",
		"msg_type": "text",
		"content": map[string]interface{}{
			"text": "test",
		},
	}
	ret := make(map[string]interface{})
	req := request.NewRequestWithNative("message/v4/send", "POST",
		request.AccessTokenTypeTenant, body, &ret,
	)
	err := api.Send(coreCtx, conf, req)
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(ret))
}

type UploadImage struct {
	ImageKey string `json:"image_key"`
}

// upload image
func UploadFile() {
	coreCtx := core.WrapContext(context.Background())
	// coreCtx.Set(constants.HTTPHeaderKeyRequestID, "2020122212081301001702714534518-xxxxx")
	var formData = request.NewFormData()
	formData.AddParam("image_type", "message")
	bs, err := ioutil.ReadFile("test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	formData.AddFile("image", request.NewFile().SetContent(bs))
	/*
		// stream upload, file implement io.Reader
		file, err := os.Open("test.png")
		if err != nil {
			fmt.Println(err)
			return
		}
		formData.AddFile("image", request.NewFile().SetContentStream(file))
	*/
	ret := &UploadImage{}
	err = api.Send(coreCtx, conf, request.NewRequestWithNative("image/v4/put", "POST",
		request.AccessTokenTypeTenant, formData, ret))
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(ret))
}

// download image
func DownloadFile() {
	coreCtx := core.WrapContext(context.Background())
	ret := &bytes.Buffer{}
	/*
		// stream download: ret implement io.Writer
		ret, err := os.Create("[file path]")
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	req := request.NewRequestWithNative("image/v4/get", "GET",
		request.AccessTokenTypeTenant, nil, ret,
		request.SetQueryParams(map[string]interface{}{"image_key": "[image key]"}), request.SetResponseStream())
	err := api.Send(coreCtx, conf, req)
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	err = ioutil.WriteFile("test_download.png", ret.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}
