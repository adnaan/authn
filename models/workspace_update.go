// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceUpdate is the builder for updating Workspace entities.
type WorkspaceUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// Where adds a new predicate for the builder.
func (wu *WorkspaceUpdate) Where(ps ...predicate.Workspace) *WorkspaceUpdate {
	wu.mutation.predicates = append(wu.mutation.predicates, ps...)
	return wu
}

// SetName sets the name field.
func (wu *WorkspaceUpdate) SetName(s string) *WorkspaceUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetPlan sets the plan field.
func (wu *WorkspaceUpdate) SetPlan(s string) *WorkspaceUpdate {
	wu.mutation.SetPlan(s)
	return wu
}

// SetDescription sets the description field.
func (wu *WorkspaceUpdate) SetDescription(s string) *WorkspaceUpdate {
	wu.mutation.SetDescription(s)
	return wu
}

// SetNillableDescription sets the description field if the given value is not nil.
func (wu *WorkspaceUpdate) SetNillableDescription(s *string) *WorkspaceUpdate {
	if s != nil {
		wu.SetDescription(*s)
	}
	return wu
}

// ClearDescription clears the value of description.
func (wu *WorkspaceUpdate) ClearDescription() *WorkspaceUpdate {
	wu.mutation.ClearDescription()
	return wu
}

// SetMetadata sets the metadata field.
func (wu *WorkspaceUpdate) SetMetadata(m map[string]interface{}) *WorkspaceUpdate {
	wu.mutation.SetMetadata(m)
	return wu
}

// ClearMetadata clears the value of metadata.
func (wu *WorkspaceUpdate) ClearMetadata() *WorkspaceUpdate {
	wu.mutation.ClearMetadata()
	return wu
}

// SetUpdatedAt sets the updated_at field.
func (wu *WorkspaceUpdate) SetUpdatedAt(t time.Time) *WorkspaceUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetOwnerID sets the owner edge to Account by id.
func (wu *WorkspaceUpdate) SetOwnerID(id uuid.UUID) *WorkspaceUpdate {
	wu.mutation.SetOwnerID(id)
	return wu
}

// SetOwner sets the owner edge to Account.
func (wu *WorkspaceUpdate) SetOwner(a *Account) *WorkspaceUpdate {
	return wu.SetOwnerID(a.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (wu *WorkspaceUpdate) AddWorkspaceRoleIDs(ids ...uuid.UUID) *WorkspaceUpdate {
	wu.mutation.AddWorkspaceRoleIDs(ids...)
	return wu
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (wu *WorkspaceUpdate) AddWorkspaceRoles(w ...*WorkspaceRole) *WorkspaceUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWorkspaceRoleIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (wu *WorkspaceUpdate) AddGroupIDs(ids ...uuid.UUID) *WorkspaceUpdate {
	wu.mutation.AddGroupIDs(ids...)
	return wu
}

// AddGroups adds the groups edges to Group.
func (wu *WorkspaceUpdate) AddGroups(g ...*Group) *WorkspaceUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wu.AddGroupIDs(ids...)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wu *WorkspaceUpdate) Mutation() *WorkspaceMutation {
	return wu.mutation
}

// ClearOwner clears the "owner" edge to type Account.
func (wu *WorkspaceUpdate) ClearOwner() *WorkspaceUpdate {
	wu.mutation.ClearOwner()
	return wu
}

// ClearWorkspaceRoles clears all "workspace_roles" edges to type WorkspaceRole.
func (wu *WorkspaceUpdate) ClearWorkspaceRoles() *WorkspaceUpdate {
	wu.mutation.ClearWorkspaceRoles()
	return wu
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (wu *WorkspaceUpdate) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) *WorkspaceUpdate {
	wu.mutation.RemoveWorkspaceRoleIDs(ids...)
	return wu
}

// RemoveWorkspaceRoles removes workspace_roles edges to WorkspaceRole.
func (wu *WorkspaceUpdate) RemoveWorkspaceRoles(w ...*WorkspaceRole) *WorkspaceUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWorkspaceRoleIDs(ids...)
}

// ClearGroups clears all "groups" edges to type Group.
func (wu *WorkspaceUpdate) ClearGroups() *WorkspaceUpdate {
	wu.mutation.ClearGroups()
	return wu
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (wu *WorkspaceUpdate) RemoveGroupIDs(ids ...uuid.UUID) *WorkspaceUpdate {
	wu.mutation.RemoveGroupIDs(ids...)
	return wu
}

// RemoveGroups removes groups edges to Group.
func (wu *WorkspaceUpdate) RemoveGroups(g ...*Group) *WorkspaceUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkspaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkspaceUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkspaceUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkspaceUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WorkspaceUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := workspace.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkspaceUpdate) check() error {
	if v, ok := wu.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := wu.mutation.Plan(); ok {
		if err := workspace.PlanValidator(v); err != nil {
			return &ValidationError{Name: "plan", err: fmt.Errorf("models: validator failed for field \"plan\": %w", err)}
		}
	}
	if v, ok := wu.mutation.Description(); ok {
		if err := workspace.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("models: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := wu.mutation.OwnerID(); wu.mutation.OwnerCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (wu *WorkspaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspace.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
	}
	if value, ok := wu.mutation.Plan(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldPlan,
		})
	}
	if value, ok := wu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldDescription,
		})
	}
	if wu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workspace.FieldDescription,
		})
	}
	if value, ok := wu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldMetadata,
		})
	}
	if wu.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: workspace.FieldMetadata,
		})
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
	}
	if wu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWorkspaceRolesIDs(); len(nodes) > 0 && !wu.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !wu.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceUpdateOne is the builder for updating a single Workspace entity.
type WorkspaceUpdateOne struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// SetName sets the name field.
func (wuo *WorkspaceUpdateOne) SetName(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetPlan sets the plan field.
func (wuo *WorkspaceUpdateOne) SetPlan(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetPlan(s)
	return wuo
}

// SetDescription sets the description field.
func (wuo *WorkspaceUpdateOne) SetDescription(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetDescription(s)
	return wuo
}

// SetNillableDescription sets the description field if the given value is not nil.
func (wuo *WorkspaceUpdateOne) SetNillableDescription(s *string) *WorkspaceUpdateOne {
	if s != nil {
		wuo.SetDescription(*s)
	}
	return wuo
}

// ClearDescription clears the value of description.
func (wuo *WorkspaceUpdateOne) ClearDescription() *WorkspaceUpdateOne {
	wuo.mutation.ClearDescription()
	return wuo
}

// SetMetadata sets the metadata field.
func (wuo *WorkspaceUpdateOne) SetMetadata(m map[string]interface{}) *WorkspaceUpdateOne {
	wuo.mutation.SetMetadata(m)
	return wuo
}

// ClearMetadata clears the value of metadata.
func (wuo *WorkspaceUpdateOne) ClearMetadata() *WorkspaceUpdateOne {
	wuo.mutation.ClearMetadata()
	return wuo
}

// SetUpdatedAt sets the updated_at field.
func (wuo *WorkspaceUpdateOne) SetUpdatedAt(t time.Time) *WorkspaceUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetOwnerID sets the owner edge to Account by id.
func (wuo *WorkspaceUpdateOne) SetOwnerID(id uuid.UUID) *WorkspaceUpdateOne {
	wuo.mutation.SetOwnerID(id)
	return wuo
}

// SetOwner sets the owner edge to Account.
func (wuo *WorkspaceUpdateOne) SetOwner(a *Account) *WorkspaceUpdateOne {
	return wuo.SetOwnerID(a.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (wuo *WorkspaceUpdateOne) AddWorkspaceRoleIDs(ids ...uuid.UUID) *WorkspaceUpdateOne {
	wuo.mutation.AddWorkspaceRoleIDs(ids...)
	return wuo
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (wuo *WorkspaceUpdateOne) AddWorkspaceRoles(w ...*WorkspaceRole) *WorkspaceUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWorkspaceRoleIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (wuo *WorkspaceUpdateOne) AddGroupIDs(ids ...uuid.UUID) *WorkspaceUpdateOne {
	wuo.mutation.AddGroupIDs(ids...)
	return wuo
}

// AddGroups adds the groups edges to Group.
func (wuo *WorkspaceUpdateOne) AddGroups(g ...*Group) *WorkspaceUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wuo.AddGroupIDs(ids...)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wuo *WorkspaceUpdateOne) Mutation() *WorkspaceMutation {
	return wuo.mutation
}

// ClearOwner clears the "owner" edge to type Account.
func (wuo *WorkspaceUpdateOne) ClearOwner() *WorkspaceUpdateOne {
	wuo.mutation.ClearOwner()
	return wuo
}

// ClearWorkspaceRoles clears all "workspace_roles" edges to type WorkspaceRole.
func (wuo *WorkspaceUpdateOne) ClearWorkspaceRoles() *WorkspaceUpdateOne {
	wuo.mutation.ClearWorkspaceRoles()
	return wuo
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (wuo *WorkspaceUpdateOne) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) *WorkspaceUpdateOne {
	wuo.mutation.RemoveWorkspaceRoleIDs(ids...)
	return wuo
}

// RemoveWorkspaceRoles removes workspace_roles edges to WorkspaceRole.
func (wuo *WorkspaceUpdateOne) RemoveWorkspaceRoles(w ...*WorkspaceRole) *WorkspaceUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWorkspaceRoleIDs(ids...)
}

// ClearGroups clears all "groups" edges to type Group.
func (wuo *WorkspaceUpdateOne) ClearGroups() *WorkspaceUpdateOne {
	wuo.mutation.ClearGroups()
	return wuo
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (wuo *WorkspaceUpdateOne) RemoveGroupIDs(ids ...uuid.UUID) *WorkspaceUpdateOne {
	wuo.mutation.RemoveGroupIDs(ids...)
	return wuo
}

// RemoveGroups removes groups edges to Group.
func (wuo *WorkspaceUpdateOne) RemoveGroups(g ...*Group) *WorkspaceUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (wuo *WorkspaceUpdateOne) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) SaveX(ctx context.Context) *Workspace {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkspaceUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WorkspaceUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := workspace.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkspaceUpdateOne) check() error {
	if v, ok := wuo.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.Plan(); ok {
		if err := workspace.PlanValidator(v); err != nil {
			return &ValidationError{Name: "plan", err: fmt.Errorf("models: validator failed for field \"plan\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.Description(); ok {
		if err := workspace.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("models: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := wuo.mutation.OwnerID(); wuo.mutation.OwnerCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (wuo *WorkspaceUpdateOne) sqlSave(ctx context.Context) (_node *Workspace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspace.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Workspace.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := wuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
	}
	if value, ok := wuo.mutation.Plan(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldPlan,
		})
	}
	if value, ok := wuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldDescription,
		})
	}
	if wuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workspace.FieldDescription,
		})
	}
	if value, ok := wuo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldMetadata,
		})
	}
	if wuo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: workspace.FieldMetadata,
		})
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
	}
	if wuo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWorkspaceRolesIDs(); len(nodes) > 0 && !wuo.mutation.WorkspaceRolesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !wuo.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Workspace{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
