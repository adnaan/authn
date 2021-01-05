// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"github.com/adnaan/authzen/models/predicate"
	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceInvitationUpdate is the builder for updating WorkspaceInvitation entities.
type WorkspaceInvitationUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceInvitationMutation
}

// Where adds a new predicate for the builder.
func (wiu *WorkspaceInvitationUpdate) Where(ps ...predicate.WorkspaceInvitation) *WorkspaceInvitationUpdate {
	wiu.mutation.predicates = append(wiu.mutation.predicates, ps...)
	return wiu
}

// SetWorkspaceID sets the workspace_id field.
func (wiu *WorkspaceInvitationUpdate) SetWorkspaceID(u uuid.UUID) *WorkspaceInvitationUpdate {
	wiu.mutation.SetWorkspaceID(u)
	return wiu
}

// SetRole sets the role field.
func (wiu *WorkspaceInvitationUpdate) SetRole(s string) *WorkspaceInvitationUpdate {
	wiu.mutation.SetRole(s)
	return wiu
}

// SetEmail sets the email field.
func (wiu *WorkspaceInvitationUpdate) SetEmail(s string) *WorkspaceInvitationUpdate {
	wiu.mutation.SetEmail(s)
	return wiu
}

// Mutation returns the WorkspaceInvitationMutation object of the builder.
func (wiu *WorkspaceInvitationUpdate) Mutation() *WorkspaceInvitationMutation {
	return wiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wiu *WorkspaceInvitationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wiu.hooks) == 0 {
		if err = wiu.check(); err != nil {
			return 0, err
		}
		affected, err = wiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wiu.check(); err != nil {
				return 0, err
			}
			wiu.mutation = mutation
			affected, err = wiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wiu.hooks) - 1; i >= 0; i-- {
			mut = wiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wiu *WorkspaceInvitationUpdate) SaveX(ctx context.Context) int {
	affected, err := wiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wiu *WorkspaceInvitationUpdate) Exec(ctx context.Context) error {
	_, err := wiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wiu *WorkspaceInvitationUpdate) ExecX(ctx context.Context) {
	if err := wiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wiu *WorkspaceInvitationUpdate) check() error {
	if v, ok := wiu.mutation.Role(); ok {
		if err := workspaceinvitation.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("models: validator failed for field \"role\": %w", err)}
		}
	}
	if v, ok := wiu.mutation.Email(); ok {
		if err := workspaceinvitation.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	return nil
}

func (wiu *WorkspaceInvitationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspaceinvitation.Table,
			Columns: workspaceinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspaceinvitation.FieldID,
			},
		},
	}
	if ps := wiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wiu.mutation.WorkspaceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspaceinvitation.FieldWorkspaceID,
		})
	}
	if value, ok := wiu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceinvitation.FieldRole,
		})
	}
	if value, ok := wiu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceinvitation.FieldEmail,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceinvitation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceInvitationUpdateOne is the builder for updating a single WorkspaceInvitation entity.
type WorkspaceInvitationUpdateOne struct {
	config
	hooks    []Hook
	mutation *WorkspaceInvitationMutation
}

// SetWorkspaceID sets the workspace_id field.
func (wiuo *WorkspaceInvitationUpdateOne) SetWorkspaceID(u uuid.UUID) *WorkspaceInvitationUpdateOne {
	wiuo.mutation.SetWorkspaceID(u)
	return wiuo
}

// SetRole sets the role field.
func (wiuo *WorkspaceInvitationUpdateOne) SetRole(s string) *WorkspaceInvitationUpdateOne {
	wiuo.mutation.SetRole(s)
	return wiuo
}

// SetEmail sets the email field.
func (wiuo *WorkspaceInvitationUpdateOne) SetEmail(s string) *WorkspaceInvitationUpdateOne {
	wiuo.mutation.SetEmail(s)
	return wiuo
}

// Mutation returns the WorkspaceInvitationMutation object of the builder.
func (wiuo *WorkspaceInvitationUpdateOne) Mutation() *WorkspaceInvitationMutation {
	return wiuo.mutation
}

// Save executes the query and returns the updated entity.
func (wiuo *WorkspaceInvitationUpdateOne) Save(ctx context.Context) (*WorkspaceInvitation, error) {
	var (
		err  error
		node *WorkspaceInvitation
	)
	if len(wiuo.hooks) == 0 {
		if err = wiuo.check(); err != nil {
			return nil, err
		}
		node, err = wiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wiuo.check(); err != nil {
				return nil, err
			}
			wiuo.mutation = mutation
			node, err = wiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wiuo.hooks) - 1; i >= 0; i-- {
			mut = wiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wiuo *WorkspaceInvitationUpdateOne) SaveX(ctx context.Context) *WorkspaceInvitation {
	node, err := wiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wiuo *WorkspaceInvitationUpdateOne) Exec(ctx context.Context) error {
	_, err := wiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wiuo *WorkspaceInvitationUpdateOne) ExecX(ctx context.Context) {
	if err := wiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wiuo *WorkspaceInvitationUpdateOne) check() error {
	if v, ok := wiuo.mutation.Role(); ok {
		if err := workspaceinvitation.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("models: validator failed for field \"role\": %w", err)}
		}
	}
	if v, ok := wiuo.mutation.Email(); ok {
		if err := workspaceinvitation.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	return nil
}

func (wiuo *WorkspaceInvitationUpdateOne) sqlSave(ctx context.Context) (_node *WorkspaceInvitation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspaceinvitation.Table,
			Columns: workspaceinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspaceinvitation.FieldID,
			},
		},
	}
	id, ok := wiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkspaceInvitation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := wiuo.mutation.WorkspaceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspaceinvitation.FieldWorkspaceID,
		})
	}
	if value, ok := wiuo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceinvitation.FieldRole,
		})
	}
	if value, ok := wiuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceinvitation.FieldEmail,
		})
	}
	_node = &WorkspaceInvitation{config: wiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, wiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceinvitation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
