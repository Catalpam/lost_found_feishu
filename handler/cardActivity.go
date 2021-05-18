package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lost_found/card"
	"lost_found/card/http"
	"lost_found/card/model"
	"lost_found/cardMessage"
	"lost_found/comander"
	"lost_found/core"
	"lost_found/core/configs"
	"lost_found/core/constants"
	"lost_found/core/tools"
	"strconv"
)

func CardEvent() gin.HandlerFunc {

	var conf = configs.FeishuConfig(constants.DomainFeiShu)

	card.SetHandler(conf, func(coreCtx *core.Context, card *model.Card) (interface{}, error) {
		fmt.Println(tools.Prettify(card.Action))
		// 当回调控件为按钮时：
		if card.Action.Tag == "button"{
			ButtonHandler(card.Action.Value,card.OpenID)
		}
		
		return nil, nil
	})

	return func(context *gin.Context) {
		http.Handle(conf, context.Request, context.Writer)
	}
}

func ButtonHandler(value map[string]interface{}, openId string) {
	switch value["buttonType"].(string) {
	case ButtonCardType.Survey:
		println("-------------收到ButtonSurvey回调！--------------------")
		buttonValue := value["buttonValue"].(map[string]interface{})
		IsTrue := buttonValue["IsTrue"].(string)
		//var LostId := buttonValue["LostId"]
		if IsTrue == "T" {
			println("-------------按下了\"是\"按钮！--------------------")
			cardMessage.SendMessage(openId,"好的，我们将撤销您的lost查询信息，恭喜您找到了您的失物。")
		} else if IsTrue == "F" {
			println("-------------按下了\"否\"按钮！--------------------")
			cardMessage.SendMessage(openId,"好的，我们将为您继续查询。")
		} else if IsTrue == "Revoke" {
			println("-------------按下了\"撤销\"按钮！--------------------")
			cardMessage.SendCardMessage(openId, cardMessage.RevokeLostCard())
		}
	case ButtonCardType.Suspected:
		buttonValue := value["buttonValue"].(map[string]interface{})
		IsTrue := buttonValue["IsTrue"].(string)
		//var LostId := buttonValue["LostId"]
		if IsTrue == "T" {
			//调试用，这里要修改的
			cardMessage.SendMessage(openId,"很开心能帮助您查找到您的物品，单击下方按钮跳转至小程序查看领取方式吧。")
		} else if IsTrue == "F" {
			//调试用，这里要修改的
			cardMessage.SendMessage(openId,"好的，我们将为您继续查询。")
		}
	case ButtonCardType.SameName:
		buttonValue := value["buttonValue"].(ButtonSameNameValueModel)
		if buttonValue.IsTrue == "T" {
			//调试用，这里要修改的
			cardMessage.SendMessage(openId,"很开心能帮助您查找到您的物品，单击下方按钮跳转至小程序查看领取方式吧。")
		} else if buttonValue.IsTrue == "F" {
			//调试用，这里要修改的
			cardMessage.SendMessage(openId,"好的，打扰您了。")
		}

	case ButtonCardType.LostAdded:
		print("\n主动修改为已找到,LostId为：")
		LostIdStr := value["LostId"].(string)
		LostId, _ := strconv.ParseUint(LostIdStr, 10, 64)
		print("\n主动修改为已找到,LostId为：")
		comander.HasFounded(LostId)
	}
}

//带按钮的卡片类型及其实例
type ButtonCardTypeModel struct {
	Survey string
	Suspected string
	SameName string
	LostAdded string
}
var ButtonCardType = ButtonCardTypeModel{
	Survey:    "survey",
	Suspected: "suspected",
	SameName:  "same_name",
	LostAdded: "cancelAdded",
}

// 每种带按钮卡片的按钮回调
type ButtonSurveyValueModel struct {
	LostId string
	IsTrue string  //T F Revoke
}

type ButtonSuspectedValueModel struct {
	LostId string
	FoundId string
	IsTrue string
}

type ButtonSameNameValueModel struct {
	FoundId string
	IsTrue string
}
type ButtonCancelAddedModel struct {
	LostId string
}

