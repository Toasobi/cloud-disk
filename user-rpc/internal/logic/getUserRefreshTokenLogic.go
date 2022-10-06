package logic

import (
	"context"
	"errors"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRefreshTokenLogic {
	return &GetUserRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 刷新token
func (l *GetUserRefreshTokenLogic) GetUserRefreshToken(in *pb.GetUserRefreshTokenReq) (*pb.GetUserRefreshTokenResp, error) {
	uc, err := helper.AnalyseToken(in.Authorization)

	if err != nil {
		return nil, errors.New("解析token失败")
	}

	//生成token
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, 20)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, 3600)

	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &pb.GetUserRefreshTokenResp{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
