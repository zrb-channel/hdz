package hdz

type (
	Config struct {
		AppId string `json:"appId"`
		Token string `json:"token"`
	}

	LoginRequest struct {
		AppID          string `json:"appId"`
		Phone          string `json:"phone"`
		ApplicationNum string `json:"applicationNum"`
		Flag           string `json:"flag"`
		Sign           string `json:"sign"`
	}

	BaseRequest[T any] struct {
		EventType      string `json:"eventType"`
		ApplicationNum string `json:"applicationNum"`
		Data           T      `json:"data"`
	}

	CreateOrderRequest struct {
		Applicant               string `json:"applicant"`               // 申请人姓名 ，最少两个汉字
		IdCardNumber            string `json:"idCardNumber"`            // 申请人身份证号， 15位或18位
		PhoneNumber             string `json:"phoneNumber"`             // 申请手机号
		CompanyName             string `json:"companyName"`             // 公司名称
		UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode"` //  统一社会信用代码
	}

	BaseResponse[T any] struct {
		EventType      string `json:"eventType"`
		Code           string `json:"code"`
		Message        string `json:"message"`
		Timestamp      int64  `json:"timestamp"`
		ApplicationNum string `json:"applicationNum"`
		PhoneNumber    string `json:"phoneNumber"`
		Data           T      `json:"data"`
	}

	CreateOrderResponse struct {
		ApplyCode string `json:"applyCode"`
		Message   string `json:"message"`
		Status    string `json:"status"`
	}
)
