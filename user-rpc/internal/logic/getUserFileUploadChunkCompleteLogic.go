package logic

import (
	"context"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/pb"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileUploadChunkCompleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileUploadChunkCompleteLogic {
	return &GetUserFileUploadChunkCompleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分片上传完成
func (l *GetUserFileUploadChunkCompleteLogic) GetUserFileUploadChunkComplete(in *pb.GetUserFileUploadChunkCompleteReq) (*pb.GetUserFileUploadChunkCompleteResp, error) {
	cs := make([]cos.Object, 0)

	for _, v := range in.CosObjects {
		pn := v.PartNumber
		cs = append(cs, cos.Object{
			ETag:       v.Etag,
			PartNumber: int(pn),
		})
	}

	err := helper.PartUploadComplete(in.Key, in.UploadId, cs)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserFileUploadChunkCompleteResp{
		Status: "上传成功",
	}, nil
}
