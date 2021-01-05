// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"github.com/adnaan/authzen/models/predicate"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// WorkspaceRoleDelete is the builder for deleting a WorkspaceRole entity.
type WorkspaceRoleDelete struct {
	config
	hooks    []Hook
	mutation *WorkspaceRoleMutation
}

// Where adds a new predicate to the delete builder.
func (wrd *WorkspaceRoleDelete) Where(ps ...predicate.WorkspaceRole) *WorkspaceRoleDelete {
	wrd.mutation.predicates = append(wrd.mutation.predicates, ps...)
	return wrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wrd *WorkspaceRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wrd.hooks) == 0 {
		affected, err = wrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wrd.mutation = mutation
			affected, err = wrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wrd.hooks) - 1; i >= 0; i-- {
			mut = wrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrd *WorkspaceRoleDelete) ExecX(ctx context.Context) int {
	n, err := wrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wrd *WorkspaceRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workspacerole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspacerole.FieldID,
			},
		},
	}
	if ps := wrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wrd.driver, _spec)
}

// WorkspaceRoleDeleteOne is the builder for deleting a single WorkspaceRole entity.
type WorkspaceRoleDeleteOne struct {
	wrd *WorkspaceRoleDelete
}

// Exec executes the deletion query.
func (wrdo *WorkspaceRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := wrdo.wrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workspacerole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wrdo *WorkspaceRoleDeleteOne) ExecX(ctx context.Context) {
	wrdo.wrd.ExecX(ctx)
}
