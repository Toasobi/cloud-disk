package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUploadLogic {
	return &GetFileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送文件
func (l *GetFileUploadLogic) GetFileUpload(in *pb.GetFileUploadReq) (*pb.GetFileUploadResp, error) {
	// todo: add your logic here and delete this line
	file := &model.RepositoryPool{
		Identity:  sql.NullString{helper.GenerateUUID(), true},
		Hash:      sql.NullString{in.Hash, true},
		Name:      sql.NullString{in.Name, true},
		Ext:       sql.NullString{in.Ext, true},
		Size:      sql.NullFloat64{float64(in.Size), true},
		Path:      sql.NullString{in.Path, true},
		CreatedAt: sql.NullTime{time.Now(), true},
		UpdatedAt: sql.NullTime{time.Now(), true},
	}

	//将file插入到数据库中
	_, err := l.svcCtx.RepositoryPool.Insert(l.ctx, file)
	if err != nil {
		return nil, errors.New("用户数据插入错误")
	} else {
		return &pb.GetFileUploadResp{
			Identity: file.Identity.String,
			Ext:      file.Ext.String,
			Name:     file.Name.String,
		}, nil
	}

}
