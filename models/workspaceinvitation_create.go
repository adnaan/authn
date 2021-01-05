// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceInvitationCreate is the builder for creating a WorkspaceInvitation entity.
type WorkspaceInvitationCreate struct {
	config
	mutation *WorkspaceInvitationMutation
	hooks    []Hook
}

// SetWorkspaceID sets the workspace_id field.
func (wic *WorkspaceInvitationCreate) SetWorkspaceID(u uuid.UUID) *WorkspaceInvitationCreate {
	wic.mutation.SetWorkspaceID(u)
	return wic
}

// SetRole sets the role field.
func (wic *WorkspaceInvitationCreate) SetRole(s string) *WorkspaceInvitationCreate {
	wic.mutation.SetRole(s)
	return wic
}

// SetEmail sets the email field.
func (wic *WorkspaceInvitationCreate) SetEmail(s string) *WorkspaceInvitationCreate {
	wic.mutation.SetEmail(s)
	return wic
}

// SetCreatedAt sets the created_at field.
func (wic *WorkspaceInvitationCreate) SetCreatedAt(t time.Time) *WorkspaceInvitationCreate {
	wic.mutation.SetCreatedAt(t)
	return wic
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (wic *WorkspaceInvitationCreate) SetNillableCreatedAt(t *time.Time) *WorkspaceInvitationCreate {
	if t != nil {
		wic.SetCreatedAt(*t)
	}
	return wic
}

// SetID sets the id field.
func (wic *WorkspaceInvitationCreate) SetID(u uuid.UUID) *WorkspaceInvitationCreate {
	wic.mutation.SetID(u)
	return wic
}

// Mutation returns the WorkspaceInvitationMutation object of the builder.
func (wic *WorkspaceInvitationCreate) Mutation() *WorkspaceInvitationMutation {
	return wic.mutation
}

// Save creates the WorkspaceInvitation in the database.
func (wic *WorkspaceInvitationCreate) Save(ctx context.Context) (*WorkspaceInvitation, error) {
	var (
		err  error
		node *WorkspaceInvitation
	)
	wic.defaults()
	if len(wic.hooks) == 0 {
		if err = wic.check(); err != nil {
			return nil, err
		}
		node, err = wic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wic.check(); err != nil {
				return nil, err
			}
			wic.mutation = mutation
			node, err = wic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wic.hooks) - 1; i >= 0; i-- {
			mut = wic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wic *WorkspaceInvitationCreate) SaveX(ctx context.Context) *WorkspaceInvitation {
	v, err := wic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wic *WorkspaceInvitationCreate) defaults() {
	if _, ok := wic.mutation.CreatedAt(); !ok {
		v := workspaceinvitation.DefaultCreatedAt()
		wic.mutation.SetCreatedAt(v)
	}
	if _, ok := wic.mutation.ID(); !ok {
		v := workspaceinvitation.DefaultID()
		wic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wic *WorkspaceInvitationCreate) check() error {
	if _, ok := wic.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New("models: missing required field \"workspace_id\"")}
	}
	if _, ok := wic.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New("models: missing required field \"role\"")}
	}
	if v, ok := wic.mutation.Role(); ok {
		if err := workspaceinvitation.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("models: validator failed for field \"role\": %w", err)}
		}
	}
	if _, ok := wic.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("models: missing required field \"email\"")}
	}
	if v, ok := wic.mutation.Email(); ok {
		if err := workspaceinvitation.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := wic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	return nil
}

func (wic *WorkspaceInvitationCreate) sqlSave(ctx context.Context) (*WorkspaceInvitation, error) {
	_node, _spec := wic.createSpec()
	if err := sqlgraph.CreateNode(ctx, wic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (wic *WorkspaceInvitationCreate) createSpec() (*WorkspaceInvitation, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkspaceInvitation{config: wic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workspaceinvitation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspaceinvitation.FieldID,
			},
		}
	)
	if id, ok := wic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wic.mutation.WorkspaceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: workspaceinvitation.FieldWorkspaceID,
		})
		_node.WorkspaceID = value
	}
	if value, ok := wic.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceinvitation.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := wic.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspaceinvitation.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := wic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspaceinvitation.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// WorkspaceInvitationCreateBulk is the builder for creating a bulk of WorkspaceInvitation entities.
type WorkspaceInvitationCreateBulk struct {
	config
	builders []*WorkspaceInvitationCreate
}

// Save creates the WorkspaceInvitation entities in the database.
func (wicb *WorkspaceInvitationCreateBulk) Save(ctx context.Context) ([]*WorkspaceInvitation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wicb.builders))
	nodes := make([]*WorkspaceInvitation, len(wicb.builders))
	mutators := make([]Mutator, len(wicb.builders))
	for i := range wicb.builders {
		func(i int, root context.Context) {
			builder := wicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkspaceInvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, wicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (wicb *WorkspaceInvitationCreateBulk) SaveX(ctx context.Context) []*WorkspaceInvitation {
	v, err := wicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
