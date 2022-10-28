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
	asFormDraftFieldNames          = builder.RawFieldNames(&AsFormDraft{})
	asFormDraftRows                = strings.Join(asFormDraftFieldNames, ",")
	asFormDraftRowsExpectAutoSet   = strings.Join(stringx.Remove(asFormDraftFieldNames, "`create_time`", "`update_time`"), ",")
	asFormDraftRowsWithPlaceHolder = strings.Join(stringx.Remove(asFormDraftFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsFormDraftIdPrefix = "cache:asset:asFormDraft:id:"
)

type (
	asFormDraftModel interface {
		Insert(ctx context.Context, data *AsFormDraft) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsFormDraft, error)
		Update(ctx context.Context, data *AsFormDraft) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsFormDraftModel struct {
		sqlc.CachedConn
		table string
	}

	AsFormDraft struct {
		Id                   string         `db:"id"`
		FormModelId          sql.NullString `db:"form_model_id"`
		Committer            sql.NullString `db:"committer"`
		CommitTime           sql.NullTime   `db:"commit_time"`
		PhoneNumber          sql.NullString `db:"phone_number"`
		Agency               sql.NullString `db:"agency"`
		FormInstSheet        sql.NullString `db:"form_inst_sheet"`
		FormInstValue        sql.NullString `db:"form_inst_value"`
		NextNodeIdList       sql.NullString `db:"next_node_id_list"`
		CandidateUserIdLists sql.NullString `db:"candidate_user_id_lists"`
	}
)

func newAsFormDraftModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsFormDraftModel {
	return &defaultAsFormDraftModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_form_draft`",
	}
}

func (m *defaultAsFormDraftModel) Insert(ctx context.Context, data *AsFormDraft) (sql.Result, error) {
	assetAsFormDraftIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormDraftIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, asFormDraftRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.FormModelId, data.Committer, data.CommitTime, data.PhoneNumber, data.Agency, data.FormInstSheet, data.FormInstValue, data.NextNodeIdList, data.CandidateUserIdLists)
	}, assetAsFormDraftIdKey)
	return ret, err
}

func (m *defaultAsFormDraftModel) FindOne(ctx context.Context, id string) (*AsFormDraft, error) {
	assetAsFormDraftIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormDraftIdPrefix, id)
	var resp AsFormDraft
	err := m.QueryRowCtx(ctx, &resp, assetAsFormDraftIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormDraftRows, m.table)
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

func (m *defaultAsFormDraftModel) Update(ctx context.Context, data *AsFormDraft) error {
	assetAsFormDraftIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormDraftIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asFormDraftRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.FormModelId, data.Committer, data.CommitTime, data.PhoneNumber, data.Agency, data.FormInstSheet, data.FormInstValue, data.NextNodeIdList, data.CandidateUserIdLists, data.Id)
	}, assetAsFormDraftIdKey)
	return err
}

func (m *defaultAsFormDraftModel) Delete(ctx context.Context, id string) error {
	assetAsFormDraftIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormDraftIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsFormDraftIdKey)
	return err
}

func (m *defaultAsFormDraftModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsFormDraftIdPrefix, primary)
}

func (m *defaultAsFormDraftModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormDraftRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsFormDraftModel) tableName() string {
	return m.table
}