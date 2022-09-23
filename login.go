package hdz

import (
	"fmt"
	"github.com/zrb-channel/utils/hash"
)

// Login
// @date: 2022-03-24 10:29:26
func Login(conf *Config, mobile, orderNo string) string {

	req := &LoginRequest{
		AppID:          conf.LoginAppId,
		Phone:          mobile,
		ApplicationNum: orderNo,
		Flag:           "0",
	}

	req.Sign = hash.MD5String(req.Flag + req.AppID + req.Phone + req.ApplicationNum + conf.LoginToken)

	return fmt.Sprintf("https://h.kungeek.com/mobile/homePage?noLoginCode=%s*%s*%s*%s*%s",
		req.Flag,
		req.AppID,
		req.Phone,
		req.ApplicationNum,
		req.Sign,
	)
}
