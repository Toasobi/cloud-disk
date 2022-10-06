package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileUpdateLogic {
	return &UserFileUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileUpdateLogic) UserFileUpdate(req *types.UserFileUpdateReq, userIdentity string) (resp *types.UserFileUpdateResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserFileUpdate(l.ctx, &pb.GetUserFileUpdateReq{
		Identity:     req.Identity,
		Name:         req.Name,
		UserIdentity: userIdentity,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserFileUpdateResp{
		Name: userResp.Name,
	}, nil

}
