package logic

import (
	"context"

	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserRpcClient.GetUserRegister(l.ctx, &pb.GetUserRegisterReq{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserRegisterResp{
		Status: userResp.Status,
	}, nil
}
