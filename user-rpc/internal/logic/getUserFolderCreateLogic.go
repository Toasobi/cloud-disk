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

type GetUserFolderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFolderCreateLogic {
	return &GetUserFolderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建文件夹
func (l *GetUserFolderCreateLogic) GetUserFolderCreate(in *pb.GetUserFolderCreateReq) (*pb.GetUserFolderCreateResp, error) {
	//首先判断同级目录下是否有相同名称的文件夹
	cnt, err := model.Engine.Table("user_repository").Where("parent_id = ? AND name = ?", in.ParentId, in.Name).Count()
	if cnt > 0 {
		return nil, errors.New("同级目录下有相同名字的文件夹")
	}

	folder := model.UserRepository{
		Identity:     sql.NullString{helper.GenerateUUID(), true},
		UserIdentity: sql.NullString{in.UserIdentity, true},
		ParentId:     sql.NullInt64{int64(in.ParentId), true},
		Name:         sql.NullString{in.Name, true},
		CreatedAt:    sql.NullTime{time.Now(), true},
		UpdatedAt:    sql.NullTime{time.Now(), true},
	}

	_, err = model.Engine.Table("user_repository").Insert(&folder)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserFolderCreateResp{
		Status: "文件夹创建成功",
	}, nil
}
