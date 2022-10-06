package logic

import (
	"context"
	"errors"

	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileDeleteLogic {
	return &GetUserFileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件删除
func (l *GetUserFileDeleteLogic) GetUserFileDelete(in *pb.GetUserFileDeleteReq) (*pb.GetUserFileDeleteResp, error) {
	//查询文件是否存在
	cnt, _ := model.Engine.Table("user_repository").Where("identity = ?", in.Identity).Count()
	if cnt == 0 {
		return &pb.GetUserFileDeleteResp{
			Status: "-1",
		}, errors.New("文件不存在")
	}
	//删除文件
	_, err := model.Engine.Table("user_repository").Where("identity = ?", in.Identity).Delete()
	if err != nil {
		return &pb.GetUserFileDeleteResp{
			Status: "-1",
		}, errors.New("删除文件失败")
	}

	return &pb.GetUserFileDeleteResp{
		Status: "删除文件成功",
	}, nil

}
