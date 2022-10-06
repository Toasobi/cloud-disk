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

type GetUserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRegisterLogic {
	return &GetUserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRegisterLogic) GetUserRegister(in *pb.GetUserRegisterReq) (*pb.GetUserRegisterResp, error) {
	// todo: add your logic here and delete this line

	//1.接收信息
	name := in.Name
	password := in.Password
	email := in.Email
	code := in.Code

	userIdentity := helper.GenerateUUID()

	//2.校验信息准确性(用户名和邮箱不能重复)

	if count, _ := l.svcCtx.Usermodel.FindByInfo(l.ctx, name, email); count != 0 {
		return &pb.GetUserRegisterResp{
			Status: "",
		}, errors.New("用户已存在或邮箱已被注册")
	}

	//从redis中查找验证码
	result, err := model.RDB.Get(l.ctx, email).Result()

	//验证验证码是否正确
	if result != code {
		return &pb.GetUserRegisterResp{
			Status: "",
		}, errors.New("验证码错误")
	}

	//创建对象
	user := &model.UserBasic{
		Name:      sql.NullString{name, true},
		Identity:  sql.NullString{userIdentity, true},
		Password:  sql.NullString{helper.GetMd5(password), true},
		Email:     sql.NullString{email, true},
		CreatedAt: sql.NullTime{time.Now(), true},
		UpdatedAt: sql.NullTime{time.Now(), true},
	}

	_, err = l.svcCtx.Usermodel.Insert(l.ctx, user)
	if err != nil {
		return &pb.GetUserRegisterResp{
			Status: "",
		}, errors.New("用户数据插入错误")
	} else {
		return &pb.GetUserRegisterResp{
			Status: "用户创建成功",
		}, nil
	}

}
