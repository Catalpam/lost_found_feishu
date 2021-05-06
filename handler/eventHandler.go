package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lost_found/api"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/event/core/model"
	"net/http"
)

var conf = configs.FeishuConfig(constants.DomainFeiShu)

func EventHandler(ctx *gin.Context)  {
	body,_ := ioutil.ReadAll(ctx.Request.Body)
	var err error

	//EncryptKey解密
	encryptKey := "Vls2mIJJSNREL8lK3Y42jhDIllPXZ88x"
	if encryptKey != "" {
		body, err = tools.Decrypt(body, encryptKey)
		if err != nil {
			fmt.Println(err)
			ctx.JSON( http.StatusBadRequest,gin.H{
				"msg":err,
			})
			return
		}
	}
	fmt.Println("---encrypt/--- \r\n "+string(body))

	//Json Decode
	fuzzy := &model.Fuzzy{}
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&fuzzy)
	if err != nil {
		fmt.Println(err)
		ctx.JSON( http.StatusBadRequest,gin.H{
			"msg":err,
		})
		return
	}

	//Model Initial
	//schema := model.Version1
	//if fuzzy.Schema != "" {
	//	schema = fuzzy.Schema
	//}

	//Token Auth
	if fuzzy.Token != "UFkMwVyQpoMzwRmtkHgmFbFhFy0HioQ1" {
		ctx.JSON( http.StatusBadRequest,gin.H{
			"msg":"Bad Token",
		})
		return
	}

	var eventType string
	if fuzzy.Event != nil {
		eventType = fuzzy.Event.Type
	} else if fuzzy.Header != nil {
		eventType = fuzzy.Header.EventType
	}

	if eventType == "url_verification" {
		ctx.JSON( http.StatusOK,gin.H{
			"challenge": fuzzy.Challenge,
		})
		return
	} else if eventType == "message"{
		content := &Message{}
		err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&content)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
			return
		}
		if content.Event.MsgType == "image" {

			sendMessage(content.Event.OpenId, "坏耶，图片有些小问题:" + content.Event.Text)
			sendImage(content.Event.OpenId, content.Event.ImageKey)

		}
		if content.Event.MsgType == "text" {
			sendMessage(content.Event.OpenId, "好耶，Your Msg is:" + content.Event.Text)
		}
		if content.Event.MsgType == "audio" {
			sendMessage(content.Event.OpenId, "阿哦，相信机器人，大家是不会听语音的")
		}


		//sendMessage(content.Event.OpenId, "好耶，Your Msg is:" + content.Event.Text)
		//sendMessage("ou_c5b13c8bae4e65657d1d89e6dbbfc098", "整个五一可算研究了一些东西，太菜了，现在后端这里可以自动保存与机器人聊天记录了，加了个小彩蛋，你发个消息，他会回你哦")
		//sendMessage(content.Event.OpenId, "整个五一可算研究了一些东西，太菜了，现在后端这里可以自动保存与机器人聊天记录了，加了个小彩蛋，你发个消息，他会回你哦")

	}
	print("\n\n\n\n\n\n\n\\n\n\n")
	println(eventType)

	ctx.JSON(200,gin.H{
		"receive":"success",
	})

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
func sendMessage(OpenId string, Text string) {
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

func sendImage(OpenId string, Key string) {
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

func SendUser() {
	coreCtx := core.WrapContext(context.Background())
	body := map[string]interface{}{
		"open_id":  "ou_273dbf68377bc685de3dd11c6102f879",
		"msg_type": "share_user",
		"content": map[string]interface{}{
			"user_id": "ou_c5b13c8bae4e65657d1d89e6dbbfc098",
			//"text":"111",
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



