package imRobot

import (
	"context"
	"encoding/json"
	"fmt"
	"lost_found/api"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	im "lost_found/service/im/v1"
)

var conf = configs.FeishuConfig(constants.DomainFeiShu)
var imService = im.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func ImSendUser(ReceiveId string,ShareId string) {
	coreCtx := core.WrapContext(context.Background())
	shareContent := "{\"user_id\":\"" + ShareId + "\"}"
	reqCall := imService.Messages.Create(coreCtx, &im.MessageCreateReqBody{
		ReceiveId: ReceiveId,
		Content:   shareContent,
		MsgType:   "share_user",
	})
	reqCall.SetReceiveIdType("open_id")
	message, err := reqCall.Do()
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return
	}
	fmt.Println(tools.Prettify(message))
}

func SendCardMessage() {
	coreCtx := core.WrapContext(context.Background())
	cardContent := "{\"config\":{\"wide_screen_mode\":true},\"i18n_elements\":{\"zh_cn\":[{\"tag\":\"div\",\"text\":{\"tag\":\"lark_md\",\"content\":\"[飞书](https://www.feishu.cn)整合即时沟通、日历、音视频会议、云文档、云盘、工作台等功能于一体，成就组织和个人，更高效、更愉悦。\"}},{\"tag\":\"action\",\"actions\":[{\"tag\":\"button\",\"text\":{\"tag\":\"plain_text\",\"content\":\"主按钮\"},\"type\":\"primary\",\"value\":{\"key\":\"primary\"}},{\"tag\":\"button\",\"text\":{\"tag\":\"plain_text\",\"content\":\"次按钮\"},\"type\":\"default\",\"value\":{\"key\":\"default\"}}]}]}}"
	card := map[string]interface{}{}
	err := json.Unmarshal([]byte(cardContent), &card)
	if err != nil {
		panic(err)
	}
	body := map[string]interface{}{
		"open_id":  "77bbc392",
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
func SendMessage(OpenId string, Text string) {
	coreCtx := core.WrapContext(context.Background())
	body := map[string]interface{}{
		"open_id":  OpenId,
		"msg_type": "text",
		"content": map[string]interface{}{
			"text": Text,
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

func SendImage(OpenId string, Key string) {
	coreCtx := core.WrapContext(context.Background())
	body := map[string]interface{}{
		"open_id":  OpenId,
		"msg_type": "image",
		"content": map[string]interface{}{
			"image_key": Key,
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
