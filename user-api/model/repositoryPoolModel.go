package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RepositoryPoolModel = (*customRepositoryPoolModel)(nil)

type (
	// RepositoryPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepositoryPoolModel.
	RepositoryPoolModel interface {
		repositoryPoolModel
	}

	customRepositoryPoolModel struct {
		*defaultRepositoryPoolModel
	}
)

// NewRepositoryPoolModel returns a model for the database table.
func NewRepositoryPoolModel(conn sqlx.SqlConn) RepositoryPoolModel {
	return &customRepositoryPoolModel{
		defaultRepositoryPoolModel: newRepositoryPoolModel(conn),
	}
}
