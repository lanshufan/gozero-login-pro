package logic

import (
	"context"
	"fmt"
	"gozero-login-pro/application/user/user"
	"gozero-login-pro/pkg"

	"gozero-login-pro/application/server/internal/svc"
	"gozero-login-pro/application/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	var reLoginResp types.LoginResponse

	res, err := l.svcCtx.UserRpcClient.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		fmt.Println("login rpc failure, ", err)
		l.Logger.Error("login rpc failure, ", err)
		return nil, err
	}

	// 账号密码验证成功，生成token
	token := pkg.GenerateToken(int64(res.UserId), 604800)

	// todo: input token to redis

	reLoginResp.UserId = int64(res.UserId)
	reLoginResp.Token = string(token)
	reLoginResp.UserName = res.UserName

	return &reLoginResp, nil
}
