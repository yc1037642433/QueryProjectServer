// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAdminFieldNames          = builder.RawFieldNames(&UserAdmin{})
	userAdminRows                = strings.Join(userAdminFieldNames, ",")
	userAdminRowsExpectAutoSet   = strings.Join(stringx.Remove(userAdminFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAdminRowsWithPlaceHolder = strings.Join(stringx.Remove(userAdminFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userAdminModel interface {
		Insert(ctx context.Context, data *UserAdmin) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAdmin, error)
		FindOneByUsername(ctx context.Context, username string) (*UserAdmin, error)
		Update(ctx context.Context, data *UserAdmin) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAdminModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserAdmin struct {
		Id         int64     `db:"id"`
		Username   string    `db:"username"`    // 登录用户名
		Passwd     string    `db:"passwd"`      // 登录用户密码
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func newUserAdminModel(conn sqlx.SqlConn) *defaultUserAdminModel {
	return &defaultUserAdminModel{
		conn:  conn,
		table: "`user_admin`",
	}
}

func (m *defaultUserAdminModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserAdminModel) FindOne(ctx context.Context, id int64) (*UserAdmin, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAdminRows, m.table)
	var resp UserAdmin
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAdminModel) FindOneByUsername(ctx context.Context, username string) (*UserAdmin, error) {
	var resp UserAdmin
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userAdminRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAdminModel) Insert(ctx context.Context, data *UserAdmin) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userAdminRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Passwd)
	return ret, err
}

func (m *defaultUserAdminModel) Update(ctx context.Context, newData *UserAdmin) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAdminRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Username, newData.Passwd, newData.Id)
	return err
}

func (m *defaultUserAdminModel) tableName() string {
	return m.table
}
