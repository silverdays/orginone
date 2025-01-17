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
	baseinfoAdministrativeAreaAllFieldNames          = builder.RawFieldNames(&BaseinfoAdministrativeAreaAll{})
	baseinfoAdministrativeAreaAllRows                = strings.Join(baseinfoAdministrativeAreaAllFieldNames, ",")
	baseinfoAdministrativeAreaAllRowsExpectAutoSet   = strings.Join(stringx.Remove(baseinfoAdministrativeAreaAllFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	baseinfoAdministrativeAreaAllRowsWithPlaceHolder = strings.Join(stringx.Remove(baseinfoAdministrativeAreaAllFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBaseinfoAdministrativeAreaAllIdPrefix = "cache:asset:baseinfoAdministrativeAreaAll:id:"
)

type (
	baseinfoAdministrativeAreaAllModel interface {
		Insert(ctx context.Context, data *BaseinfoAdministrativeAreaAll) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BaseinfoAdministrativeAreaAll, error)
		Update(ctx context.Context, data *BaseinfoAdministrativeAreaAll) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBaseinfoAdministrativeAreaAllModel struct {
		sqlc.CachedConn
		table string
	}

	BaseinfoAdministrativeAreaAll struct {
		Id         int64          `db:"id"`          // 主键
		Code       sql.NullString `db:"code"`        // 行政区域编码
		Name       sql.NullString `db:"name"`        // 行政区域名称
		Province   sql.NullString `db:"province"`    // 省/直辖市
		City       sql.NullString `db:"city"`        // 市
		Area       sql.NullString `db:"area"`        // 区
		Town       sql.NullString `db:"town"`        // 城镇地区
		CreateUser sql.NullInt64  `db:"create_user"` // 创建人
		CreateTime sql.NullTime   `db:"create_time"` // 创建时间
		UpdateUser sql.NullInt64  `db:"update_user"` // 更新人
		UpdateTime sql.NullTime   `db:"update_time"` // 更新时间
		IsDeleted  sql.NullInt64  `db:"is_deleted"`  // 删除标注
		TsVersion  sql.NullInt64  `db:"ts_version"`  // 乐观锁字段
		Status     sql.NullInt64  `db:"status"`      // 状态
		AllName    sql.NullString `db:"all_name"`    // 全名
		Pid        sql.NullInt64  `db:"pid"`         // 上级行政区域
		Tp         sql.NullInt64  `db:"tp"`          // 行政区域级别
	}
)

func newBaseinfoAdministrativeAreaAllModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBaseinfoAdministrativeAreaAllModel {
	return &defaultBaseinfoAdministrativeAreaAllModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`baseinfo_administrative_area_all`",
	}
}

func (m *defaultBaseinfoAdministrativeAreaAllModel) Insert(ctx context.Context, data *BaseinfoAdministrativeAreaAll) (sql.Result, error) {
	assetBaseinfoAdministrativeAreaAllIdKey := fmt.Sprintf("%s%v", cacheAssetBaseinfoAdministrativeAreaAllIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, baseinfoAdministrativeAreaAllRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Code, data.Name, data.Province, data.City, data.Area, data.Town, data.CreateUser, data.UpdateUser, data.IsDeleted, data.TsVersion, data.Status, data.AllName, data.Pid, data.Tp)
	}, assetBaseinfoAdministrativeAreaAllIdKey)
	return ret, err
}

func (m *defaultBaseinfoAdministrativeAreaAllModel) FindOne(ctx context.Context, id int64) (*BaseinfoAdministrativeAreaAll, error) {
	assetBaseinfoAdministrativeAreaAllIdKey := fmt.Sprintf("%s%v", cacheAssetBaseinfoAdministrativeAreaAllIdPrefix, id)
	var resp BaseinfoAdministrativeAreaAll
	err := m.QueryRowCtx(ctx, &resp, assetBaseinfoAdministrativeAreaAllIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseinfoAdministrativeAreaAllRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultBaseinfoAdministrativeAreaAllModel) Update(ctx context.Context, data *BaseinfoAdministrativeAreaAll) error {
	assetBaseinfoAdministrativeAreaAllIdKey := fmt.Sprintf("%s%v", cacheAssetBaseinfoAdministrativeAreaAllIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, baseinfoAdministrativeAreaAllRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Code, data.Name, data.Province, data.City, data.Area, data.Town, data.CreateUser, data.UpdateUser, data.IsDeleted, data.TsVersion, data.Status, data.AllName, data.Pid, data.Tp, data.Id)
	}, assetBaseinfoAdministrativeAreaAllIdKey)
	return err
}

func (m *defaultBaseinfoAdministrativeAreaAllModel) Delete(ctx context.Context, id int64) error {
	assetBaseinfoAdministrativeAreaAllIdKey := fmt.Sprintf("%s%v", cacheAssetBaseinfoAdministrativeAreaAllIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetBaseinfoAdministrativeAreaAllIdKey)
	return err
}

func (m *defaultBaseinfoAdministrativeAreaAllModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBaseinfoAdministrativeAreaAllIdPrefix, primary)
}

func (m *defaultBaseinfoAdministrativeAreaAllModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", baseinfoAdministrativeAreaAllRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBaseinfoAdministrativeAreaAllModel) tableName() string {
	return m.table
}
