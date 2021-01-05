// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AccountRoleCreate is the builder for creating a AccountRole entity.
type AccountRoleCreate struct {
	config
	mutation *AccountRoleMutation
	hooks    []Hook
}

// SetName sets the name field.
func (arc *AccountRoleCreate) SetName(s string) *AccountRoleCreate {
	arc.mutation.SetName(s)
	return arc
}

// SetAccountID sets the account_id field.
func (arc *AccountRoleCreate) SetAccountID(u uuid.UUID) *AccountRoleCreate {
	arc.mutation.SetAccountID(u)
	return arc
}

// SetMetadata sets the metadata field.
func (arc *AccountRoleCreate) SetMetadata(m map[string]interface{}) *AccountRoleCreate {
	arc.mutation.SetMetadata(m)
	return arc
}

// SetCreatedAt sets the created_at field.
func (arc *AccountRoleCreate) SetCreatedAt(t time.Time) *AccountRoleCreate {
	arc.mutation.SetCreatedAt(t)
	return arc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (arc *AccountRoleCreate) SetNillableCreatedAt(t *time.Time) *AccountRoleCreate {
	if t != nil {
		arc.SetCreatedAt(*t)
	}
	return arc
}

// SetUpdatedAt sets the updated_at field.
func (arc *AccountRoleCreate) SetUpdatedAt(t time.Time) *AccountRoleCreate {
	arc.mutation.SetUpdatedAt(t)
	return arc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (arc *AccountRoleCreate) SetNillableUpdatedAt(t *time.Time) *AccountRoleCreate {
	if t != nil {
		arc.SetUpdatedAt(*t)
	}
	return arc
}

// SetID sets the id field.
func (arc *AccountRoleCreate) SetID(u uuid.UUID) *AccountRoleCreate {
	arc.mutation.SetID(u)
	return arc
}

// SetAccountsID sets the accounts edge to Account by id.
func (arc *AccountRoleCreate) SetAccountsID(id uuid.UUID) *AccountRoleCreate {
	arc.mutation.SetAccountsID(id)
	return arc
}

// SetAccounts sets the accounts edge to Account.
func (arc *AccountRoleCreate) SetAccounts(a *Account) *AccountRoleCreate {
	return arc.SetAccountsID(a.ID)
}

// Mutation returns the AccountRoleMutation object of the builder.
func (arc *AccountRoleCreate) Mutation() *AccountRoleMutation {
	return arc.mutation
}

// Save creates the AccountRole in the database.
func (arc *AccountRoleCreate) Save(ctx context.Context) (*AccountRole, error) {
	var (
		err  error
		node *AccountRole
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			node, err = arc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			mut = arc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AccountRoleCreate) SaveX(ctx context.Context) *AccountRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (arc *AccountRoleCreate) defaults() {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		v := accountrole.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		v := accountrole.DefaultUpdatedAt()
		arc.mutation.SetUpdatedAt(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		v := accountrole.DefaultID()
		arc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AccountRoleCreate) check() error {
	if _, ok := arc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("models: missing required field \"name\"")}
	}
	if v, ok := arc.mutation.Name(); ok {
		if err := accountrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := arc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New("models: missing required field \"account_id\"")}
	}
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("models: missing required field \"updated_at\"")}
	}
	if _, ok := arc.mutation.AccountsID(); !ok {
		return &ValidationError{Name: "accounts", err: errors.New("models: missing required edge \"accounts\"")}
	}
	return nil
}

func (arc *AccountRoleCreate) sqlSave(ctx context.Context) (*AccountRole, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (arc *AccountRoleCreate) createSpec() (*AccountRole, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountRole{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: accountrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountrole.FieldID,
			},
		}
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountrole.FieldName,
		})
		_node.Name = value
	}
	if value, ok := arc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: accountrole.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := arc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: accountrole.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accountrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accountrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := arc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountrole.AccountsTable,
			Columns: []string{accountrole.AccountsColumn},
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

// AccountRoleCreateBulk is the builder for creating a bulk of AccountRole entities.
type AccountRoleCreateBulk struct {
	config
	builders []*AccountRoleCreate
}

// Save creates the AccountRole entities in the database.
func (arcb *AccountRoleCreateBulk) Save(ctx context.Context) ([]*AccountRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AccountRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (arcb *AccountRoleCreateBulk) SaveX(ctx context.Context) []*AccountRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
