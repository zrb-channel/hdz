package hdz

import (
	"context"
	json "github.com/json-iterator/go"
	"io"
	"net/http"
)

type (
	HandlerFunc func(context.Context, string, []byte) error

	ApprovalStatus struct {
		State uint8
		Text  string
	}
)

var (
	handlers = make(map[string]HandlerFunc)
)

type NotifyHandlers interface {
	OnCreate(ctx context.Context, orderNo string, body *CreateOrderResponse) error

	OnApproval(ctx context.Context, orderNo string, body *ApprovalOrderResponse) error

	OnLoanApply(ctx context.Context, orderNo string, body []byte) error

	OnLoan(ctx context.Context, orderNo string, body *OrderLoanResponse) error

	OnCheckAgreement(ctx context.Context, orderNo string, body []byte) error
}

var notifyHandlers NotifyHandlers

func init() {
	handlers["APPLY_CREATE_STATUS"] = OnCreate         // 订单创建
	handlers["APPROVAL_RESULT"] = OnApproval           // 订单审批
	handlers["LOAN_APPLY"] = OnLoanApply               // 申请放款
	handlers["LOAN_RESULT"] = OnLoan                   // 放款成功
	handlers["USER_CHECKS_AGREEMENT"] = CheckAgreement // 放款成功
}

func RegisterNotifyHandlers(handlers NotifyHandlers) {
	notifyHandlers = handlers
}

// Notify
// @param req
// @date 2022-09-21 16:31:29
func Notify(req *http.Request) error {
	base := &NotifyBaseResponse{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, base); err != nil {
		return err
	}

	ctx := req.Context()
	handler, ok := handlers[base.EventType]
	if ok {
		return handler(ctx, base.OrderNo, base.Data)
	}

	return nil
}

// onNotify
// @param ctx
// @param msg
// @date 2022-09-21 16:31:28
func onNotify(ctx context.Context, msg []byte) error {
	result := &NotifyBaseResponse{}
	if err := json.Unmarshal(msg, result); err != nil {
		return err
	}

	handler, ok := handlers[result.EventType]
	if ok {
		return handler(ctx, result.OrderNo, result.Data)
	}

	return nil
}

// OnCreate
// @date: 2022-03-21 20:23:07
func OnCreate(ctx context.Context, orderNo string, msg []byte) error {
	result := &CreateOrderResponse{}
	if err := json.Unmarshal(msg, result); err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}
	return notifyHandlers.OnCreate(ctx, orderNo, result)
}

// OnApproval
// @date: 2022-03-24 01:13:58
func OnApproval(ctx context.Context, orderNo string, msg []byte) error {
	result := &ApprovalOrderResponse{}
	if err := json.Unmarshal(msg, result); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}
	return notifyHandlers.OnApproval(ctx, orderNo, result)
}

// OnLoanApply
// @date: 2022-03-24 01:13:59
func OnLoanApply(ctx context.Context, orderNo string, msg []byte) error {
	if notifyHandlers == nil {
		return nil
	}
	return notifyHandlers.OnLoanApply(ctx, orderNo, msg)
}

// OnLoan
// @date: 2022-03-24 01:14:01
func OnLoan(ctx context.Context, orderNo string, msg []byte) error {
	result := &OrderLoanResponse{}
	if err := json.Unmarshal(msg, result); err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}
	return notifyHandlers.OnLoan(ctx, orderNo, result)
}

// CheckAgreement
// @param ctx
// @param orderNo
// @param _
// @date 2022-09-21 16:24:23
func CheckAgreement(ctx context.Context, orderNo string, msg []byte) error {
	if notifyHandlers == nil {
		return nil
	}
	return notifyHandlers.OnCheckAgreement(ctx, orderNo, msg)
}
