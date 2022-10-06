package logic

import (
	"context"
	"fmt"
	"time"

	"cloud-disk/user-api/define"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/internal/types"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileListLogic {
	return &GetUserFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 列出文件情况
func (l *GetUserFileListLogic) GetUserFileList(in *pb.GetUserFileListReq) (*pb.GetUserFileListResp, error) {
	uf := make([]*types.UserFile, 0)

	size := in.Size
	if size == 0 {
		size = int64(define.PageSize)
	}

	page := in.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	//查询用户文件列表
	err := model.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", in.Id, in.UserIdentity).
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,user_repository.ext,"+
			"user_repository.name,repository_pool.size,repository_pool.path").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format("2006-01-02 15:04:05")).Limit(int(size), int(offset)).Find(&uf)

	if err != nil {
		fmt.Println("数据库查找失败")
		return nil, err
	}

	cnt, err := model.Engine.Where("parent_id = ? AND user_identity = ?", in.Id, in.UserIdentity).Count(new(model.UserRepository))
	if err != nil {
		fmt.Println("我超最后一步都能错")
		return nil, err
	}

	pf := make([]*pb.GetUserFileListResp_UserFile, 0)
	for _, v := range uf {
		pf = append(pf, &pb.GetUserFileListResp_UserFile{
			Id:                 v.Id,
			Ext:                v.Ext,
			Identity:           v.Identity,
			RepositoryIdentity: v.RepositoryIdentity,
			Name:               v.Name,
			Path:               v.Path,
			Size:               v.Size,
		})
	}

	return &pb.GetUserFileListResp{
		List:  pf,
		Count: cnt,
	}, nil

}
