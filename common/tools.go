package common

func TimeSessionToChinese(timeSession string) string {
	timeSessionInChinese := ""
	switch timeSession {
	case "morning": timeSessionInChinese = "上午（6：00-11：00）"
	case "noon": timeSessionInChinese = "中午（11：00-2：00）"
	case "afternoon": timeSessionInChinese = "下午（2：00-19：00）"
	case "evening": timeSessionInChinese = "晚上（19：00-22：00）"
	case "night": timeSessionInChinese = "夜间（00：00-6：00，22：00-24：00）"
	}
	return timeSessionInChinese
}
