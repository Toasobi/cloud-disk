package logic

import (
	"context"
	"database/sql"
	"errors"

	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileUpdateLogic {
	return &GetUserFileUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文件
func (l *GetUserFileUpdateLogic) GetUserFileUpdate(in *pb.GetUserFileUpdateReq) (*pb.GetUserFileUpdateResp, error) {
	//先判断是否存在文件
	user_repository := &model.UserRepository{}
	_, err := model.Engine.Table("user_repository").Where("user_identity = ? AND identity = ?", in.UserIdentity, in.Identity).Get(user_repository)
	if err != nil {
		return nil, err
	}
	if user_repository.Identity.String == "" {
		return nil, errors.New("文件不存在")
	}

	//查看统计目录下是否有相同名称的文件
	cnt, err := model.Engine.Table("user_repository").Where("parent_id = ? AND name = ?", user_repository.ParentId, in.Name).Count()
	if cnt > 0 {
		return nil, errors.New("同级目录下有相同名字文件")
	}

	updated := &model.UserRepository{
		Name: sql.NullString{in.Name, true},
	}
	_, err = model.Engine.Table("user_repository").Where("user_identity = ? AND identity = ?", in.UserIdentity, in.Identity).Update(updated)

	if err != nil {
		return nil, err
	}
	return &pb.GetUserFileUpdateResp{
		Name: in.Name,
	}, nil
}
