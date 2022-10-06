package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteReq, userIdentity string) (resp *types.UserFileDeleteResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserFileDelete(l.ctx, &pb.GetUserFileDeleteReq{
		Identity:     req.Identity,
		UserIdentity: userIdentity,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserFileDeleteResp{
		Status: userResp.Status,
	}, nil

}
