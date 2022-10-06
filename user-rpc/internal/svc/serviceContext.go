package svc

import (
	"cloud-disk/user-rpc/internal/config"
	"cloud-disk/user-rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	Usermodel      model.UserBasicModel
	RepositoryPool model.RepositoryPoolModel
	UserRepository model.UserRepositoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Usermodel:      model.NewUserBasicModel(sqlx.NewMysql(c.DB.DataSource)),
		RepositoryPool: model.NewRepositoryPoolModel(sqlx.NewMysql(c.DB.DataSource)),
		UserRepository: model.NewUserRepositoryModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
