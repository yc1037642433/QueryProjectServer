package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserAdminModel = (*customUserAdminModel)(nil)

type (
	// UserAdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAdminModel.
	UserAdminModel interface {
		userAdminModel
	}

	customUserAdminModel struct {
		*defaultUserAdminModel
	}
)

// NewUserAdminModel returns a model for the database table.
func NewUserAdminModel(conn sqlx.SqlConn) UserAdminModel {
	return &customUserAdminModel{
		defaultUserAdminModel: newUserAdminModel(conn),
	}
}
