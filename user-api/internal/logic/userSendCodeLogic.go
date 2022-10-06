package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSendCodeLogic {
	return &UserSendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSendCodeLogic) UserSendCode(req *types.UserSendCodeReq) (resp *types.UserSendCodeResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserSendCode(l.ctx, &pb.GetUserSendCodeReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserSendCodeResp{
		Status: userResp.Status,
	}, nil

}
