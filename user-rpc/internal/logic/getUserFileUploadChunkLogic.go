package logic

import (
	"context"

	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileUploadChunkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileUploadChunkLogic {
	return &GetUserFileUploadChunkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分片上传
func (l *GetUserFileUploadChunkLogic) GetUserFileUploadChunk(in *pb.GetUserFileUploadChunkReq) (*pb.GetUserFileUploadChunkResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserFileUploadChunkResp{}, nil
}
