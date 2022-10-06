package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-rpc/internal/svc"
	"cloud-disk/user-rpc/model"
	"cloud-disk/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSendCodeLogic {
	return &GetUserSendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送验证码
func (l *GetUserSendCodeLogic) GetUserSendCode(in *pb.GetUserSendCodeReq) (*pb.GetUserSendCodeResp, error) {
	email := in.Email

	if email == "" {
		return &pb.GetUserSendCodeResp{
			Status: "-1",
		}, errors.New("邮箱不能为空")
	}

	//发送验证码

	code := helper.GenerateCode()
	if err := helper.SendCode(email, code); err != nil {
		return &pb.GetUserSendCodeResp{
			Status: "-1",
		}, errors.New("验证码发送失败")
	}

	//将数据存入redis当中
	_, err := model.RDB.Set(l.ctx, email, code, time.Minute*5).Result()
	if err != nil {
		fmt.Println(err)
		return &pb.GetUserSendCodeResp{
			Status: "-1",
		}, errors.New("redis数据库报错")
	}

	//没有问题后返回ok
	return &pb.GetUserSendCodeResp{
		Status: "验证码发送成功",
	}, nil

}
