// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	repositoryPoolFieldNames          = builder.RawFieldNames(&RepositoryPool{})
	repositoryPoolRows                = strings.Join(repositoryPoolFieldNames, ",")
	repositoryPoolRowsExpectAutoSet   = strings.Join(stringx.Remove(repositoryPoolFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	repositoryPoolRowsWithPlaceHolder = strings.Join(stringx.Remove(repositoryPoolFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	repositoryPoolModel interface {
		Insert(ctx context.Context, data *RepositoryPool) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*RepositoryPool, error)
		Update(ctx context.Context, data *RepositoryPool) error
		Delete(ctx context.Context, id uint64) error
		FindOneByHash(ctx context.Context, hash string) (*RepositoryPool, error)
		FindOneByDBIdentity(ctx context.Context, identity string) (*RepositoryPool, error)
	}

	defaultRepositoryPoolModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RepositoryPool struct {
		Id        uint64          `db:"id"`
		Identity  sql.NullString  `db:"identity"`
		Hash      sql.NullString  `db:"hash"` // 文件的唯一标识
		Name      sql.NullString  `db:"name"`
		Ext       sql.NullString  `db:"ext"`  // 文件扩展名
		Size      sql.NullFloat64 `db:"size"` // 文件大小
		Path      sql.NullString  `db:"path"` // 文件路径
		CreatedAt sql.NullTime    `db:"created_at"`
		UpdatedAt sql.NullTime    `db:"updated_at"`
		DeletedAt sql.NullTime    `db:"deleted_at"`
	}
)

func newRepositoryPoolModel(conn sqlx.SqlConn) *defaultRepositoryPoolModel {
	return &defaultRepositoryPoolModel{
		conn:  conn,
		table: "`repository_pool`",
	}
}

func (m *defaultRepositoryPoolModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRepositoryPoolModel) FindOne(ctx context.Context, id uint64) (*RepositoryPool, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", repositoryPoolRows, m.table)
	var resp RepositoryPool
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

func (m *defaultRepositoryPoolModel) FindOneByHash(ctx context.Context, hash string) (*RepositoryPool, error){
	query := fmt.Sprintf("select %s from %s where `hash` = ? limit 1", repositoryPoolRows, m.table)
	var resp RepositoryPool
	err := m.conn.QueryRowCtx(ctx, &resp, query, hash)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultRepositoryPoolModel)FindOneByDBIdentity(ctx context.Context, identity string) (*RepositoryPool, error){
	query := fmt.Sprintf("select %s from %s where `identity` = ? limit 1", repositoryPoolRows, m.table)
	var resp RepositoryPool
	err := m.conn.QueryRowCtx(ctx, &resp, query, identity)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultRepositoryPoolModel) Insert(ctx context.Context, data *RepositoryPool) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, repositoryPoolRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Identity, data.Hash, data.Name, data.Ext, data.Size, data.Path, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return ret, err
}

func (m *defaultRepositoryPoolModel) Update(ctx context.Context, data *RepositoryPool) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, repositoryPoolRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Identity, data.Hash, data.Name, data.Ext, data.Size, data.Path, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	return err
}

func (m *defaultRepositoryPoolModel) tableName() string {
	return m.table
}
