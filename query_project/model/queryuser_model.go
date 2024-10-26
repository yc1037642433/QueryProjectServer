package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ QueryuserModel = (*customQueryuserModel)(nil)

type (
	// QueryuserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQueryuserModel.
	QueryuserModel interface {
		queryuserModel
	}

	customQueryuserModel struct {
		*defaultQueryuserModel
	}
)

// NewQueryuserModel returns a model for the database table.
func NewQueryuserModel(conn sqlx.SqlConn) QueryuserModel {
	return &customQueryuserModel{
		defaultQueryuserModel: newQueryuserModel(conn),
	}
}
