package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRepositorySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRepositorySaveLogic {
	return &GetUserRepositorySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关联用户数据库
func (l *GetUserRepositorySaveLogic) GetUserRepositorySave(in *pb.GetUserRepositorySaveReq) (*pb.GetUserRepositorySaveResp, error) {
	result, err := l.svcCtx.RepositoryPool.FindOneByDBIdentity(l.ctx, in.RepositoryIdentity)

	if result == nil {
		return &pb.GetUserRepositorySaveResp{
			Status: "-1",
		}, errors.New("没有查询到该仓库")
	}

	Identity := result.Hash.String

	userSave := &model.UserRepository{
		Identity:           sql.NullString{Identity, true},
		UserIdentity:       sql.NullString{in.UserIdentity, true},
		ParentId:           sql.NullInt64{int64(in.ParentId), true},
		RepositoryIdentity: sql.NullString{in.RepositoryIdentity, true},
		Ext:                sql.NullString{in.Ext, true},
		Name:               sql.NullString{in.Name, true},
		CreatedAt:          sql.NullTime{time.Now(), true},
		UpdatedAt:          sql.NullTime{time.Now(), true},
	}

	//插入数据
	_, err = l.svcCtx.UserRepository.Insert(l.ctx, userSave)
	if err != nil {
		return &pb.GetUserRepositorySaveResp{
			Status: "-1",
		}, errors.New("数据插入失败")
	}

	return &pb.GetUserRepositorySaveResp{
		Status: "关联成功！",
	}, nil

}
