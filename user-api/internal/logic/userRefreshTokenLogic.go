package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRefreshTokenLogic {
	return &UserRefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRefreshTokenLogic) UserRefreshToken(req *types.UserRefreshTokenReq, authorization string) (resp *types.UserRefreshTokenResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserRefreshToken(l.ctx, &pb.GetUserRefreshTokenReq{
		Authorization: authorization,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserRefreshTokenResp{
		Token:        userResp.Token,
		RefreshToken: userResp.RefreshToken,
	}, nil

}
