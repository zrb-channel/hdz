package hdz

import (
	json "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

type (
	NotifyBaseResponse struct {
		EventType string `json:"eventType"`
		Code      string `json:"code"`

		Message string `json:"message"`

		Timestamp int64 `json:"timestamp"`

		OrderNo string `json:"applicationNum"`
		
		Mobile string `json:"phoneNumber"`

		Data json.RawMessage `json:"data"`
	}

	NotifyCreateOrderResponse struct {
		ApplyCode string `json:"applyCode"`

		Message string `json:"message"`

		Status string `json:"status"`
	}

	ApprovalOrderResponse struct {
		Amount decimal.Decimal `json:"amount"`

		FileCategories string `json:"fileCategories"`

		RateLevel string `json:"rateLevel"`

		RejectReason string `json:"rejectReason"`

		Remark string `json:"remark"`

		RepayType string `json:"repayType"`

		Result string `json:"result"`
	}

	OrderLoanResponse struct {
		Result string `json:"result"`

		DisburseCode string `json:"disburseCode"`

		LoanDate string `json:"loanDate"`

		Amount decimal.Decimal `json:"amount"`

		Period int `json:"period"`

		RepayType string `json:"repayType"`

		RepayCard string `json:"repayCard"`

		RepayBank string `json:"repayBank"`
	}
)
