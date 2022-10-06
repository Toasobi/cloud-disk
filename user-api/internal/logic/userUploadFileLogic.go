package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUploadFileLogic {
	return &UserUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUploadFileLogic) UserUploadFile(req *types.FileUploadReq) (resp *types.FileUploadResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetFileUpload(l.ctx, &pb.GetFileUploadReq{
		Hash: req.Hash,
		Name: req.Name,
		Ext:  req.Ext,
		Path: req.Path,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &types.FileUploadResp{
		Identity: userResp.Identity,
		Ext:      userResp.Ext,
		Name:     userResp.Name,
	}, nil
}
