package cardMessage

import "fmt"

type FeedbackUrlModel struct {
	HasClaim2Founder string
	HasSendThanks    string
}

var FeedbackUrl = FeedbackUrlModel{
	HasClaim2Founder: "https://wenjuan.feishu.cn/m?t=szwWLUIernsi-a1rz",
	HasSendThanks:    "https://wenjuan.feishu.cn/m?t=szwWLUIernsi-a1rz",
}

func Str2MultiUrl(url string) string {
	return fmt.Sprintf("\"multi_url\": {\n\t\"url\": \"%s\",\n\t\"android_url\": \"%s\",\n\t\"ios_url\": \"%s\",\n\t\"pc_url\": \"%s\"\n},", url, url, url, url)
}
