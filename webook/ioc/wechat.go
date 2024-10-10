package ioc

import (
	"github.com/ClearloveHn/webook/webook/internal/service/oauth2/wechat"
)

func InitWechatService() wechat.Service {
	return wechat.NewService("test", "test")
}
