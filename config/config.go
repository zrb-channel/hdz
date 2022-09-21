package config

const (
	// Addr 测试环境地址: https://hxduat.kungeek.com/api/hdz/biz
	// Addr 生产环境地址: https://hxdapi.kungeek.com/api/hdz/biz
	Addr = "https://hxdapi.kungeek.com/api/hdz/biz"

	Token = "95A8A72ADADD2024BB26E20ADC5E1BFF"
	AppID = "GDYSS"

	//LoginAppID="MDAwMDQ2NjU1IUlGExWP900"
	//LoginToken = "U2FsdGVkX1s2OMrMhUdaUCjKDyrUmIJ59FEQdLcyk8="
)

var (
	codeMapping = map[string]string{
		"201": "失败",
		"202": "参数不合法",
		"203": "未知的事件类型",
		"204": "未知的 App-Id",
		"205": "订单不存在",
		"206": "影像资料上传失败",
		"207": "订单不可放款",
		"208": "文件地址不可用",
		"209": "存疑(启信宝查询失败",
	}
)

func CodeMessage(code string) string {
	return codeMapping[code]
}

// https://hxduat.kungeek.com/mobile/homePage?noLoginCode=MDAwMDM4NDk1aZFzMbdrBDF*17606518462*DJXD2021052000001*87f14e2d76b49c9ba3b37cfa2e2673ec
