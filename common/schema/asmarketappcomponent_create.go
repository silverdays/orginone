// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappcomponent"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppComponentCreate is the builder for creating a AsMarketAppComponent entity.
type AsMarketAppComponentCreate struct {
	config
	mutation *AsMarketAppComponentMutation
	hooks    []Hook
}

// SetAppID sets the "app_id" field.
func (amacc *AsMarketAppComponentCreate) SetAppID(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetAppID(i)
	return amacc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableAppID(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetAppID(*i)
	}
	return amacc
}

// SetCode sets the "code" field.
func (amacc *AsMarketAppComponentCreate) SetCode(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetCode(s)
	return amacc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableCode(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetCode(*s)
	}
	return amacc
}

// SetName sets the "name" field.
func (amacc *AsMarketAppComponentCreate) SetName(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetName(s)
	return amacc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableName(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetName(*s)
	}
	return amacc
}

// SetURL sets the "url" field.
func (amacc *AsMarketAppComponentCreate) SetURL(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetURL(s)
	return amacc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableURL(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetURL(*s)
	}
	return amacc
}

// SetType sets the "type" field.
func (amacc *AsMarketAppComponentCreate) SetType(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetType(i)
	return amacc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableType(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetType(*i)
	}
	return amacc
}

// SetPreviewPic sets the "preview_pic" field.
func (amacc *AsMarketAppComponentCreate) SetPreviewPic(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetPreviewPic(s)
	return amacc
}

// SetNillablePreviewPic sets the "preview_pic" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillablePreviewPic(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetPreviewPic(*s)
	}
	return amacc
}

// SetLayoutType sets the "layout_type" field.
func (amacc *AsMarketAppComponentCreate) SetLayoutType(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetLayoutType(s)
	return amacc
}

// SetNillableLayoutType sets the "layout_type" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableLayoutType(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetLayoutType(*s)
	}
	return amacc
}

// SetLayoutConfig sets the "layout_config" field.
func (amacc *AsMarketAppComponentCreate) SetLayoutConfig(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetLayoutConfig(s)
	return amacc
}

// SetNillableLayoutConfig sets the "layout_config" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableLayoutConfig(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetLayoutConfig(*s)
	}
	return amacc
}

// SetTenantCode sets the "tenant_code" field.
func (amacc *AsMarketAppComponentCreate) SetTenantCode(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetTenantCode(s)
	return amacc
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableTenantCode(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetTenantCode(*s)
	}
	return amacc
}

// SetSource sets the "source" field.
func (amacc *AsMarketAppComponentCreate) SetSource(s string) *AsMarketAppComponentCreate {
	amacc.mutation.SetSource(s)
	return amacc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableSource(s *string) *AsMarketAppComponentCreate {
	if s != nil {
		amacc.SetSource(*s)
	}
	return amacc
}

// SetIsDeleted sets the "is_deleted" field.
func (amacc *AsMarketAppComponentCreate) SetIsDeleted(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetIsDeleted(i)
	return amacc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableIsDeleted(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetIsDeleted(*i)
	}
	return amacc
}

// SetStatus sets the "status" field.
func (amacc *AsMarketAppComponentCreate) SetStatus(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetStatus(i)
	return amacc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableStatus(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetStatus(*i)
	}
	return amacc
}

// SetCreateUser sets the "create_user" field.
func (amacc *AsMarketAppComponentCreate) SetCreateUser(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetCreateUser(i)
	return amacc
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableCreateUser(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetCreateUser(*i)
	}
	return amacc
}

// SetUpdateUser sets the "update_user" field.
func (amacc *AsMarketAppComponentCreate) SetUpdateUser(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetUpdateUser(i)
	return amacc
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableUpdateUser(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetUpdateUser(*i)
	}
	return amacc
}

// SetCreateTime sets the "create_time" field.
func (amacc *AsMarketAppComponentCreate) SetCreateTime(dt date.DateTime) *AsMarketAppComponentCreate {
	amacc.mutation.SetCreateTime(dt)
	return amacc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableCreateTime(dt *date.DateTime) *AsMarketAppComponentCreate {
	if dt != nil {
		amacc.SetCreateTime(*dt)
	}
	return amacc
}

// SetUpdateTime sets the "update_time" field.
func (amacc *AsMarketAppComponentCreate) SetUpdateTime(dt date.DateTime) *AsMarketAppComponentCreate {
	amacc.mutation.SetUpdateTime(dt)
	return amacc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableUpdateTime(dt *date.DateTime) *AsMarketAppComponentCreate {
	if dt != nil {
		amacc.SetUpdateTime(*dt)
	}
	return amacc
}

// SetID sets the "id" field.
func (amacc *AsMarketAppComponentCreate) SetID(i int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetID(i)
	return amacc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableID(i *int64) *AsMarketAppComponentCreate {
	if i != nil {
		amacc.SetID(*i)
	}
	return amacc
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (amacc *AsMarketAppComponentCreate) SetAppxID(id int64) *AsMarketAppComponentCreate {
	amacc.mutation.SetAppxID(id)
	return amacc
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (amacc *AsMarketAppComponentCreate) SetNillableAppxID(id *int64) *AsMarketAppComponentCreate {
	if id != nil {
		amacc = amacc.SetAppxID(*id)
	}
	return amacc
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (amacc *AsMarketAppComponentCreate) SetAppx(a *AsMarketApp) *AsMarketAppComponentCreate {
	return amacc.SetAppxID(a.ID)
}

// Mutation returns the AsMarketAppComponentMutation object of the builder.
func (amacc *AsMarketAppComponentCreate) Mutation() *AsMarketAppComponentMutation {
	return amacc.mutation
}

// Save creates the AsMarketAppComponent in the database.
func (amacc *AsMarketAppComponentCreate) Save(ctx context.Context) (*AsMarketAppComponent, error) {
	var (
		err  error
		node *AsMarketAppComponent
	)
	amacc.defaults()
	if len(amacc.hooks) == 0 {
		if err = amacc.check(); err != nil {
			return nil, err
		}
		node, err = amacc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketAppComponentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = amacc.check(); err != nil {
				return nil, err
			}
			amacc.mutation = mutation
			if node, err = amacc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(amacc.hooks) - 1; i >= 0; i-- {
			if amacc.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amacc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amacc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (amacc *AsMarketAppComponentCreate) SaveX(ctx context.Context) *AsMarketAppComponent {
	v, err := amacc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amacc *AsMarketAppComponentCreate) Exec(ctx context.Context) error {
	_, err := amacc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amacc *AsMarketAppComponentCreate) ExecX(ctx context.Context) {
	if err := amacc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amacc *AsMarketAppComponentCreate) defaults() {
	if _, ok := amacc.mutation.IsDeleted(); !ok {
		v := asmarketappcomponent.DefaultIsDeleted
		amacc.mutation.SetIsDeleted(v)
	}
	if _, ok := amacc.mutation.Status(); !ok {
		v := asmarketappcomponent.DefaultStatus
		amacc.mutation.SetStatus(v)
	}
	if _, ok := amacc.mutation.CreateTime(); !ok {
		v := asmarketappcomponent.DefaultCreateTime()
		amacc.mutation.SetCreateTime(v)
	}
	if _, ok := amacc.mutation.UpdateTime(); !ok {
		v := asmarketappcomponent.DefaultUpdateTime()
		amacc.mutation.SetUpdateTime(v)
	}
	if _, ok := amacc.mutation.ID(); !ok {
		v := asmarketappcomponent.DefaultID()
		amacc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (amacc *AsMarketAppComponentCreate) check() error {
	if _, ok := amacc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`schema: missing required field "AsMarketAppComponent.is_deleted"`)}
	}
	return nil
}

func (amacc *AsMarketAppComponentCreate) sqlSave(ctx context.Context) (*AsMarketAppComponent, error) {
	_node, _spec := amacc.createSpec()
	if err := sqlgraph.CreateNode(ctx, amacc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (amacc *AsMarketAppComponentCreate) createSpec() (*AsMarketAppComponent, *sqlgraph.CreateSpec) {
	var (
		_node = &AsMarketAppComponent{config: amacc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asmarketappcomponent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappcomponent.FieldID,
			},
		}
	)
	if id, ok := amacc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := amacc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := amacc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldName,
		})
		_node.Name = value
	}
	if value, ok := amacc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := amacc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponent.FieldType,
		})
		_node.Type = value
	}
	if value, ok := amacc.mutation.PreviewPic(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldPreviewPic,
		})
		_node.PreviewPic = value
	}
	if value, ok := amacc.mutation.LayoutType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldLayoutType,
		})
		_node.LayoutType = value
	}
	if value, ok := amacc.mutation.LayoutConfig(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldLayoutConfig,
		})
		_node.LayoutConfig = value
	}
	if value, ok := amacc.mutation.TenantCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldTenantCode,
		})
		_node.TenantCode = value
	}
	if value, ok := amacc.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponent.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := amacc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponent.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := amacc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponent.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := amacc.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponent.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := amacc.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponent.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := amacc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketappcomponent.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := amacc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketappcomponent.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if nodes := amacc.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketappcomponent.AppxTable,
			Columns: []string{asmarketappcomponent.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AsMarketAppComponentCreateBulk is the builder for creating many AsMarketAppComponent entities in bulk.
type AsMarketAppComponentCreateBulk struct {
	config
	builders []*AsMarketAppComponentCreate
}

// Save creates the AsMarketAppComponent entities in the database.
func (amaccb *AsMarketAppComponentCreateBulk) Save(ctx context.Context) ([]*AsMarketAppComponent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(amaccb.builders))
	nodes := make([]*AsMarketAppComponent, len(amaccb.builders))
	mutators := make([]Mutator, len(amaccb.builders))
	for i := range amaccb.builders {
		func(i int, root context.Context) {
			builder := amaccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AsMarketAppComponentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, amaccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, amaccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, amaccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (amaccb *AsMarketAppComponentCreateBulk) SaveX(ctx context.Context) []*AsMarketAppComponent {
	v, err := amaccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amaccb *AsMarketAppComponentCreateBulk) Exec(ctx context.Context) error {
	_, err := amaccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amaccb *AsMarketAppComponentCreateBulk) ExecX(ctx context.Context) {
	if err := amaccb.Exec(ctx); err != nil {
		panic(err)
	}
}
