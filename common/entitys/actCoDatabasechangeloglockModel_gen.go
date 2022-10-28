// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	actCoDatabasechangeloglockFieldNames          = builder.RawFieldNames(&ActCoDatabasechangeloglock{})
	actCoDatabasechangeloglockRows                = strings.Join(actCoDatabasechangeloglockFieldNames, ",")
	actCoDatabasechangeloglockRowsExpectAutoSet   = strings.Join(stringx.Remove(actCoDatabasechangeloglockFieldNames, "`create_time`", "`update_time`"), ",")
	actCoDatabasechangeloglockRowsWithPlaceHolder = strings.Join(stringx.Remove(actCoDatabasechangeloglockFieldNames, "`ID`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActCoDatabasechangeloglockIDPrefix = "cache:asset:actCoDatabasechangeloglock:iD:"
)

type (
	actCoDatabasechangeloglockModel interface {
		Insert(ctx context.Context, data *ActCoDatabasechangeloglock) (sql.Result, error)
		FindOne(ctx context.Context, iD int64) (*ActCoDatabasechangeloglock, error)
		Update(ctx context.Context, data *ActCoDatabasechangeloglock) error
		Delete(ctx context.Context, iD int64) error
	}

	defaultActCoDatabasechangeloglockModel struct {
		sqlc.CachedConn
		table string
	}

	ActCoDatabasechangeloglock struct {
		ID          int64          `db:"ID"`
		LOCKED      byte           `db:"LOCKED"`
		LOCKGRANTED sql.NullTime   `db:"LOCKGRANTED"`
		LOCKEDBY    sql.NullString `db:"LOCKEDBY"`
	}
)

func newActCoDatabasechangeloglockModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActCoDatabasechangeloglockModel {
	return &defaultActCoDatabasechangeloglockModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_co_databasechangeloglock`",
	}
}

func (m *defaultActCoDatabasechangeloglockModel) Insert(ctx context.Context, data *ActCoDatabasechangeloglock) (sql.Result, error) {
	assetActCoDatabasechangeloglockIDKey := fmt.Sprintf("%s%v", cacheAssetActCoDatabasechangeloglockIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, actCoDatabasechangeloglockRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.LOCKED, data.LOCKGRANTED, data.LOCKEDBY)
	}, assetActCoDatabasechangeloglockIDKey)
	return ret, err
}

func (m *defaultActCoDatabasechangeloglockModel) FindOne(ctx context.Context, iD int64) (*ActCoDatabasechangeloglock, error) {
	assetActCoDatabasechangeloglockIDKey := fmt.Sprintf("%s%v", cacheAssetActCoDatabasechangeloglockIDPrefix, iD)
	var resp ActCoDatabasechangeloglock
	err := m.QueryRowCtx(ctx, &resp, assetActCoDatabasechangeloglockIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID` = ? limit 1", actCoDatabasechangeloglockRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, iD)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultActCoDatabasechangeloglockModel) Update(ctx context.Context, data *ActCoDatabasechangeloglock) error {
	assetActCoDatabasechangeloglockIDKey := fmt.Sprintf("%s%v", cacheAssetActCoDatabasechangeloglockIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID` = ?", m.table, actCoDatabasechangeloglockRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.LOCKED, data.LOCKGRANTED, data.LOCKEDBY, data.ID)
	}, assetActCoDatabasechangeloglockIDKey)
	return err
}

func (m *defaultActCoDatabasechangeloglockModel) Delete(ctx context.Context, iD int64) error {
	assetActCoDatabasechangeloglockIDKey := fmt.Sprintf("%s%v", cacheAssetActCoDatabasechangeloglockIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActCoDatabasechangeloglockIDKey)
	return err
}

func (m *defaultActCoDatabasechangeloglockModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActCoDatabasechangeloglockIDPrefix, primary)
}

func (m *defaultActCoDatabasechangeloglockModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID` = ? limit 1", actCoDatabasechangeloglockRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActCoDatabasechangeloglockModel) tableName() string {
	return m.table
}