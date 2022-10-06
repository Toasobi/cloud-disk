package logic

import (
	"context"
	"errors"
	"fmt"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLoginLogic {
	return &GetUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *GetUserLoginLogic) GetUserLogin(in *pb.GetUserLoginReq) (*pb.GetUserLoginResp, error) {
	// todo: add your logic here and delete this line
	//1.查询是否有此人
	user, err := l.svcCtx.Usermodel.FindOneByName(l.ctx, in.Username)

	if err != nil && err != model.ErrNotFound {

		fmt.Println(err)

		return nil, errors.New("查询数据失败")

	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	//2.签发token
	if user.Password.String == helper.GetMd5(in.Password) {

		token, err := helper.GenerateToken(int(user.Id), user.Identity.String, user.Name.String, 20)
		if err != nil {
			return nil, errors.New("用户签发token失败")
		}

		refreshToken, err := helper.GenerateToken(int(user.Id), user.Identity.String, user.Name.String, 3600)
		if err != nil {
			return nil, errors.New("用户签发token失败")
		}
		return &pb.GetUserLoginResp{
			Token:        token,
			RefreshToken: refreshToken,
		}, nil

	}

	return nil, errors.New("未知的错误")

}
