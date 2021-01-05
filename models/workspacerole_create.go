// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceRoleCreate is the builder for creating a WorkspaceRole entity.
type WorkspaceRoleCreate struct {
	config
	mutation *WorkspaceRoleMutation
	hooks    []Hook
}

// SetName sets the name field.
func (wrc *WorkspaceRoleCreate) SetName(s string) *WorkspaceRoleCreate {
	wrc.mutation.SetName(s)
	return wrc
}

// SetWorkspaceID sets the workspace_id field.
func (wrc *WorkspaceRoleCreate) SetWorkspaceID(u uuid.UUID) *WorkspaceRoleCreate {
	wrc.mutation.SetWorkspaceID(u)
	return wrc
}

// SetAccountID sets the account_id field.
func (wrc *WorkspaceRoleCreate) SetAccountID(u uuid.UUID) *WorkspaceRoleCreate {
	wrc.mutation.SetAccountID(u)
	return wrc
}

// SetMetadata sets the metadata field.
func (wrc *WorkspaceRoleCreate) SetMetadata(m map[string]interface{}) *WorkspaceRoleCreate {
	wrc.mutation.SetMetadata(m)
	return wrc
}

// SetCreatedAt sets the created_at field.
func (wrc *WorkspaceRoleCreate) SetCreatedAt(t time.Time) *WorkspaceRoleCreate {
	wrc.mutation.SetCreatedAt(t)
	return wrc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (wrc *WorkspaceRoleCreate) SetNillableCreatedAt(t *time.Time) *WorkspaceRoleCreate {
	if t != nil {
		wrc.SetCreatedAt(*t)
	}
	return wrc
}

// SetUpdatedAt sets the updated_at field.
func (wrc *WorkspaceRoleCreate) SetUpdatedAt(t time.Time) *WorkspaceRoleCreate {
	wrc.mutation.SetUpdatedAt(t)
	return wrc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (wrc *WorkspaceRoleCreate) SetNillableUpdatedAt(t *time.Time) *WorkspaceRoleCreate {
	if t != nil {
		wrc.SetUpdatedAt(*t)
	}
	return wrc
}

// SetID sets the id field.
func (wrc *WorkspaceRoleCreate) SetID(u uuid.UUID) *WorkspaceRoleCreate {
	wrc.mutation.SetID(u)
	return wrc
}

// SetWorkspacesID sets the workspaces edge to Workspace by id.
func (wrc *WorkspaceRoleCreate) SetWorkspacesID(id uuid.UUID) *WorkspaceRoleCreate {
	wrc.mutation.SetWorkspacesID(id)
	return wrc
}

// SetWorkspaces sets the workspaces edge to Workspace.
func (wrc *WorkspaceRoleCreate) SetWorkspaces(w *Workspace) *WorkspaceRoleCreate {
	return wrc.SetWorkspacesID(w.ID)
}

// SetAccountsID sets the accounts edge to Account by id.
func (wrc *WorkspaceRoleCreate) SetAccountsID(id uuid.UUID) *WorkspaceRoleCreate {
	wrc.mutation.SetAccountsID(id)
	return wrc
}

// SetAccounts sets the accounts edge to Account.
func (wrc *WorkspaceRoleCreate) SetAccounts(a *Account) *WorkspaceRoleCreate {
	return wrc.SetAccountsID(a.ID)
}

// Mutation returns the WorkspaceRoleMutation object of the builder.
func (wrc *WorkspaceRoleCreate) Mutation() *WorkspaceRoleMutation {
	return wrc.mutation
}

// Save creates the WorkspaceRole in the database.
func (wrc *WorkspaceRoleCreate) Save(ctx context.Context) (*WorkspaceRole, error) {
	var (
		err  error
		node *WorkspaceRole
	)
	wrc.defaults()
	if len(wrc.hooks) == 0 {
		if err = wrc.check(); err != nil {
			return nil, err
		}
		node, err = wrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wrc.check(); err != nil {
				return nil, err
			}
			wrc.mutation = mutation
			node, err = wrc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wrc.hooks) - 1; i >= 0; i-- {
			mut = wrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wrc *WorkspaceRoleCreate) SaveX(ctx context.Context) *WorkspaceRole {
	v, err := wrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wrc *WorkspaceRoleCreate) defaults() {
	if _, ok := wrc.mutation.CreatedAt(); !ok {
		v := workspacerole.DefaultCreatedAt()
		wrc.mutation.SetCreatedAt(v)
	}
	if _, ok := wrc.mutation.UpdatedAt(); !ok {
		v := workspacerole.DefaultUpdatedAt()
		wrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wrc.mutation.ID(); !ok {
		v := workspacerole.DefaultID()
		wrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrc *WorkspaceRoleCreate) check() error {
	if _, ok := wrc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("models: missing required field \"name\"")}
	}
	if v, ok := wrc.mutation.Name(); ok {
		if err := workspacerole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wrc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New("models: missing required field \"workspace_id\"")}
	}
	if _, ok := wrc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New("models: missing required field \"account_id\"")}
	}
	if _, ok := wrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	if _, ok := wrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("models: missing required field \"updated_at\"")}
	}
	if _, ok := wrc.mutation.WorkspacesID(); !ok {
		return &ValidationError{Name: "workspaces", err: errors.New("models: missing required edge \"workspaces\"")}
	}
	if _, ok := wrc.mutation.AccountsID(); !ok {
		return &ValidationError{Name: "accounts", err: errors.New("models: missing required edge \"accounts\"")}
	}
	return nil
}

func (wrc *WorkspaceRoleCreate) sqlSave(ctx context.Context) (*WorkspaceRole, error) {
	_node, _spec := wrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (wrc *WorkspaceRoleCreate) createSpec() (*WorkspaceRole, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkspaceRole{config: wrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workspacerole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspacerole.FieldID,
			},
		}
	)
	if id, ok := wrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wrc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspacerole.FieldName,
		})
		_node.Name = value
	}
	if value, ok := wrc.mutation.WorkspaceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspacerole.FieldWorkspaceID,
		})
		_node.WorkspaceID = value
	}
	if value, ok := wrc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspacerole.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := wrc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspacerole.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := wrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspacerole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspacerole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := wrc.mutation.WorkspacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.WorkspacesTable,
			Columns: []string{workspacerole.WorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wrc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspacerole.AccountsTable,
			Columns: []string{workspacerole.AccountsColumn},
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

// WorkspaceRoleCreateBulk is the builder for creating a bulk of WorkspaceRole entities.
type WorkspaceRoleCreateBulk struct {
	config
	builders []*WorkspaceRoleCreate
}

// Save creates the WorkspaceRole entities in the database.
func (wrcb *WorkspaceRoleCreateBulk) Save(ctx context.Context) ([]*WorkspaceRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wrcb.builders))
	nodes := make([]*WorkspaceRole, len(wrcb.builders))
	mutators := make([]Mutator, len(wrcb.builders))
	for i := range wrcb.builders {
		func(i int, root context.Context) {
			builder := wrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkspaceRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, wrcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wrcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (wrcb *WorkspaceRoleCreateBulk) SaveX(ctx context.Context) []*WorkspaceRole {
	v, err := wrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
