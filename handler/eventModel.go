package handler

type Message struct {
	Event *struct {
		Type string `json:"type"`
		MsgType string `json:"msg_type"`
		OpenId string `json:"open_id"`
		UserAgent string `json:"user_agent"`
		ImageKey string `json:"image_key"`
		ImageUrl string `json:"image_url"`
		Text string `json:"text"`
		OpenMessageId string `json:"open_message_id"`
	} `json:"event"`
}