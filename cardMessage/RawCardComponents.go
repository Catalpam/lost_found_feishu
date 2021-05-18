package cardMessage

import "fmt"

func ReSurveyButton(lostId string) string {
	return fmt.Sprintf(ReSurveyButtonRaw,lostId,lostId,lostId)
}
func ReSuspectedButton(lostId string,foundId string) string {
	return fmt.Sprintf(ReSuspectedButtonRaw,lostId,foundId,lostId,foundId)
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
const ReSuspectedButtonRaw = `
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
		  "FoundId": "%s",
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
		"buttonValue": {
		  "LostId": "%s"
		}
	  },
	  "type": "primary"
	}
  ]
}
`



