// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppGroupDistributionQuery is the builder for querying AsMarketAppGroupDistribution entities.
type AsMarketAppGroupDistributionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsMarketAppGroupDistribution
	// eager-loading edges.
	withAppx   *AsMarketAppQuery
	withGroupx *AsAllGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsMarketAppGroupDistributionQuery builder.
func (amagdq *AsMarketAppGroupDistributionQuery) Where(ps ...predicate.AsMarketAppGroupDistribution) *AsMarketAppGroupDistributionQuery {
	amagdq.predicates = append(amagdq.predicates, ps...)
	return amagdq
}

// Limit adds a limit step to the query.
func (amagdq *AsMarketAppGroupDistributionQuery) Limit(limit int) *AsMarketAppGroupDistributionQuery {
	amagdq.limit = &limit
	return amagdq
}

// Offset adds an offset step to the query.
func (amagdq *AsMarketAppGroupDistributionQuery) Offset(offset int) *AsMarketAppGroupDistributionQuery {
	amagdq.offset = &offset
	return amagdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (amagdq *AsMarketAppGroupDistributionQuery) Unique(unique bool) *AsMarketAppGroupDistributionQuery {
	amagdq.unique = &unique
	return amagdq
}

// Order adds an order step to the query.
func (amagdq *AsMarketAppGroupDistributionQuery) Order(o ...OrderFunc) *AsMarketAppGroupDistributionQuery {
	amagdq.order = append(amagdq.order, o...)
	return amagdq
}

// QueryAppx chains the current query on the "appx" edge.
func (amagdq *AsMarketAppGroupDistributionQuery) QueryAppx() *AsMarketAppQuery {
	query := &AsMarketAppQuery{config: amagdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amagdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amagdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asmarketappgroupdistribution.Table, asmarketappgroupdistribution.FieldID, selector),
			sqlgraph.To(asmarketapp.Table, asmarketapp.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asmarketappgroupdistribution.AppxTable, asmarketappgroupdistribution.AppxColumn),
		)
		fromU = sqlgraph.SetNeighbors(amagdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroupx chains the current query on the "groupx" edge.
func (amagdq *AsMarketAppGroupDistributionQuery) QueryGroupx() *AsAllGroupQuery {
	query := &AsAllGroupQuery{config: amagdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amagdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amagdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asmarketappgroupdistribution.Table, asmarketappgroupdistribution.FieldID, selector),
			sqlgraph.To(asallgroup.Table, asallgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asmarketappgroupdistribution.GroupxTable, asmarketappgroupdistribution.GroupxColumn),
		)
		fromU = sqlgraph.SetNeighbors(amagdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsMarketAppGroupDistribution entity from the query.
// Returns a *NotFoundError when no AsMarketAppGroupDistribution was found.
func (amagdq *AsMarketAppGroupDistributionQuery) First(ctx context.Context) (*AsMarketAppGroupDistribution, error) {
	nodes, err := amagdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asmarketappgroupdistribution.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) FirstX(ctx context.Context) *AsMarketAppGroupDistribution {
	node, err := amagdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsMarketAppGroupDistribution ID from the query.
// Returns a *NotFoundError when no AsMarketAppGroupDistribution ID was found.
func (amagdq *AsMarketAppGroupDistributionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = amagdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asmarketappgroupdistribution.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := amagdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsMarketAppGroupDistribution entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsMarketAppGroupDistribution entity is found.
// Returns a *NotFoundError when no AsMarketAppGroupDistribution entities are found.
func (amagdq *AsMarketAppGroupDistributionQuery) Only(ctx context.Context) (*AsMarketAppGroupDistribution, error) {
	nodes, err := amagdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		return nil, &NotSingularError{asmarketappgroupdistribution.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) OnlyX(ctx context.Context) *AsMarketAppGroupDistribution {
	node, err := amagdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsMarketAppGroupDistribution ID in the query.
// Returns a *NotSingularError when more than one AsMarketAppGroupDistribution ID is found.
// Returns a *NotFoundError when no entities are found.
func (amagdq *AsMarketAppGroupDistributionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = amagdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = &NotSingularError{asmarketappgroupdistribution.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := amagdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsMarketAppGroupDistributions.
func (amagdq *AsMarketAppGroupDistributionQuery) All(ctx context.Context) ([]*AsMarketAppGroupDistribution, error) {
	if err := amagdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return amagdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) AllX(ctx context.Context) []*AsMarketAppGroupDistribution {
	nodes, err := amagdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsMarketAppGroupDistribution IDs.
func (amagdq *AsMarketAppGroupDistributionQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := amagdq.Select(asmarketappgroupdistribution.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := amagdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (amagdq *AsMarketAppGroupDistributionQuery) Count(ctx context.Context) (int64, error) {
	if err := amagdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return amagdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) CountX(ctx context.Context) int64 {
	count, err := amagdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (amagdq *AsMarketAppGroupDistributionQuery) Exist(ctx context.Context) (bool, error) {
	if err := amagdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return amagdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (amagdq *AsMarketAppGroupDistributionQuery) ExistX(ctx context.Context) bool {
	exist, err := amagdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsMarketAppGroupDistributionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (amagdq *AsMarketAppGroupDistributionQuery) Clone() *AsMarketAppGroupDistributionQuery {
	if amagdq == nil {
		return nil
	}
	return &AsMarketAppGroupDistributionQuery{
		config:     amagdq.config,
		limit:      amagdq.limit,
		offset:     amagdq.offset,
		order:      append([]OrderFunc{}, amagdq.order...),
		predicates: append([]predicate.AsMarketAppGroupDistribution{}, amagdq.predicates...),
		withAppx:   amagdq.withAppx.Clone(),
		withGroupx: amagdq.withGroupx.Clone(),
		// clone intermediate query.
		sql:    amagdq.sql.Clone(),
		path:   amagdq.path,
		unique: amagdq.unique,
	}
}

// WithAppx tells the query-builder to eager-load the nodes that are connected to
// the "appx" edge. The optional arguments are used to configure the query builder of the edge.
func (amagdq *AsMarketAppGroupDistributionQuery) WithAppx(opts ...func(*AsMarketAppQuery)) *AsMarketAppGroupDistributionQuery {
	query := &AsMarketAppQuery{config: amagdq.config}
	for _, opt := range opts {
		opt(query)
	}
	amagdq.withAppx = query
	return amagdq
}

// WithGroupx tells the query-builder to eager-load the nodes that are connected to
// the "groupx" edge. The optional arguments are used to configure the query builder of the edge.
func (amagdq *AsMarketAppGroupDistributionQuery) WithGroupx(opts ...func(*AsAllGroupQuery)) *AsMarketAppGroupDistributionQuery {
	query := &AsAllGroupQuery{config: amagdq.config}
	for _, opt := range opts {
		opt(query)
	}
	amagdq.withGroupx = query
	return amagdq
}

// ThenAppx tells the query-builder to eager-load the nodes that are connected to
// the "appx" edge. The optional arguments are used to configure the query builder of the edge.
func (amagdq *AsMarketAppGroupDistributionQuery) ThenAppx(opts ...func(*AsMarketAppQuery)) *AsMarketAppGroupDistributionQuery {
	query := &AsMarketAppQuery{config: amagdq.config}
	for _, opt := range opts {
		opt(query.Where(asmarketapp.IsDeleted(0)))
	}
	amagdq.withAppx = query
	return amagdq
}

// ThenGroupx tells the query-builder to eager-load the nodes that are connected to
// the "groupx" edge. The optional arguments are used to configure the query builder of the edge.
func (amagdq *AsMarketAppGroupDistributionQuery) ThenGroupx(opts ...func(*AsAllGroupQuery)) *AsMarketAppGroupDistributionQuery {
	query := &AsAllGroupQuery{config: amagdq.config}
	for _, opt := range opts {
		opt(query.Where(asallgroup.IsDeleted(0)))
	}
	amagdq.withGroupx = query
	return amagdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Remark string `json:"remark"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsMarketAppGroupDistribution.Query().
//		GroupBy(asmarketappgroupdistribution.FieldRemark).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (amagdq *AsMarketAppGroupDistributionQuery) GroupBy(field string, fields ...string) *AsMarketAppGroupDistributionGroupBy {
	group := &AsMarketAppGroupDistributionGroupBy{config: amagdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := amagdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return amagdq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Remark string `json:"remark"`
//	}
//
//	client.AsMarketAppGroupDistribution.Query().
//		Select(asmarketappgroupdistribution.FieldRemark).
//		Scan(ctx, &v)
//
func (amagdq *AsMarketAppGroupDistributionQuery) Select(fields ...string) *AsMarketAppGroupDistributionSelect {
	amagdq.fields = append(amagdq.fields, fields...)
	return &AsMarketAppGroupDistributionSelect{AsMarketAppGroupDistributionQuery: amagdq}
}

func (amagdq *AsMarketAppGroupDistributionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range amagdq.fields {
		if !asmarketappgroupdistribution.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if amagdq.path != nil {
		prev, err := amagdq.path(ctx)
		if err != nil {
			return err
		}
		amagdq.sql = prev
	}
	return nil
}

func (amagdq *AsMarketAppGroupDistributionQuery) sqlAll(ctx context.Context) ([]*AsMarketAppGroupDistribution, error) {
	var (
		nodes       = []*AsMarketAppGroupDistribution{}
		_spec       = amagdq.querySpec()
		loadedTypes = [2]bool{
			amagdq.withAppx != nil,
			amagdq.withGroupx != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsMarketAppGroupDistribution{config: amagdq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("schema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, amagdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := amagdq.withAppx; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsMarketAppGroupDistribution)
		for i := range nodes {
			fk := nodes[i].AppID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asmarketapp.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "app_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Appx = n
			}
		}
	}

	if query := amagdq.withGroupx; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsMarketAppGroupDistribution)
		for i := range nodes {
			fk := nodes[i].GroupID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asallgroup.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Groupx = n
			}
		}
	}

	return nodes, nil
}

func (amagdq *AsMarketAppGroupDistributionQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := amagdq.querySpec()
	_spec.Node.Columns = amagdq.fields
	if len(amagdq.fields) > 0 {
		_spec.Unique = amagdq.unique != nil && *amagdq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, amagdq.driver, _spec)
	return int64(c), err
}

func (amagdq *AsMarketAppGroupDistributionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := amagdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (amagdq *AsMarketAppGroupDistributionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketappgroupdistribution.Table,
			Columns: asmarketappgroupdistribution.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappgroupdistribution.FieldID,
			},
		},
		From:   amagdq.sql,
		Unique: true,
	}
	if unique := amagdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := amagdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asmarketappgroupdistribution.FieldID)
		for i := range fields {
			if fields[i] != asmarketappgroupdistribution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := amagdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := amagdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := amagdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := amagdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (amagdq *AsMarketAppGroupDistributionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(amagdq.driver.Dialect())
	t1 := builder.Table(asmarketappgroupdistribution.Table)
	columns := amagdq.fields
	if len(columns) == 0 {
		columns = asmarketappgroupdistribution.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if amagdq.sql != nil {
		selector = amagdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if amagdq.unique != nil && *amagdq.unique {
		selector.Distinct()
	}
	for _, p := range amagdq.predicates {
		p(selector)
	}
	for _, p := range amagdq.order {
		p(selector)
	}
	if offset := amagdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := amagdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsMarketAppGroupDistributionGroupBy is the group-by builder for AsMarketAppGroupDistribution entities.
type AsMarketAppGroupDistributionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Aggregate(fns ...AggregateFunc) *AsMarketAppGroupDistributionGroupBy {
	amagdgb.fns = append(amagdgb.fns, fns...)
	return amagdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := amagdgb.path(ctx)
	if err != nil {
		return err
	}
	amagdgb.sql = query
	return amagdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := amagdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(amagdgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := amagdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) StringsX(ctx context.Context) []string {
	v, err := amagdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amagdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) StringX(ctx context.Context) string {
	v, err := amagdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(amagdgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := amagdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) IntsX(ctx context.Context) []int {
	v, err := amagdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amagdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) IntX(ctx context.Context) int {
	v, err := amagdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(amagdgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := amagdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := amagdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amagdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := amagdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(amagdgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := amagdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := amagdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amagdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) BoolX(ctx context.Context) bool {
	v, err := amagdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(amagdgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := amagdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := amagdgb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = amagdgb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (amagdgb *AsMarketAppGroupDistributionGroupBy) Int64X(ctx context.Context) int64 {
	v, err := amagdgb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amagdgb *AsMarketAppGroupDistributionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range amagdgb.fields {
		if !asmarketappgroupdistribution.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := amagdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := amagdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (amagdgb *AsMarketAppGroupDistributionGroupBy) sqlQuery() *sql.Selector {
	selector := amagdgb.sql.Select()
	aggregation := make([]string, 0, len(amagdgb.fns))
	for _, fn := range amagdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(amagdgb.fields)+len(amagdgb.fns))
		for _, f := range amagdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(amagdgb.fields...)...)
}

// AsMarketAppGroupDistributionSelect is the builder for selecting fields of AsMarketAppGroupDistribution entities.
type AsMarketAppGroupDistributionSelect struct {
	*AsMarketAppGroupDistributionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (amagds *AsMarketAppGroupDistributionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := amagds.prepareQuery(ctx); err != nil {
		return err
	}
	amagds.sql = amagds.AsMarketAppGroupDistributionQuery.sqlQuery(ctx)
	return amagds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := amagds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(amagds.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := amagds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) StringsX(ctx context.Context) []string {
	v, err := amagds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amagds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) StringX(ctx context.Context) string {
	v, err := amagds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(amagds.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := amagds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) IntsX(ctx context.Context) []int {
	v, err := amagds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amagds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) IntX(ctx context.Context) int {
	v, err := amagds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(amagds.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := amagds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := amagds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amagds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) Float64X(ctx context.Context) float64 {
	v, err := amagds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(amagds.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := amagds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) BoolsX(ctx context.Context) []bool {
	v, err := amagds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amagds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) BoolX(ctx context.Context) bool {
	v, err := amagds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(amagds.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppGroupDistributionSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := amagds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) Int64sX(ctx context.Context) []int64 {
	v, err := amagds.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (amagds *AsMarketAppGroupDistributionSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = amagds.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappgroupdistribution.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppGroupDistributionSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (amagds *AsMarketAppGroupDistributionSelect) Int64X(ctx context.Context) int64 {
	v, err := amagds.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amagds *AsMarketAppGroupDistributionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := amagds.sql.Query()
	if err := amagds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
