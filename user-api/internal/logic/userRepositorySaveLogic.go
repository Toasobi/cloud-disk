package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveReq, userIdentity string) (resp *types.UserRepositorySaveResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserRepositorySave(l.ctx, &pb.GetUserRepositorySaveReq{
		ParentId:           int64(req.ParentId),
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
		UserIdentity:       userIdentity,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserRepositorySaveResp{
		Status: userResp.Status,
	}, nil
}
