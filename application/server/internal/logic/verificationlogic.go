package logic

import (
	"context"

	"gozero-login-pro/application/server/internal/svc"
	"gozero-login-pro/application/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.WithContext(l.ctx).Info("测试日志输出")
	return nil, nil
}
