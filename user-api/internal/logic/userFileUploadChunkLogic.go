package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileUploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileUploadChunkLogic {
	return &UserFileUploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileUploadChunkLogic) UserFileUploadChunk(req *types.UserFileUploadChunkReq) (resp *types.UserFileUploadChunkResp, err error) {
	// todo: add your logic here and delete this line
	return
}
