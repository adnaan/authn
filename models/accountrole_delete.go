// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"github.com/adnaan/authzen/models/accountrole"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AccountRoleDelete is the builder for deleting a AccountRole entity.
type AccountRoleDelete struct {
	config
	hooks    []Hook
	mutation *AccountRoleMutation
}

// Where adds a new predicate to the delete builder.
func (ard *AccountRoleDelete) Where(ps ...predicate.AccountRole) *AccountRoleDelete {
	ard.mutation.predicates = append(ard.mutation.predicates, ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AccountRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ard.hooks) == 0 {
		affected, err = ard.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ard.mutation = mutation
			affected, err = ard.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ard.hooks) - 1; i >= 0; i-- {
			mut = ard.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ard.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AccountRoleDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AccountRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: accountrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountrole.FieldID,
			},
		},
	}
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
}

// AccountRoleDeleteOne is the builder for deleting a single AccountRole entity.
type AccountRoleDeleteOne struct {
	ard *AccountRoleDelete
}

// Exec executes the deletion query.
func (ardo *AccountRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AccountRoleDeleteOne) ExecX(ctx context.Context) {
	ardo.ard.ExecX(ctx)
}
