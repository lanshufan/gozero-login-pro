package logic

import (
	"context"
	"fmt"
	"gozero-login-pro/application/user/user"

	"gozero-login-pro/application/server/internal/svc"
	"gozero-login-pro/application/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// get userId from token
	//userIdVal := l.ctx.Value("userId")
	//fmt.Println("userId: ", userIdVal)

	var reUserInfo types.UserInfoResponse
	userInfo, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{})
	if err != nil {
		fmt.Println("rpc failure, ", err)
		l.Logger.Info(err)
		return nil, err
	}

	reUserInfo.UserId = int64(userInfo.UserId)
	reUserInfo.UserName = userInfo.UserName
	reUserInfo.Mobile = userInfo.Mobile

	// todo: parse token and return user info

	return &reUserInfo, nil
}
