package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListReq, userIdentity string) (resp *types.UserFileListResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserFileList(l.ctx, &pb.GetUserFileListReq{
		Id:           req.Id,
		Page:         req.Page,
		Size:         req.Size,
		UserIdentity: userIdentity,
	})
	if err != nil {
		return nil, err
	}

	userFile := make([]*types.UserFile, 0)
	for _, v := range userResp.List {
		userFile = append(userFile, &types.UserFile{
			Id:                 v.Id,
			Identity:           v.Identity,
			RepositoryIdentity: v.RepositoryIdentity,
			Ext:                v.Ext,
			Name:               v.Name,
			Path:               v.Path,
			Size:               v.Size,
		})
	}

	return &types.UserFileListResp{
		List:  userFile,
		Count: userResp.Count,
	}, nil
}
