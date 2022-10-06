package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateReq, userIdenity string) (resp *types.UserFolderCreateResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserFolderCreate(l.ctx, &pb.GetUserFolderCreateReq{
		Name:         req.Name,
		ParentId:     int64(req.ParentId),
		UserIdentity: userIdenity,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserFolderCreateResp{
		Status: userResp.Status,
	}, nil

}
