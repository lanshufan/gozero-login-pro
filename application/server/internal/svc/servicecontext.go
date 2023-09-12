package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gozero-login-pro/application/server/internal/config"
	"gozero-login-pro/application/user/user"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: user.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn()),
	}
}
