package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"lost_found/event/core/model"
	"lost_found/imRobot"
	im "lost_found/service/im/v1"
	"net/http"
)

var conf = configs.FeishuConfig(constants.DomainFeiShu)
var imService = im.NewService(configs.FeishuConfig(constants.DomainFeiShu))

func EventHandler(ctx *gin.Context)  {
	body,_ := ioutil.ReadAll(ctx.Request.Body)
	var err error
	fmt.Println("-------未解密前body/------- \r\n "+string(body))

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
	fmt.Println("-------使用EncryptKey解密后body/------- \r\n "+string(body))

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

	//验证Token是否有效
	if fuzzy.Token != "UFkMwVyQpoMzwRmtkHgmFbFhFy0HioQ1" {
		ctx.JSON( http.StatusBadRequest,gin.H{
			"msg":"Bad Token",
		})
		return
	}

	//当在飞书开发者后台设置事件响应URL时的Challenge响应
	if fuzzy.Type == "url_verification" {
		ctx.JSON( http.StatusOK,gin.H{
			"challenge": fuzzy.Challenge,
		})
		return
	}

	//解析事件类型eventType
	var eventType string
	if fuzzy.Event != nil {
		eventType = fuzzy.Event.Type
	}
	if fuzzy.Header != nil {
		eventType = fuzzy.Header.EventType
	}
	fmt.Println("-------eventType/------- \r\n "+string(eventType))

	// 当eventType为message时的响应
	if eventType == "message" {
		err = messageWebhook(body)
		ctx.JSON(200,gin.H{"msg":"Succeed"})
		return
	}

	// 当eventType为 时的响应
	if eventType == ""{

	}

	//未知的Event类型响应
	ctx.JSON(200,gin.H{"msg":"Unknown Event Class"})
}

func messageWebhook (body []byte) error{
	content := &Message{}
	var err error

	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if content.Event.MsgType == "image" {
		imRobot.SendMessage(content.Event.OpenId, "坏耶，图片有些小问题:" + content.Event.Text)
		imRobot.SendImage(content.Event.OpenId, content.Event.ImageKey)
	}
	if content.Event.MsgType == "text" {
		imRobot.SendMessage(content.Event.OpenId, "好耶，Your Msg is:" + content.Event.Text)
	}
	if content.Event.MsgType == "audio" {
		imRobot.SendMessage(content.Event.OpenId, "阿哦，相信机器人，大家是不会听语音的")
	}
	return nil
}