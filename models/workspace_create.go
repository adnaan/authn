// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceCreate is the builder for creating a Workspace entity.
type WorkspaceCreate struct {
	config
	mutation *WorkspaceMutation
	hooks    []Hook
}

// SetName sets the name field.
func (wc *WorkspaceCreate) SetName(s string) *WorkspaceCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetPlan sets the plan field.
func (wc *WorkspaceCreate) SetPlan(s string) *WorkspaceCreate {
	wc.mutation.SetPlan(s)
	return wc
}

// SetDescription sets the description field.
func (wc *WorkspaceCreate) SetDescription(s string) *WorkspaceCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableDescription(s *string) *WorkspaceCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetMetadata sets the metadata field.
func (wc *WorkspaceCreate) SetMetadata(m map[string]interface{}) *WorkspaceCreate {
	wc.mutation.SetMetadata(m)
	return wc
}

// SetCreatedAt sets the created_at field.
func (wc *WorkspaceCreate) SetCreatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableCreatedAt(t *time.Time) *WorkspaceCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the updated_at field.
func (wc *WorkspaceCreate) SetUpdatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableUpdatedAt(t *time.Time) *WorkspaceCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetID sets the id field.
func (wc *WorkspaceCreate) SetID(u uuid.UUID) *WorkspaceCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetOwnerID sets the owner edge to Account by id.
func (wc *WorkspaceCreate) SetOwnerID(id uuid.UUID) *WorkspaceCreate {
	wc.mutation.SetOwnerID(id)
	return wc
}

// SetOwner sets the owner edge to Account.
func (wc *WorkspaceCreate) SetOwner(a *Account) *WorkspaceCreate {
	return wc.SetOwnerID(a.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (wc *WorkspaceCreate) AddWorkspaceRoleIDs(ids ...uuid.UUID) *WorkspaceCreate {
	wc.mutation.AddWorkspaceRoleIDs(ids...)
	return wc
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (wc *WorkspaceCreate) AddWorkspaceRoles(w ...*WorkspaceRole) *WorkspaceCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWorkspaceRoleIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (wc *WorkspaceCreate) AddGroupIDs(ids ...uuid.UUID) *WorkspaceCreate {
	wc.mutation.AddGroupIDs(ids...)
	return wc
}

// AddGroups adds the groups edges to Group.
func (wc *WorkspaceCreate) AddGroups(g ...*Group) *WorkspaceCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wc.AddGroupIDs(ids...)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wc *WorkspaceCreate) Mutation() *WorkspaceMutation {
	return wc.mutation
}

// Save creates the Workspace in the database.
func (wc *WorkspaceCreate) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			node, err = wc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkspaceCreate) SaveX(ctx context.Context) *Workspace {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wc *WorkspaceCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workspace.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := workspace.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workspace.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkspaceCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("models: missing required field \"name\"")}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wc.mutation.Plan(); !ok {
		return &ValidationError{Name: "plan", err: errors.New("models: missing required field \"plan\"")}
	}
	if v, ok := wc.mutation.Plan(); ok {
		if err := workspace.PlanValidator(v); err != nil {
			return &ValidationError{Name: "plan", err: fmt.Errorf("models: validator failed for field \"plan\": %w", err)}
		}
	}
	if v, ok := wc.mutation.Description(); ok {
		if err := workspace.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("models: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("models: missing required field \"created_at\"")}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("models: missing required field \"updated_at\"")}
	}
	if _, ok := wc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("models: missing required edge \"owner\"")}
	}
	return nil
}

func (wc *WorkspaceCreate) sqlSave(ctx context.Context) (*Workspace, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (wc *WorkspaceCreate) createSpec() (*Workspace, *sqlgraph.CreateSpec) {
	var (
		_node = &Workspace{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workspace.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspace.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
		_node.Name = value
	}
	if value, ok := wc.mutation.Plan(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldPlan,
		})
		_node.Plan = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := wc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := wc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workspace.OwnerTable,
			Columns: []string{workspace.OwnerColumn},
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
	if nodes := wc.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceRolesTable,
			Columns: []string{workspace.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.GroupsTable,
			Columns: []string{workspace.GroupsColumn},
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
	return _node, _spec
}

// WorkspaceCreateBulk is the builder for creating a bulk of Workspace entities.
type WorkspaceCreateBulk struct {
	config
	builders []*WorkspaceCreate
}

// Save creates the Workspace entities in the database.
func (wcb *WorkspaceCreateBulk) Save(ctx context.Context) ([]*Workspace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workspace, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkspaceMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (wcb *WorkspaceCreateBulk) SaveX(ctx context.Context) []*Workspace {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
