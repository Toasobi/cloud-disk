package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileUploadChunkCompleteLogic {
	return &UserFileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileUploadChunkCompleteLogic) UserFileUploadChunkComplete(req *types.UserFileUploadChunkCompleteReq) (resp *types.UserFileUploadChunkCompleteResp, err error) {
	cosOb := make([]*pb.GetUserFileUploadChunkCompleteReq_CosObject, 0)
	for _, v := range req.CosObjects {
		cosOb = append(cosOb, &pb.GetUserFileUploadChunkCompleteReq_CosObject{
			PartNumber: v.PartNumber,
			Etag:       v.Etag,
		})
	}

	userResp, err := l.svcCtx.UserRpcClient.GetUserFileUploadChunkComplete(l.ctx, &pb.GetUserFileUploadChunkCompleteReq{
		Key:        req.Key,
		UploadId:   req.UploadId,
		CosObjects: cosOb,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserFileUploadChunkCompleteResp{
		Status: userResp.Status,
	}, nil
}
