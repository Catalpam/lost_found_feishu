package cardMessage

func ReSurveyButton(lostId string) string {
	return ""
}

const ReSurveyButtonRaw =
`
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