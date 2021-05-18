package cardMessage

import (
	"fmt"
	"net/url"
	"strconv"
)

func ReSurveyButton(lostId string) string {
	return fmt.Sprintf(ReSurveyButtonRaw,lostId,lostId,lostId)
}
func SuspectedButton(lostId uint,foundId uint) string {
	var lostIdStr string =  strconv.FormatUint(uint64(lostId), 10)
	var foundIdStr string =  strconv.FormatUint(uint64(foundId), 10)

	// Let's start with a base url

	baseUrlStr := "https://applink.feishu.cn/client/mini_program/open?appId=cli_a00d67c8e5f8500c&mode=window&path="
	indexUrl,err := url.Parse("pages/index")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		print("\nIndexUrl编码严重错误：")
		print(err)
		print("\n")
	}
	params := url.Values{}
	params.Add("lost_id", lostIdStr)
	params.Add("found_id", foundIdStr)
	// Add Query Parameters to the URL
	indexUrl.RawQuery = params.Encode() // Escape Query Parameters
	fmt.Printf("跳转的链接为：URL is %q\n", baseUrlStr+indexUrl.String())
	encodeUrl:= url.QueryEscape(indexUrl.String())
	fmt.Printf("跳转的转码链接为：Encoded URL is %q\n", baseUrlStr+encodeUrl)
	return fmt.Sprintf(SuspectedButtonRaw,baseUrlStr+encodeUrl,lostIdStr,foundIdStr)
}

func LostAddedButton(lostId string) string {
	return fmt.Sprintf(LostAddedButtonRaw,lostId)
}


//ReSurvey卡片按钮
const ReSurveyButtonRaw = `
{
  "tag": "action",
  "layout": "bisected",
  "actions": [
	{
	  "tag": "button",
	  "text": {
		"tag": "plain_text",
		"content": "是"
	  },
	  "value": {
		"buttonType": "survey",
		"buttonValue": {
		  "LostId": "%s",
		  "IsTrue": "T"
		}
	  },
	  "type": "primary"
	},
	{
	  "tag": "button",
	  "text": {
		"tag": "plain_text",
		"content": "否"
      },
	  "value": {
		"buttonType": "survey",
		"buttonValue": {
		  "LostId": "%s",
		  "IsTrue": "F"
		}
	  },
	  "type": "primary"
	}
  ]
},
{
  "tag": "action",
  "actions": [
	{
	  "tag": "button",
	  "text": {
		"tag": "plain_text",
		"content": "撤销查询"
	  },
	  "value": {
		"buttonType": "survey",
		"buttonValue": {
		  "LostId": "%s",
		  "IsTrue": "Revoke"
		}
	  },
	  "type": "default"
	}
  ]
}
`

//Suspected卡片按钮
const SuspectedButtonRaw = `
{
  "tag": "action",
  "layout": "bisected",
  "actions": [
	{
	  "tag": "button",
	  "text": {
		"tag": "plain_text",
		"content": "是"
	  },
	  "url":"%s",
	  "type": "primary"
	},
	{
	  "tag": "button",
	  "text": {
		"tag": "plain_text",
		"content": "否"
      },
	  "value": {
		"buttonType": "suspected",
		"buttonValue": {
		  "LostId": "%s",
		  "FoundId": "%s",
		  "IsTrue": "F"
		}
	  },
	  "type": "primary"
	}
  ]
}
`

//LostAdded卡片按钮
const LostAddedButtonRaw = `
{
  "tag": "action",
  "layout": "bisected",
  "actions": [
	{
	  "tag": "button",
	  "text": {
		"tag": "plain_text",
		"content": "已找到"
	  },
	  "value": {
		"buttonType": "cancelAdded",
		"LostId": "%s"
	  },
	  "type": "primary"
	}
  ]
}
`



