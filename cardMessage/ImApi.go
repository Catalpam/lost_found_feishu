package cardMessage

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

func SendCardMessage(OpenID string, cardContent string) {
	coreCtx := core.WrapContext(context.Background())
	card := map[string]interface{}{}
	println(cardContent)
	err := json.Unmarshal([]byte(cardContent), &card)
	if err != nil {
		panic(err)
	}
	body := map[string]interface{}{
		"open_id":  OpenID,
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

func FoundClaimCard(foundClaim FoundClaim) string {
	var cardContent = "{\"a\":\"a\"}"
	formatCard := "{\n\t\"header\": {\n\t\t\"title\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"招领信息被认领\"\n\t\t},\n\t\t\"template\": \"green\"\n\t},\n\t\"config\": {\n\t\t\"wide_screen_mode\": true\n\t},\n\t\"elements\": [{\n\t\t\"tag\": \"div\",\n\t\t\"text\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"感谢您的热心，您发布的招领信息已有失主认领啦！\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"img\",\n\t\t\"title\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"物品种类：%s\\n发现时间：2020/3/20\"\n\t\t},\n\t\t\"img_key\": \"%s\",\n\t\t\"alt\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"图片\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"hr\"\n\t}, {\n\t\t\"tag\": \"div\",\n\t\t\"text\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"%s\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"hr\"\n\t}, {\n\t\t\"tag\": \"div\",\n\t\t\"text\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"再次感谢您的热心相助，也欢迎您给我们留言，反馈您本次使用的体验~\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"action\",\n\t\t\"actions\": [{\n\t\t\t\"tag\": \"button\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\"content\": \"提交反馈\"\n\t\t\t},\n\t\t\t\"type\": \"primary\"\n\t\t}]\n\t}]\n}"
	println(formatCard)
	cardContent = fmt.Sprintf(formatCard, foundClaim.ItemSubtype, foundClaim.ImageKey, foundClaim.LeaveMessage)
	println(cardContent)
	return cardContent
}

type FoundClaim struct {
	ItemSubtype  string
	ImageKey     string
	LeaveMessage string
}
