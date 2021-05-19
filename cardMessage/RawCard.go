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
        "content": "已根据您之前添加的的Lost查询到如下Found招领信息，请问是您的物品吗？"
      }
    },
    {
      "tag": "img",
      "title": {
        "tag": "lark_md",
        "content": "物品种类：%s\n发现地点：%s\n发现时间：%s"
      },
      "img_key": "%s",
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

const rawSameNameCard = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "有人拾到了含有您姓名的物品，详细信息如下"
      }
    },
    {
      "tag": "hr"
    },
    {
      "tag": "img",
      "title": {
        "tag": "lark_md",
        "content": "拾取地点：%s"
      },
      "img_key": "%s",
      "alt": {
        "tag": "plain_text",
        "content": "图片"
      }
    },
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "若该物品为您遗失的物品，您可以点击以下按钮获取链接，进行领取，若并非您的物品，请忽略本消息。"
      },
	%s
    }
  ]
}
`