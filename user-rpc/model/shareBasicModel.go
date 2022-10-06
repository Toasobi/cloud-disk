package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ShareBasicModel = (*customShareBasicModel)(nil)

type (
	// ShareBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShareBasicModel.
	ShareBasicModel interface {
		shareBasicModel
	}

	customShareBasicModel struct {
		*defaultShareBasicModel
	}
)

// NewShareBasicModel returns a model for the database table.
func NewShareBasicModel(conn sqlx.SqlConn) ShareBasicModel {
	return &customShareBasicModel{
		defaultShareBasicModel: newShareBasicModel(conn),
	}
}
