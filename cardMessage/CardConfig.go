package cardMessage

import "fmt"

type FeedbackUrlModel struct {
	HasClaim2Founder string
	HasSendThanks    string
}

var FeedbackUrl = FeedbackUrlModel{
	HasClaim2Founder: "https://wenjuan.feishu.cn/m?t=s4Ihv8ra93ri-euhz",
	HasSendThanks:    "https://wenjuan.feishu.cn/m?t=s4Ihv8ra93ri-euhz",
}

func Str2MultiUrl(url string) string {
	return fmt.Sprintf("\"multi_url\": {\n\t\"url\": \"%s\",\n\t\"android_url\": \"%s\",\n\t\"ios_url\": \"%s\",\n\t\"pc_url\": \"%s\"\n},", url, url, url, url)
}
