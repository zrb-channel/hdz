package hdz

import (
	"context"
	"errors"
	"github.com/zrb-channel/hdz/config"
	"github.com/zrb-channel/utils"

	json "github.com/json-iterator/go"
)

// CreateOrder
// @param ctx
// @param orderNo
// @param body
// @date 2022-09-21 16:23:46
func CreateOrder(ctx context.Context, orderNo string, body *CreateOrderRequest) (*CreateOrderResponse, error) {

	req := BaseRequest[*CreateOrderRequest]{
		EventType:      "APPLY_CREATE",
		ApplicationNum: orderNo,
		Data:           body,
	}

	headers := map[string]string{
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": config.Token,
		"App-Id":        config.AppID,
	}

	resp, err := utils.Request(ctx).
		SetBody(req).
		SetHeaders(headers).
		Post(config.Addr)

	if err != nil {
		return nil, err
	}

	result := &BaseResponse[*CreateOrderResponse]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}

	if result.Code != "200" {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
