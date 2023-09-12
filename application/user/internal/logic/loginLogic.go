package logic

import (
	"context"
	"errors"
	"gozero-login-pro/application/user/internal/svc"
	"gozero-login-pro/application/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	l.Logger.WithContext(l.ctx).Info("调用记录")

	mobile := in.GetMobile()
	password := in.GetPassword()
	if mobile != "13800000000" || password != "redhat" {
		return nil, errors.New("账号密码不正确")
	}

	return &user.LoginResponse{
		UserId:   1,
		Mobile:   mobile,
		UserName: "admin",
	}, nil
}
