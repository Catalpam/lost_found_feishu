package cardMessage

import "fmt"

func FoundClaimCard(cardModel FoundClaim) string {
	var cardContent = "{\"a\":\"a\"}"
	formatRawCard := "{\n\t\"header\": {\n\t\t\"title\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"招领信息被认领\"\n\t\t},\n\t\t\"template\": \"green\"\n\t},\n\t\"config\": {\n\t\t\"wide_screen_mode\": true\n\t},\n\t\"elements\": [{\n\t\t\"tag\": \"div\",\n\t\t\"text\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"感谢您的热心，您发布的招领信息已有失主认领啦！\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"img\",\n\t\t\"title\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"物品种类：%s\\n发现时间：2020/3/20\"\n\t\t},\n\t\t\"img_key\": \"%s\",\n\t\t\"alt\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"图片\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"hr\"\n\t}, {\n\t\t\"tag\": \"div\",\n\t\t\"text\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"来自失主的感谢信息：%s\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"hr\"\n\t}, {\n\t\t\"tag\": \"div\",\n\t\t\"text\": {\n\t\t\t\"tag\": \"lark_md\",\n\t\t\t\"content\": \"再次感谢您的热心相助，也欢迎您给我们留言，反馈您本次使用的体验~\"\n\t\t}\n\t}, {\n\t\t\"tag\": \"action\",\n\t\t\"actions\": [{\n\t\t\t\"tag\": \"button\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\"content\": \"提交反馈\"\n\t\t\t},\n\t\t\t%s\"type\": \"primary\"\n\t\t}]\n\t}]\n}"
	cardContent = fmt.Sprintf(formatRawCard, cardModel.ItemSubtype, cardModel.ImageKey, cardModel.LeaveMessage, Str2MultiUrl(FeedbackUrl.HasClaim2Founder))
	return cardContent
}

type FoundClaim struct {
	ItemSubtype  string
	ImageKey     string
	LeaveMessage string
}

func ThanksHasSendCard(cardModel ThanksHasSend) string {
	var cardContent = "{\"a\":\"a\"}"
	formatRawCard := "{\n\t\"config\": {\n\t\t\"wide_screen_mode\": true\n\t},\n\t\"header\": {\n\t\t\"title\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"感谢留言已发送\"\n\t\t},\n\t\t\"template\": \"blue\"\n\t},\n\t\"elements\": [{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"您的感谢留言已发送给found发布者~\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"img\",\n\t\t\t\"title\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"物品种类：%s\\n发现时间：%s\"\n\t\t\t},\n\t\t\t\"img_key\": \"%s\",\n\t\t\t\"alt\": {\n\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\"content\": \"图片\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"hr\"\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"很开心我们的系统能够帮助您找到失物，也欢迎您给我们留言，反馈您本次使用的体验\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"action\",\n\t\t\t\"actions\": [{\n\t\t\t\t\"tag\": \"button\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\t\"content\": \"提交反馈\"\n\t\t\t\t},\n\t\t\t\t%s\n\t\t\t\t\"type\": \"primary\"\n\t\t\t}]\n\t\t}\n\t]\n}"
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, cardModel.ItemSubtype, cardModel.FoundDate, cardModel.ImageKey, Str2MultiUrl(FeedbackUrl.HasSendThanks))
	//println(cardContent)
	return cardContent
}

type ThanksHasSend struct {
	ItemSubtype string
	FoundDate   string
	ImageKey    string
}

func SendUser2FounderCard(cardModel SendUser2Founder) string {
	var cardContent = "{\"a\":\"a\"}"
	formatRawCard := "{\n\t\"config\": {\n\t\t\"wide_screen_mode\": true\n\t},\n\t\"header\": {\n\t\t\"title\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"招领消息已认领，请联系失主\"\n\t\t},\n\t\t\"template\": \"green\"\n\t},\n\t\"elements\": [{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"您发布的如下招领信息已被人认领啦，您可以通过以下名片与失主联系。\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"img\",\n\t\t\t\"title\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"物品种类：%s\\n发现时间: %s\"\n\t\t\t},\n\t\t\t\"img_key\": \"%s\",\n\t\t\t\"alt\": {\n\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\"content\": \"图片\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"hr\"\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"再次感谢您的热心相助，也欢迎您给我们留言，反馈您本次使用的体验~\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"action\",\n\t\t\t\"actions\": [{\n\t\t\t\t\"tag\": \"button\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\t\"content\": \"提交反馈\"\n\t\t\t\t},\n\t\t\t\t%s\n\t\t\t\t\"type\": \"primary\"\n\t\t\t}]\n\t\t}\n\t]\n}"
	println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, cardModel.ItemSubtype, cardModel.FoundDate, cardModel.ImageKey, Str2MultiUrl(FeedbackUrl.HasSendThanks))
	println(cardContent)
	return cardContent
}

type SendUser2Founder struct {
	ItemSubtype string
	FoundDate   string
	ImageKey    string
}

func SendUser2LosterCard(cardModel SendUser2Loster) string {
	var cardContent = "{\"a\":\"a\"}"
	formatRawCard := "{\n\t\"config\": {\n\t\t\"wide_screen_mode\": true\n\t},\n\t\"header\": {\n\t\t\"title\": {\n\t\t\t\"tag\": \"plain_text\",\n\t\t\t\"content\": \"您已认领招领消息，请联系拾物者\"\n\t\t},\n\t\t\"template\": \"green\"\n\t},\n\t\"elements\": [{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"您的物品目前由%s保管，您可以通过以下名片与TA联系。\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"img\",\n\t\t\t\"title\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"物品种类：%s\\n发现时间：%s\"\n\t\t\t},\n\t\t\t\"img_key\": \"%s\",\n\t\t\t\"alt\": {\n\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\"content\": \"图片\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"hr\"\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"很开心我们的系统能够帮助您找到失物，也欢迎您给我们留言，反馈您本次使用的体验~\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"action\",\n\t\t\t\"actions\": [{\n\t\t\t\t\"tag\": \"button\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\t\"content\": \"提交反馈\"\n\t\t\t\t},\n\t\t\t\t%s \n\t\t\t\t\"type\": \"primary\"\n\t\t\t}]\n\t\t}\n\t]\n}"
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, cardModel.FounderName, cardModel.ItemSubtype, cardModel.FoundDate, cardModel.ImageKey, Str2MultiUrl(FeedbackUrl.HasSendThanks))
	//println(cardContent)
	return cardContent
}

type SendUser2Loster struct {
	FounderName string
	ItemSubtype string
	FoundDate   string
	ImageKey    string
}

func RevokeLostCard() string {
	var cardContent = "{\"a\":\"a\"}"
	formatRawCard := "{\n\t\"config\": {\n\t\t\"wide_screen_mode\": true\n\t},\n\t\"elements\": [{\n\t\t\t\"tag\": \"div\",\n\t\t\t\"text\": {\n\t\t\t\t\"tag\": \"lark_md\",\n\t\t\t\t\"content\": \"很抱歉未能帮您解决问题，若您对本次使用体验感到不满，可以点击反馈按钮进行吐槽\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"tag\": \"action\",\n\t\t\t\"actions\": [{\n\t\t\t\t\"tag\": \"button\",\n\t\t\t\t\"text\": {\n\t\t\t\t\t\"tag\": \"plain_text\",\n\t\t\t\t\t\"content\": \"提交反馈\"\n\t\t\t\t},\n\t\t\t\t%s\n\t\t\t\t\"type\": \"primary\"\n\t\t\t}]\n\t\t}\n\t]\n}"
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, Str2MultiUrl(FeedbackUrl.HasSendThanks))
	//println(cardContent)
	return cardContent
}

func ReSurveyCard(cardmodel ReSurvey) string {
	var cardContent = ""
	formatRawCard := rawSurveyCard
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, cardmodel.ItemSubtype,cardmodel.FoundDate, cardmodel.ImageKey,ReSurveyButton(cardmodel.LostId))
	//println(cardContent)
	return cardContent
}
type ReSurvey struct {
	LostId 		string
	ItemSubtype string
	FoundDate   string
	ImageKey    string
}

func SuspectedCard(cardmodel Suspected) string {
	var cardContent = ""
	formatRawCard := rawSuspectedCard
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, cardmodel.ItemSubtype,cardmodel.FoundPlace,cardmodel.FoundDate, cardmodel.ImageKey,SuspectedButton(cardmodel.FoundId,cardmodel.FoundId))
	//println(cardContent)
	return cardContent
}
type Suspected struct {
	LostId 		uint
	FoundId 	uint
	ItemSubtype string
	FoundPlace string
	FoundDate   string
	ImageKey    string
}
func LostAddedCard(cardmodel LostAdded) string {
	var cardContent = ""
	formatRawCard := rawLostAddedCard
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard, cardmodel.ItemSubtype,cardmodel.LostDate, cardmodel.LostPlace,LostAddedButton(cardmodel.LostId))
	//println(cardContent)
	return cardContent
}
type LostAdded struct {
	LostId 		string
	ItemSubtype string
	LostDate    string
	LostPlace 	string
}

func SameNameCard(cardmodel SameName) string {
	var cardContent = ""
	formatRawCard := rawSameNameCard
	//println(formatRawCard)
	cardContent = fmt.Sprintf(formatRawCard,cardmodel.FoundPlace,cardmodel.ImageKey,SameNameUrlButton(cardmodel.FoundId))
	//println(cardContent)
	return cardContent
}
type SameName struct {
	FoundId 	uint
	FoundPlace string
	ImageKey    string
}






