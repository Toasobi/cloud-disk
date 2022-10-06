package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUploadPrepareLogic {
	return &UserUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUploadPrepareLogic) UserUploadPrepare(req *types.UserUploadPrepareReq) (resp *types.UserUploadPrepareResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserUploadPrepare(l.ctx, &pb.GetUserUploadPrepareReq{
		Md5:  req.Md5,
		Name: req.Name,
		Ext:  req.Ext,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserUploadPrepareResp{
		Identity: userResp.Identity,
		UploadId: userResp.UploadId,
		Key:      userResp.Key,
	}, nil
}
