package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserLogin(l.ctx, &pb.GetUserLoginReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserLoginResp{
		Token:        userResp.Token,
		RefreshToken: userResp.RefreshToken,
	}, nil
}
