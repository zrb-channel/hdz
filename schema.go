package hdz

type (
	Config struct {
		LoginToken string `json:"loginToken"`

		LoginAppId string `json:"loginAppId"`

		AppId string `json:"appId"`

		Token string `json:"token"`
	}

	LoginRequest struct {
		AppID string `json:"appId"`

		Phone string `json:"phone"`

		ApplicationNum string `json:"applicationNum"`

		Flag string `json:"flag"`

		Sign string `json:"sign"`
	}

	BaseRequest[T any] struct {
		EventType string `json:"eventType"`

		ApplicationNum string `json:"applicationNum"`

		Data T `json:"data"`
	}

	CreateOrderRequest struct {
		// Applicant 申请人姓名 ，最少两个汉字
		Applicant string `json:"applicant"`

		// IdCardNumber 申请人身份证号， 15位或18位
		IdCardNumber string `json:"idCardNumber"`

		// PhoneNumber 申请手机号
		PhoneNumber string `json:"phoneNumber"`

		// CompanyName 公司名称
		CompanyName string `json:"companyName"`

		// UnifiedSocialCreditCode 统一社会信用代码
		UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode"`
	}

	BaseResponse[T any] struct {
		EventType string `json:"eventType"`

		Code string `json:"code"`

		Message string `json:"message"`

		Timestamp int64 `json:"timestamp"`

		ApplicationNum string `json:"applicationNum"`

		PhoneNumber string `json:"phoneNumber"`

		Data T `json:"data"`
	}

	CreateOrderResponse struct {
		ApplyCode string `json:"applyCode"`

		Message string `json:"message"`

		Status string `json:"status"`
	}
)
