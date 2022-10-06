package logic

import (
	"context"
	"errors"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserUploadPrepareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserUploadPrepareLogic {
	return &GetUserUploadPrepareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分片上传准备
func (l *GetUserUploadPrepareLogic) GetUserUploadPrepare(in *pb.GetUserUploadPrepareReq) (*pb.GetUserUploadPrepareResp, error) {
	rp := &model.RepositoryPool{}
	has, err := model.Engine.Table("repository_pool").Where("hash = ? ", in.Md5).Get(rp)
	if err != nil {
		return nil, errors.New("数据库发生错误")
	}
	resp := new(pb.GetUserUploadPrepareResp)
	if has {
		//秒传
		resp.Identity = rp.Identity.String
		return resp, nil
	} else {
		//获取该文件的uploadID ,key用来进行文件分片上传
		key, uploadId, err := helper.CosInitUploader(in.Ext)
		if err != nil {
			return nil, err
		}
		resp.Key = key
		resp.UploadId = uploadId
		return resp, nil

	}
}
