// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	actFoDatabasechangelogFieldNames          = builder.RawFieldNames(&ActFoDatabasechangelog{})
	actFoDatabasechangelogRows                = strings.Join(actFoDatabasechangelogFieldNames, ",")
	actFoDatabasechangelogRowsExpectAutoSet   = strings.Join(stringx.Remove(actFoDatabasechangelogFieldNames, "`create_time`", "`update_time`"), ",")
	actFoDatabasechangelogRowsWithPlaceHolder = strings.Join(stringx.Remove(actFoDatabasechangelogFieldNames, "`ID`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActFoDatabasechangelogIDPrefix = "cache:asset:actFoDatabasechangelog:iD:"
)

type (
	actFoDatabasechangelogModel interface {
		Insert(ctx context.Context, data *ActFoDatabasechangelog) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActFoDatabasechangelog, error)
		Update(ctx context.Context, data *ActFoDatabasechangelog) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActFoDatabasechangelogModel struct {
		sqlc.CachedConn
		table string
	}

	ActFoDatabasechangelog struct {
		ID            string         `db:"ID"`
		AUTHOR        string         `db:"AUTHOR"`
		FILENAME      string         `db:"FILENAME"`
		DATEEXECUTED  time.Time      `db:"DATEEXECUTED"`
		ORDEREXECUTED int64          `db:"ORDEREXECUTED"`
		EXECTYPE      string         `db:"EXECTYPE"`
		MD5SUM        sql.NullString `db:"MD5SUM"`
		DESCRIPTION   sql.NullString `db:"DESCRIPTION"`
		COMMENTS      sql.NullString `db:"COMMENTS"`
		TAG           sql.NullString `db:"TAG"`
		LIQUIBASE     sql.NullString `db:"LIQUIBASE"`
		CONTEXTS      sql.NullString `db:"CONTEXTS"`
		LABELS        sql.NullString `db:"LABELS"`
		DEPLOYMENTID  sql.NullString `db:"DEPLOYMENT_ID"`
	}
)

func newActFoDatabasechangelogModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActFoDatabasechangelogModel {
	return &defaultActFoDatabasechangelogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_fo_databasechangelog`",
	}
}

func (m *defaultActFoDatabasechangelogModel) Insert(ctx context.Context, data *ActFoDatabasechangelog) (sql.Result, error) {
	assetActFoDatabasechangelogIDKey := fmt.Sprintf("%s%v", cacheAssetActFoDatabasechangelogIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actFoDatabasechangelogRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.AUTHOR, data.FILENAME, data.DATEEXECUTED, data.ORDEREXECUTED, data.EXECTYPE, data.MD5SUM, data.DESCRIPTION, data.COMMENTS, data.TAG, data.LIQUIBASE, data.CONTEXTS, data.LABELS, data.DEPLOYMENTID)
	}, assetActFoDatabasechangelogIDKey)
	return ret, err
}

func (m *defaultActFoDatabasechangelogModel) FindOne(ctx context.Context, iD string) (*ActFoDatabasechangelog, error) {
	assetActFoDatabasechangelogIDKey := fmt.Sprintf("%s%v", cacheAssetActFoDatabasechangelogIDPrefix, iD)
	var resp ActFoDatabasechangelog
	err := m.QueryRowCtx(ctx, &resp, assetActFoDatabasechangelogIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID` = ? limit 1", actFoDatabasechangelogRows, m.table)
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

func (m *defaultActFoDatabasechangelogModel) Update(ctx context.Context, data *ActFoDatabasechangelog) error {
	assetActFoDatabasechangelogIDKey := fmt.Sprintf("%s%v", cacheAssetActFoDatabasechangelogIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID` = ?", m.table, actFoDatabasechangelogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AUTHOR, data.FILENAME, data.DATEEXECUTED, data.ORDEREXECUTED, data.EXECTYPE, data.MD5SUM, data.DESCRIPTION, data.COMMENTS, data.TAG, data.LIQUIBASE, data.CONTEXTS, data.LABELS, data.DEPLOYMENTID, data.ID)
	}, assetActFoDatabasechangelogIDKey)
	return err
}

func (m *defaultActFoDatabasechangelogModel) Delete(ctx context.Context, iD string) error {
	assetActFoDatabasechangelogIDKey := fmt.Sprintf("%s%v", cacheAssetActFoDatabasechangelogIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActFoDatabasechangelogIDKey)
	return err
}

func (m *defaultActFoDatabasechangelogModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActFoDatabasechangelogIDPrefix, primary)
}

func (m *defaultActFoDatabasechangelogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID` = ? limit 1", actFoDatabasechangelogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActFoDatabasechangelogModel) tableName() string {
	return m.table
}