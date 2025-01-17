// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asuser"
	"orginone/common/schema/asworkingdata"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsWorkingDataUpdate is the builder for updating AsWorkingData entities.
type AsWorkingDataUpdate struct {
	config
	hooks    []Hook
	mutation *AsWorkingDataMutation
}

// Where appends a list predicates to the AsWorkingDataUpdate builder.
func (awdu *AsWorkingDataUpdate) Where(ps ...predicate.AsWorkingData) *AsWorkingDataUpdate {
	awdu.mutation.Where(ps...)
	return awdu
}

// SetAppID sets the "app_id" field.
func (awdu *AsWorkingDataUpdate) SetAppID(i int64) *AsWorkingDataUpdate {
	awdu.mutation.ResetAppID()
	awdu.mutation.SetAppID(i)
	return awdu
}

// AddAppID adds i to the "app_id" field.
func (awdu *AsWorkingDataUpdate) AddAppID(i int64) *AsWorkingDataUpdate {
	awdu.mutation.AddAppID(i)
	return awdu
}

// SetUserID sets the "user_id" field.
func (awdu *AsWorkingDataUpdate) SetUserID(i int64) *AsWorkingDataUpdate {
	awdu.mutation.SetUserID(i)
	return awdu
}

// SetType sets the "type" field.
func (awdu *AsWorkingDataUpdate) SetType(i int64) *AsWorkingDataUpdate {
	awdu.mutation.ResetType()
	awdu.mutation.SetType(i)
	return awdu
}

// AddType adds i to the "type" field.
func (awdu *AsWorkingDataUpdate) AddType(i int64) *AsWorkingDataUpdate {
	awdu.mutation.AddType(i)
	return awdu
}

// SetIsDeleted sets the "is_deleted" field.
func (awdu *AsWorkingDataUpdate) SetIsDeleted(i int64) *AsWorkingDataUpdate {
	awdu.mutation.ResetIsDeleted()
	awdu.mutation.SetIsDeleted(i)
	return awdu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (awdu *AsWorkingDataUpdate) SetNillableIsDeleted(i *int64) *AsWorkingDataUpdate {
	if i != nil {
		awdu.SetIsDeleted(*i)
	}
	return awdu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (awdu *AsWorkingDataUpdate) AddIsDeleted(i int64) *AsWorkingDataUpdate {
	awdu.mutation.AddIsDeleted(i)
	return awdu
}

// SetStatus sets the "status" field.
func (awdu *AsWorkingDataUpdate) SetStatus(i int64) *AsWorkingDataUpdate {
	awdu.mutation.ResetStatus()
	awdu.mutation.SetStatus(i)
	return awdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (awdu *AsWorkingDataUpdate) SetNillableStatus(i *int64) *AsWorkingDataUpdate {
	if i != nil {
		awdu.SetStatus(*i)
	}
	return awdu
}

// AddStatus adds i to the "status" field.
func (awdu *AsWorkingDataUpdate) AddStatus(i int64) *AsWorkingDataUpdate {
	awdu.mutation.AddStatus(i)
	return awdu
}

// ClearStatus clears the value of the "status" field.
func (awdu *AsWorkingDataUpdate) ClearStatus() *AsWorkingDataUpdate {
	awdu.mutation.ClearStatus()
	return awdu
}

// SetCreateUser sets the "create_user" field.
func (awdu *AsWorkingDataUpdate) SetCreateUser(i int64) *AsWorkingDataUpdate {
	awdu.mutation.ResetCreateUser()
	awdu.mutation.SetCreateUser(i)
	return awdu
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (awdu *AsWorkingDataUpdate) SetNillableCreateUser(i *int64) *AsWorkingDataUpdate {
	if i != nil {
		awdu.SetCreateUser(*i)
	}
	return awdu
}

// AddCreateUser adds i to the "create_user" field.
func (awdu *AsWorkingDataUpdate) AddCreateUser(i int64) *AsWorkingDataUpdate {
	awdu.mutation.AddCreateUser(i)
	return awdu
}

// ClearCreateUser clears the value of the "create_user" field.
func (awdu *AsWorkingDataUpdate) ClearCreateUser() *AsWorkingDataUpdate {
	awdu.mutation.ClearCreateUser()
	return awdu
}

// SetUpdateUser sets the "update_user" field.
func (awdu *AsWorkingDataUpdate) SetUpdateUser(i int64) *AsWorkingDataUpdate {
	awdu.mutation.ResetUpdateUser()
	awdu.mutation.SetUpdateUser(i)
	return awdu
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (awdu *AsWorkingDataUpdate) SetNillableUpdateUser(i *int64) *AsWorkingDataUpdate {
	if i != nil {
		awdu.SetUpdateUser(*i)
	}
	return awdu
}

// AddUpdateUser adds i to the "update_user" field.
func (awdu *AsWorkingDataUpdate) AddUpdateUser(i int64) *AsWorkingDataUpdate {
	awdu.mutation.AddUpdateUser(i)
	return awdu
}

// ClearUpdateUser clears the value of the "update_user" field.
func (awdu *AsWorkingDataUpdate) ClearUpdateUser() *AsWorkingDataUpdate {
	awdu.mutation.ClearUpdateUser()
	return awdu
}

// SetUpdateTime sets the "update_time" field.
func (awdu *AsWorkingDataUpdate) SetUpdateTime(dt date.DateTime) *AsWorkingDataUpdate {
	awdu.mutation.SetUpdateTime(dt)
	return awdu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (awdu *AsWorkingDataUpdate) ClearUpdateTime() *AsWorkingDataUpdate {
	awdu.mutation.ClearUpdateTime()
	return awdu
}

// SetUser sets the "user" edge to the AsUser entity.
func (awdu *AsWorkingDataUpdate) SetUser(a *AsUser) *AsWorkingDataUpdate {
	return awdu.SetUserID(a.ID)
}

// Mutation returns the AsWorkingDataMutation object of the builder.
func (awdu *AsWorkingDataUpdate) Mutation() *AsWorkingDataMutation {
	return awdu.mutation
}

// ClearUser clears the "user" edge to the AsUser entity.
func (awdu *AsWorkingDataUpdate) ClearUser() *AsWorkingDataUpdate {
	awdu.mutation.ClearUser()
	return awdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (awdu *AsWorkingDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	awdu.defaults()
	if len(awdu.hooks) == 0 {
		if err = awdu.check(); err != nil {
			return 0, err
		}
		affected, err = awdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsWorkingDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = awdu.check(); err != nil {
				return 0, err
			}
			awdu.mutation = mutation
			affected, err = awdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(awdu.hooks) - 1; i >= 0; i-- {
			if awdu.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = awdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (awdu *AsWorkingDataUpdate) SaveX(ctx context.Context) int {
	affected, err := awdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (awdu *AsWorkingDataUpdate) Exec(ctx context.Context) error {
	_, err := awdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awdu *AsWorkingDataUpdate) ExecX(ctx context.Context) {
	if err := awdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (awdu *AsWorkingDataUpdate) defaults() {
	if _, ok := awdu.mutation.UpdateTime(); !ok && !awdu.mutation.UpdateTimeCleared() {
		v := asworkingdata.UpdateDefaultUpdateTime()
		awdu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awdu *AsWorkingDataUpdate) check() error {
	if _, ok := awdu.mutation.UserID(); awdu.mutation.UserCleared() && !ok {
		return errors.New(`schema: clearing a required unique edge "AsWorkingData.user"`)
	}
	return nil
}

func (awdu *AsWorkingDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asworkingdata.Table,
			Columns: asworkingdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asworkingdata.FieldID,
			},
		},
	}
	if ps := awdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := awdu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldAppID,
		})
	}
	if value, ok := awdu.mutation.AddedAppID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldAppID,
		})
	}
	if value, ok := awdu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldType,
		})
	}
	if value, ok := awdu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldType,
		})
	}
	if value, ok := awdu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldIsDeleted,
		})
	}
	if value, ok := awdu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldIsDeleted,
		})
	}
	if value, ok := awdu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldStatus,
		})
	}
	if value, ok := awdu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldStatus,
		})
	}
	if awdu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asworkingdata.FieldStatus,
		})
	}
	if value, ok := awdu.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldCreateUser,
		})
	}
	if value, ok := awdu.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldCreateUser,
		})
	}
	if awdu.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asworkingdata.FieldCreateUser,
		})
	}
	if value, ok := awdu.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldUpdateUser,
		})
	}
	if value, ok := awdu.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldUpdateUser,
		})
	}
	if awdu.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asworkingdata.FieldUpdateUser,
		})
	}
	if awdu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asworkingdata.FieldCreateTime,
		})
	}
	if value, ok := awdu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asworkingdata.FieldUpdateTime,
		})
	}
	if awdu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asworkingdata.FieldUpdateTime,
		})
	}
	if awdu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asworkingdata.UserTable,
			Columns: []string{asworkingdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awdu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asworkingdata.UserTable,
			Columns: []string{asworkingdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, awdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asworkingdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsWorkingDataUpdateOne is the builder for updating a single AsWorkingData entity.
type AsWorkingDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsWorkingDataMutation
}

// SetAppID sets the "app_id" field.
func (awduo *AsWorkingDataUpdateOne) SetAppID(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.ResetAppID()
	awduo.mutation.SetAppID(i)
	return awduo
}

// AddAppID adds i to the "app_id" field.
func (awduo *AsWorkingDataUpdateOne) AddAppID(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.AddAppID(i)
	return awduo
}

// SetUserID sets the "user_id" field.
func (awduo *AsWorkingDataUpdateOne) SetUserID(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.SetUserID(i)
	return awduo
}

// SetType sets the "type" field.
func (awduo *AsWorkingDataUpdateOne) SetType(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.ResetType()
	awduo.mutation.SetType(i)
	return awduo
}

// AddType adds i to the "type" field.
func (awduo *AsWorkingDataUpdateOne) AddType(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.AddType(i)
	return awduo
}

// SetIsDeleted sets the "is_deleted" field.
func (awduo *AsWorkingDataUpdateOne) SetIsDeleted(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.ResetIsDeleted()
	awduo.mutation.SetIsDeleted(i)
	return awduo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (awduo *AsWorkingDataUpdateOne) SetNillableIsDeleted(i *int64) *AsWorkingDataUpdateOne {
	if i != nil {
		awduo.SetIsDeleted(*i)
	}
	return awduo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (awduo *AsWorkingDataUpdateOne) AddIsDeleted(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.AddIsDeleted(i)
	return awduo
}

// SetStatus sets the "status" field.
func (awduo *AsWorkingDataUpdateOne) SetStatus(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.ResetStatus()
	awduo.mutation.SetStatus(i)
	return awduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (awduo *AsWorkingDataUpdateOne) SetNillableStatus(i *int64) *AsWorkingDataUpdateOne {
	if i != nil {
		awduo.SetStatus(*i)
	}
	return awduo
}

// AddStatus adds i to the "status" field.
func (awduo *AsWorkingDataUpdateOne) AddStatus(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.AddStatus(i)
	return awduo
}

// ClearStatus clears the value of the "status" field.
func (awduo *AsWorkingDataUpdateOne) ClearStatus() *AsWorkingDataUpdateOne {
	awduo.mutation.ClearStatus()
	return awduo
}

// SetCreateUser sets the "create_user" field.
func (awduo *AsWorkingDataUpdateOne) SetCreateUser(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.ResetCreateUser()
	awduo.mutation.SetCreateUser(i)
	return awduo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (awduo *AsWorkingDataUpdateOne) SetNillableCreateUser(i *int64) *AsWorkingDataUpdateOne {
	if i != nil {
		awduo.SetCreateUser(*i)
	}
	return awduo
}

// AddCreateUser adds i to the "create_user" field.
func (awduo *AsWorkingDataUpdateOne) AddCreateUser(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.AddCreateUser(i)
	return awduo
}

// ClearCreateUser clears the value of the "create_user" field.
func (awduo *AsWorkingDataUpdateOne) ClearCreateUser() *AsWorkingDataUpdateOne {
	awduo.mutation.ClearCreateUser()
	return awduo
}

// SetUpdateUser sets the "update_user" field.
func (awduo *AsWorkingDataUpdateOne) SetUpdateUser(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.ResetUpdateUser()
	awduo.mutation.SetUpdateUser(i)
	return awduo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (awduo *AsWorkingDataUpdateOne) SetNillableUpdateUser(i *int64) *AsWorkingDataUpdateOne {
	if i != nil {
		awduo.SetUpdateUser(*i)
	}
	return awduo
}

// AddUpdateUser adds i to the "update_user" field.
func (awduo *AsWorkingDataUpdateOne) AddUpdateUser(i int64) *AsWorkingDataUpdateOne {
	awduo.mutation.AddUpdateUser(i)
	return awduo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (awduo *AsWorkingDataUpdateOne) ClearUpdateUser() *AsWorkingDataUpdateOne {
	awduo.mutation.ClearUpdateUser()
	return awduo
}

// SetUpdateTime sets the "update_time" field.
func (awduo *AsWorkingDataUpdateOne) SetUpdateTime(dt date.DateTime) *AsWorkingDataUpdateOne {
	awduo.mutation.SetUpdateTime(dt)
	return awduo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (awduo *AsWorkingDataUpdateOne) ClearUpdateTime() *AsWorkingDataUpdateOne {
	awduo.mutation.ClearUpdateTime()
	return awduo
}

// SetUser sets the "user" edge to the AsUser entity.
func (awduo *AsWorkingDataUpdateOne) SetUser(a *AsUser) *AsWorkingDataUpdateOne {
	return awduo.SetUserID(a.ID)
}

// Mutation returns the AsWorkingDataMutation object of the builder.
func (awduo *AsWorkingDataUpdateOne) Mutation() *AsWorkingDataMutation {
	return awduo.mutation
}

// ClearUser clears the "user" edge to the AsUser entity.
func (awduo *AsWorkingDataUpdateOne) ClearUser() *AsWorkingDataUpdateOne {
	awduo.mutation.ClearUser()
	return awduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (awduo *AsWorkingDataUpdateOne) Select(field string, fields ...string) *AsWorkingDataUpdateOne {
	awduo.fields = append([]string{field}, fields...)
	return awduo
}

// Save executes the query and returns the updated AsWorkingData entity.
func (awduo *AsWorkingDataUpdateOne) Save(ctx context.Context) (*AsWorkingData, error) {
	var (
		err  error
		node *AsWorkingData
	)
	awduo.defaults()
	if len(awduo.hooks) == 0 {
		if err = awduo.check(); err != nil {
			return nil, err
		}
		node, err = awduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsWorkingDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = awduo.check(); err != nil {
				return nil, err
			}
			awduo.mutation = mutation
			node, err = awduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(awduo.hooks) - 1; i >= 0; i-- {
			if awduo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = awduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (awduo *AsWorkingDataUpdateOne) SaveX(ctx context.Context) *AsWorkingData {
	node, err := awduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (awduo *AsWorkingDataUpdateOne) Exec(ctx context.Context) error {
	_, err := awduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awduo *AsWorkingDataUpdateOne) ExecX(ctx context.Context) {
	if err := awduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (awduo *AsWorkingDataUpdateOne) defaults() {
	if _, ok := awduo.mutation.UpdateTime(); !ok && !awduo.mutation.UpdateTimeCleared() {
		v := asworkingdata.UpdateDefaultUpdateTime()
		awduo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awduo *AsWorkingDataUpdateOne) check() error {
	if _, ok := awduo.mutation.UserID(); awduo.mutation.UserCleared() && !ok {
		return errors.New(`schema: clearing a required unique edge "AsWorkingData.user"`)
	}
	return nil
}

func (awduo *AsWorkingDataUpdateOne) sqlSave(ctx context.Context) (_node *AsWorkingData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asworkingdata.Table,
			Columns: asworkingdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asworkingdata.FieldID,
			},
		},
	}
	id, ok := awduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsWorkingData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := awduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asworkingdata.FieldID)
		for _, f := range fields {
			if !asworkingdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != asworkingdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := awduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := awduo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldAppID,
		})
	}
	if value, ok := awduo.mutation.AddedAppID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldAppID,
		})
	}
	if value, ok := awduo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldType,
		})
	}
	if value, ok := awduo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldType,
		})
	}
	if value, ok := awduo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldIsDeleted,
		})
	}
	if value, ok := awduo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldIsDeleted,
		})
	}
	if value, ok := awduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldStatus,
		})
	}
	if value, ok := awduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldStatus,
		})
	}
	if awduo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asworkingdata.FieldStatus,
		})
	}
	if value, ok := awduo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldCreateUser,
		})
	}
	if value, ok := awduo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldCreateUser,
		})
	}
	if awduo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asworkingdata.FieldCreateUser,
		})
	}
	if value, ok := awduo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldUpdateUser,
		})
	}
	if value, ok := awduo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asworkingdata.FieldUpdateUser,
		})
	}
	if awduo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asworkingdata.FieldUpdateUser,
		})
	}
	if awduo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asworkingdata.FieldCreateTime,
		})
	}
	if value, ok := awduo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asworkingdata.FieldUpdateTime,
		})
	}
	if awduo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asworkingdata.FieldUpdateTime,
		})
	}
	if awduo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asworkingdata.UserTable,
			Columns: []string{asworkingdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awduo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asworkingdata.UserTable,
			Columns: []string{asworkingdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AsWorkingData{config: awduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, awduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asworkingdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
