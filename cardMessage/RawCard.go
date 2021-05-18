package cardMessage

const rawSurveyCard = `
{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "失物查询回访"
    },
    "template": "orange"
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "您于7天前发布的如下lost信息尚未找到失物，请问您是否通过其他方式找到了呢？"
      }
    },
    {
      "tag": "img",
      "title": {
        "tag": "lark_md",
        "content": "物品种类：%s\n发现时间：%s"
      },
      "img_key": "%s",
      "alt": {
        "tag": "plain_text",
        "content": "失物图片"
      }
    },
	%s
  ]
}
`
const rawSuspectedCard = `
{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "疑似失物推送"
    },
    "template": "orange"
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "已查询到您如下失物查询的可能的招领信息，请问是您的物品吗？"
      }
    },
    {
      "tag": "img",
      "title": {
        "tag": "lark_md",
        "content": "物品种类：%s\n发现时间：%s"
      },
      "img_key": "img_e344c476-1e58-4492-b40d-7dcffe9d6dfg",
      "alt": {
        "tag": "plain_text",
        "content": "图片"
      }
    },
	%s
  ]
}
`
const rawLostAddedCard = `
{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "收到lost查询信息"
    },
    "template": "blue"
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "已收到您的lost信息，详情如下："
      }
    },
    {
      "tag": "hr"
    },
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "物品种类：%s\n时间：%s\n可能遗失地点: %s"
      }
    },
    {
      "tag": "hr"
    },
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "我们将实时查询，帮助您找到您的物品，若您自行找到遗失物品，也可以点击下方按钮，撤销查询"
      }
    },
	%s
  ]
}
`