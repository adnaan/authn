// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/grouprole"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// GroupRoleCreate is the builder for creating a GroupRole entity.
type GroupRoleCreate struct {
	config
	mutation *GroupRoleMutation
	hooks    []Hook
}

// SetName sets the name field.
func (grc *GroupRoleCreate) SetName(s string) *GroupRoleCreate {
	grc.mutation.SetName(s)
	return grc
}

// SetGroupID sets the group_id field.
func (grc *GroupRoleCreate) SetGroupID(u uuid.UUID) *GroupRoleCreate {
	grc.mutation.SetGroupID(u)
	return grc
}

// SetAccountID sets the account_id field.
func (grc *GroupRoleCreate) SetAccountID(u uuid.UUID) *GroupRoleCreate {
	grc.mutation.SetAccountID(u)
	return grc
}

// SetMetadata sets the metadata field.
func (grc *GroupRoleCreate) SetMetadata(m map[string]interface{}) *GroupRoleCreate {
	grc.mutation.SetMetadata(m)
	return grc
}

// SetCreatedAt sets the created_at field.
func (grc *GroupRoleCreate) SetCreatedAt(t time.Time) *GroupRoleCreate {
	grc.mutation.SetCreatedAt(t)
	return grc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (grc *GroupRoleCreate) SetNillableCreatedAt(t *time.Time) *GroupRoleCreate {
	if t != nil {
		grc.SetCreatedAt(*t)
	}
	return grc
}

// SetUpdatedAt sets the updated_at field.
func (grc *GroupRoleCreate) SetUpdatedAt(t time.Time) *GroupRoleCreate {
	grc.mutation.SetUpdatedAt(t)
	return grc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (grc *GroupRoleCreate) SetNillableUpdatedAt(t *time.Time) *GroupRoleCreate {
	if t != nil {
		grc.SetUpdatedAt(*t)
	}
	return grc
}

// SetID sets the id field.
func (grc *GroupRoleCreate) SetID(u uuid.UUID) *GroupRoleCreate {
	grc.mutation.SetID(u)
	return grc
}

// SetGroupsID sets the groups edge to Group by id.
func (grc *GroupRoleCreate) SetGroupsID(id uuid.UUID) *GroupRoleCreate {
	grc.mutation.SetGroupsID(id)
	return grc
}

// SetGroups sets the groups edge to Group.
func (grc *GroupRoleCreate) SetGroups(g *Group) *GroupRoleCreate {
	return grc.SetGroupsID(g.ID)
}

// SetAccountsID sets the accounts edge to Account by id.
func (grc *GroupRoleCreate) SetAccountsID(id uuid.UUID) *GroupRoleCreate {
	grc.mutation.SetAccountsID(id)
	return grc
}

// SetAccounts sets the accounts edge to Account.
func (grc *GroupRoleCreate) SetAccounts(a *Account) *GroupRoleCreate {
	return grc.SetAccountsID(a.ID)
}

// Mutation returns the GroupRoleMutation object of the builder.
func (grc *GroupRoleCreate) Mutation() *GroupRoleMutation {
	return grc.mutation
}

// Save creates the GroupRole in the database.
func (grc *GroupRoleCreate) Save(ctx context.Context) (*GroupRole, error) {
	var (
		err  error
		node *GroupRole
	)
	grc.defaults()
	if len(grc.hooks) == 0 {
		if err = grc.check(); err != nil {
			return nil, err
		}
		node, err = grc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = grc.check(); err != nil {
				return nil, err
			}
			grc.mutation = mutation
			node, err = grc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(grc.hooks) - 1; i >= 0; i-- {
			mut = grc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, grc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (grc *GroupRoleCreate) SaveX(ctx context.Context) *GroupRole {
	v, err := grc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (grc *GroupRoleCreate) defaults() {
	if _, ok := grc.mutation.CreatedAt(); !ok {
		v := grouprole.DefaultCreatedAt()
		grc.mutation.SetCreatedAt(v)
	}
	if _, ok := grc.mutation.UpdatedAt(); !ok {
		v := grouprole.DefaultUpdatedAt()
		grc.mutation.SetUpdatedAt(v)
	}
	if _, ok := grc.mutation.ID(); !ok {
		v := grouprole.DefaultID()
		grc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (grc *GroupRoleCreate) check() error {
	if _, ok := grc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("models: missing required field \"name\"")}
	}
	if v, ok := grc.mutation.Name(); ok {
		if err := grouprole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := grc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New("models: missing required field \"group_id\"")}
	}
	if _, ok := grc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New("models: missing required field \"account_id\"")}
	}
	if _, ok := grc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	if _, ok := grc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("models: missing required field \"updated_at\"")}
	}
	if _, ok := grc.mutation.GroupsID(); !ok {
		return &ValidationError{Name: "groups", err: errors.New("models: missing required edge \"groups\"")}
	}
	if _, ok := grc.mutation.AccountsID(); !ok {
		return &ValidationError{Name: "accounts", err: errors.New("models: missing required edge \"accounts\"")}
	}
	return nil
}

func (grc *GroupRoleCreate) sqlSave(ctx context.Context) (*GroupRole, error) {
	_node, _spec := grc.createSpec()
	if err := sqlgraph.CreateNode(ctx, grc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (grc *GroupRoleCreate) createSpec() (*GroupRole, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupRole{config: grc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: grouprole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: grouprole.FieldID,
			},
		}
	)
	if id, ok := grc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := grc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grouprole.FieldName,
		})
		_node.Name = value
	}
	if value, ok := grc.mutation.GroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: grouprole.FieldGroupID,
		})
		_node.GroupID = value
	}
	if value, ok := grc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: grouprole.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := grc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: grouprole.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := grc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grouprole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := grc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grouprole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := grc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   grouprole.GroupsTable,
			Columns: []string{grouprole.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := grc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   grouprole.AccountsTable,
			Columns: []string{grouprole.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupRoleCreateBulk is the builder for creating a bulk of GroupRole entities.
type GroupRoleCreateBulk struct {
	config
	builders []*GroupRoleCreate
}

// Save creates the GroupRole entities in the database.
func (grcb *GroupRoleCreateBulk) Save(ctx context.Context) ([]*GroupRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(grcb.builders))
	nodes := make([]*GroupRole, len(grcb.builders))
	mutators := make([]Mutator, len(grcb.builders))
	for i := range grcb.builders {
		func(i int, root context.Context) {
			builder := grcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, grcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, grcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, grcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (grcb *GroupRoleCreateBulk) SaveX(ctx context.Context) []*GroupRole {
	v, err := grcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
