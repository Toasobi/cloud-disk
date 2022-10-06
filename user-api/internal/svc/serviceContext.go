package svc

import (
	"cloud-disk/user-api/internal/config"
	"cloud-disk/user-rpc/usercenter"

	"cloud-disk/user-api/model"

	"cloud-disk/user-api/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Usermodel      model.UserBasicModel
	RepositoryPool model.RepositoryPoolModel
	UserRepository model.UserRepositoryModel
	Auth           rest.Middleware

	UserRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Usermodel:      model.NewUserBasicModel(sqlx.NewMysql(c.DB.DataSource)),
		RepositoryPool: model.NewRepositoryPoolModel(sqlx.NewMysql(c.DB.DataSource)),
		UserRepository: model.NewUserRepositoryModel(sqlx.NewMysql(c.DB.DataSource)),
		Auth:           middleware.NewAuthMiddleware().Handle,

		UserRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
