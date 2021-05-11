package cardMessage

import (
	"context"
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

func ImSendUser(ReceiveId string, ShareId string) {
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
