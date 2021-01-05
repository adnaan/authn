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
)

// WorkspaceInvitationDelete is the builder for deleting a WorkspaceInvitation entity.
type WorkspaceInvitationDelete struct {
	config
	hooks    []Hook
	mutation *WorkspaceInvitationMutation
}

// Where adds a new predicate to the delete builder.
func (wid *WorkspaceInvitationDelete) Where(ps ...predicate.WorkspaceInvitation) *WorkspaceInvitationDelete {
	wid.mutation.predicates = append(wid.mutation.predicates, ps...)
	return wid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wid *WorkspaceInvitationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wid.hooks) == 0 {
		affected, err = wid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wid.mutation = mutation
			affected, err = wid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wid.hooks) - 1; i >= 0; i-- {
			mut = wid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wid *WorkspaceInvitationDelete) ExecX(ctx context.Context) int {
	n, err := wid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wid *WorkspaceInvitationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workspaceinvitation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspaceinvitation.FieldID,
			},
		},
	}
	if ps := wid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wid.driver, _spec)
}

// WorkspaceInvitationDeleteOne is the builder for deleting a single WorkspaceInvitation entity.
type WorkspaceInvitationDeleteOne struct {
	wid *WorkspaceInvitationDelete
}

// Exec executes the deletion query.
func (wido *WorkspaceInvitationDeleteOne) Exec(ctx context.Context) error {
	n, err := wido.wid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workspaceinvitation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wido *WorkspaceInvitationDeleteOne) ExecX(ctx context.Context) {
	wido.wid.ExecX(ctx)
}
