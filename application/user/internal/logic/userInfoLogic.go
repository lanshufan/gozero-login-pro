package logic

import (
	"context"

	"gozero-login-pro/application/user/internal/svc"
	"gozero-login-pro/application/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var reInfo = user.UserInfoResponse{
		UserId:   1,
		UserName: "admin",
		Mobile:   "13800000000",
	}

	//return &user.UserInfoResponse{}, nil
	return &reInfo, nil
}
