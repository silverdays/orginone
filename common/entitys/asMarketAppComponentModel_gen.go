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
	asMarketAppComponentFieldNames          = builder.RawFieldNames(&AsMarketAppComponent{})
	asMarketAppComponentRows                = strings.Join(asMarketAppComponentFieldNames, ",")
	asMarketAppComponentRowsExpectAutoSet   = strings.Join(stringx.Remove(asMarketAppComponentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asMarketAppComponentRowsWithPlaceHolder = strings.Join(stringx.Remove(asMarketAppComponentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsMarketAppComponentIdPrefix = "cache:asset:asMarketAppComponent:id:"
)

type (
	asMarketAppComponentModel interface {
		Insert(ctx context.Context, data *AsMarketAppComponent) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsMarketAppComponent, error)
		Update(ctx context.Context, data *AsMarketAppComponent) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsMarketAppComponentModel struct {
		sqlc.CachedConn
		table string
	}

	AsMarketAppComponent struct {
		Id           int64          `db:"id"`            // 主键
		AppId        sql.NullInt64  `db:"app_id"`        // app主键
		Code         sql.NullString `db:"code"`          // 编码
		Name         sql.NullString `db:"name"`          // 名称
		Url          sql.NullString `db:"url"`           // url
		Tp           sql.NullInt64  `db:"tp"`            // 类型
		PreviewPic   sql.NullString `db:"preview_pic"`   // 预览图片
		LayoutType   sql.NullString `db:"layout_type"`   // 布局类型
		LayoutConfig sql.NullString `db:"layout_config"` // 布局配置
		TenantCode   sql.NullString `db:"tenant_code"`   // 租户编码
		Source       sql.NullString `db:"source"`        // 平台还剩应用
		CreateUser   sql.NullInt64  `db:"create_user"`   // 创建人
		CreateTime   sql.NullTime   `db:"create_time"`   // 创建时间
		UpdateUser   sql.NullInt64  `db:"update_user"`   // 修改人
		UpdateTime   sql.NullTime   `db:"update_time"`   // 修改时间
		Status       sql.NullInt64  `db:"status"`        // 状态
		IsDeleted    sql.NullInt64  `db:"is_deleted"`    // 是否已删除
	}
)

func newAsMarketAppComponentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsMarketAppComponentModel {
	return &defaultAsMarketAppComponentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_market_app_component`",
	}
}

func (m *defaultAsMarketAppComponentModel) Insert(ctx context.Context, data *AsMarketAppComponent) (sql.Result, error) {
	assetAsMarketAppComponentIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppComponentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, asMarketAppComponentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AppId, data.Code, data.Name, data.Url, data.Tp, data.PreviewPic, data.LayoutType, data.LayoutConfig, data.TenantCode, data.Source, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted)
	}, assetAsMarketAppComponentIdKey)
	return ret, err
}

func (m *defaultAsMarketAppComponentModel) FindOne(ctx context.Context, id int64) (*AsMarketAppComponent, error) {
	assetAsMarketAppComponentIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppComponentIdPrefix, id)
	var resp AsMarketAppComponent
	err := m.QueryRowCtx(ctx, &resp, assetAsMarketAppComponentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asMarketAppComponentRows, m.table)
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

func (m *defaultAsMarketAppComponentModel) Update(ctx context.Context, data *AsMarketAppComponent) error {
	assetAsMarketAppComponentIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppComponentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asMarketAppComponentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AppId, data.Code, data.Name, data.Url, data.Tp, data.PreviewPic, data.LayoutType, data.LayoutConfig, data.TenantCode, data.Source, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.Id)
	}, assetAsMarketAppComponentIdKey)
	return err
}

func (m *defaultAsMarketAppComponentModel) Delete(ctx context.Context, id int64) error {
	assetAsMarketAppComponentIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppComponentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsMarketAppComponentIdKey)
	return err
}

func (m *defaultAsMarketAppComponentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsMarketAppComponentIdPrefix, primary)
}

func (m *defaultAsMarketAppComponentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asMarketAppComponentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsMarketAppComponentModel) tableName() string {
	return m.table
}
