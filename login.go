package hdz

import (
	"context"
	"fmt"
	"github.com/zrb-channel/utils/hash"
)

// Login
// @date: 2022-03-24 10:29:26
func Login(ctx context.Context, conf *Config, mobile, orderNo string) string {
	req := &LoginRequest{
		AppID:          conf.AppId,
		Phone:          mobile,
		ApplicationNum: orderNo,
		Flag:           "0",
	}

	req.Sign = hash.MD5String(req.Flag + req.AppID + req.Phone + req.ApplicationNum + conf.Token)

	return fmt.Sprintf("https://h.kungeek.com/mobile/homePage?noLoginCode=%s*%s*%s*%s*%s",
		req.Flag,
		req.AppID,
		req.Phone,
		req.ApplicationNum,
		req.Sign,
	)
}
