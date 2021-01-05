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
	"github.com/adnaan/authzen/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// GroupRoleUpdate is the builder for updating GroupRole entities.
type GroupRoleUpdate struct {
	config
	hooks    []Hook
	mutation *GroupRoleMutation
}

// Where adds a new predicate for the builder.
func (gru *GroupRoleUpdate) Where(ps ...predicate.GroupRole) *GroupRoleUpdate {
	gru.mutation.predicates = append(gru.mutation.predicates, ps...)
	return gru
}

// SetName sets the name field.
func (gru *GroupRoleUpdate) SetName(s string) *GroupRoleUpdate {
	gru.mutation.SetName(s)
	return gru
}

// SetGroupID sets the group_id field.
func (gru *GroupRoleUpdate) SetGroupID(u uuid.UUID) *GroupRoleUpdate {
	gru.mutation.SetGroupID(u)
	return gru
}

// SetAccountID sets the account_id field.
func (gru *GroupRoleUpdate) SetAccountID(u uuid.UUID) *GroupRoleUpdate {
	gru.mutation.SetAccountID(u)
	return gru
}

// SetAttributes sets the attributes field.
func (gru *GroupRoleUpdate) SetAttributes(m map[string]interface{}) *GroupRoleUpdate {
	gru.mutation.SetAttributes(m)
	return gru
}

// ClearAttributes clears the value of attributes.
func (gru *GroupRoleUpdate) ClearAttributes() *GroupRoleUpdate {
	gru.mutation.ClearAttributes()
	return gru
}

// SetUpdatedAt sets the updated_at field.
func (gru *GroupRoleUpdate) SetUpdatedAt(t time.Time) *GroupRoleUpdate {
	gru.mutation.SetUpdatedAt(t)
	return gru
}

// SetGroupsID sets the groups edge to Group by id.
func (gru *GroupRoleUpdate) SetGroupsID(id uuid.UUID) *GroupRoleUpdate {
	gru.mutation.SetGroupsID(id)
	return gru
}

// SetGroups sets the groups edge to Group.
func (gru *GroupRoleUpdate) SetGroups(g *Group) *GroupRoleUpdate {
	return gru.SetGroupsID(g.ID)
}

// SetAccountsID sets the accounts edge to Account by id.
func (gru *GroupRoleUpdate) SetAccountsID(id uuid.UUID) *GroupRoleUpdate {
	gru.mutation.SetAccountsID(id)
	return gru
}

// SetAccounts sets the accounts edge to Account.
func (gru *GroupRoleUpdate) SetAccounts(a *Account) *GroupRoleUpdate {
	return gru.SetAccountsID(a.ID)
}

// Mutation returns the GroupRoleMutation object of the builder.
func (gru *GroupRoleUpdate) Mutation() *GroupRoleMutation {
	return gru.mutation
}

// ClearGroups clears the "groups" edge to type Group.
func (gru *GroupRoleUpdate) ClearGroups() *GroupRoleUpdate {
	gru.mutation.ClearGroups()
	return gru
}

// ClearAccounts clears the "accounts" edge to type Account.
func (gru *GroupRoleUpdate) ClearAccounts() *GroupRoleUpdate {
	gru.mutation.ClearAccounts()
	return gru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GroupRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gru.defaults()
	if len(gru.hooks) == 0 {
		if err = gru.check(); err != nil {
			return 0, err
		}
		affected, err = gru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gru.check(); err != nil {
				return 0, err
			}
			gru.mutation = mutation
			affected, err = gru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gru.hooks) - 1; i >= 0; i-- {
			mut = gru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GroupRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GroupRoleUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GroupRoleUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gru *GroupRoleUpdate) defaults() {
	if _, ok := gru.mutation.UpdatedAt(); !ok {
		v := grouprole.UpdateDefaultUpdatedAt()
		gru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gru *GroupRoleUpdate) check() error {
	if v, ok := gru.mutation.Name(); ok {
		if err := grouprole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := gru.mutation.GroupsID(); gru.mutation.GroupsCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"groups\"")
	}
	if _, ok := gru.mutation.AccountsID(); gru.mutation.AccountsCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"accounts\"")
	}
	return nil
}

func (gru *GroupRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grouprole.Table,
			Columns: grouprole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: grouprole.FieldID,
			},
		},
	}
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grouprole.FieldName,
		})
	}
	if value, ok := gru.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: grouprole.FieldGroupID,
		})
	}
	if value, ok := gru.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: grouprole.FieldAccountID,
		})
	}
	if value, ok := gru.mutation.Attributes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: grouprole.FieldAttributes,
		})
	}
	if gru.mutation.AttributesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: grouprole.FieldAttributes,
		})
	}
	if value, ok := gru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grouprole.FieldUpdatedAt,
		})
	}
	if gru.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gru.mutation.AccountsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.AccountsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grouprole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupRoleUpdateOne is the builder for updating a single GroupRole entity.
type GroupRoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *GroupRoleMutation
}

// SetName sets the name field.
func (gruo *GroupRoleUpdateOne) SetName(s string) *GroupRoleUpdateOne {
	gruo.mutation.SetName(s)
	return gruo
}

// SetGroupID sets the group_id field.
func (gruo *GroupRoleUpdateOne) SetGroupID(u uuid.UUID) *GroupRoleUpdateOne {
	gruo.mutation.SetGroupID(u)
	return gruo
}

// SetAccountID sets the account_id field.
func (gruo *GroupRoleUpdateOne) SetAccountID(u uuid.UUID) *GroupRoleUpdateOne {
	gruo.mutation.SetAccountID(u)
	return gruo
}

// SetAttributes sets the attributes field.
func (gruo *GroupRoleUpdateOne) SetAttributes(m map[string]interface{}) *GroupRoleUpdateOne {
	gruo.mutation.SetAttributes(m)
	return gruo
}

// ClearAttributes clears the value of attributes.
func (gruo *GroupRoleUpdateOne) ClearAttributes() *GroupRoleUpdateOne {
	gruo.mutation.ClearAttributes()
	return gruo
}

// SetUpdatedAt sets the updated_at field.
func (gruo *GroupRoleUpdateOne) SetUpdatedAt(t time.Time) *GroupRoleUpdateOne {
	gruo.mutation.SetUpdatedAt(t)
	return gruo
}

// SetGroupsID sets the groups edge to Group by id.
func (gruo *GroupRoleUpdateOne) SetGroupsID(id uuid.UUID) *GroupRoleUpdateOne {
	gruo.mutation.SetGroupsID(id)
	return gruo
}

// SetGroups sets the groups edge to Group.
func (gruo *GroupRoleUpdateOne) SetGroups(g *Group) *GroupRoleUpdateOne {
	return gruo.SetGroupsID(g.ID)
}

// SetAccountsID sets the accounts edge to Account by id.
func (gruo *GroupRoleUpdateOne) SetAccountsID(id uuid.UUID) *GroupRoleUpdateOne {
	gruo.mutation.SetAccountsID(id)
	return gruo
}

// SetAccounts sets the accounts edge to Account.
func (gruo *GroupRoleUpdateOne) SetAccounts(a *Account) *GroupRoleUpdateOne {
	return gruo.SetAccountsID(a.ID)
}

// Mutation returns the GroupRoleMutation object of the builder.
func (gruo *GroupRoleUpdateOne) Mutation() *GroupRoleMutation {
	return gruo.mutation
}

// ClearGroups clears the "groups" edge to type Group.
func (gruo *GroupRoleUpdateOne) ClearGroups() *GroupRoleUpdateOne {
	gruo.mutation.ClearGroups()
	return gruo
}

// ClearAccounts clears the "accounts" edge to type Account.
func (gruo *GroupRoleUpdateOne) ClearAccounts() *GroupRoleUpdateOne {
	gruo.mutation.ClearAccounts()
	return gruo
}

// Save executes the query and returns the updated entity.
func (gruo *GroupRoleUpdateOne) Save(ctx context.Context) (*GroupRole, error) {
	var (
		err  error
		node *GroupRole
	)
	gruo.defaults()
	if len(gruo.hooks) == 0 {
		if err = gruo.check(); err != nil {
			return nil, err
		}
		node, err = gruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gruo.check(); err != nil {
				return nil, err
			}
			gruo.mutation = mutation
			node, err = gruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gruo.hooks) - 1; i >= 0; i-- {
			mut = gruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GroupRoleUpdateOne) SaveX(ctx context.Context) *GroupRole {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GroupRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GroupRoleUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gruo *GroupRoleUpdateOne) defaults() {
	if _, ok := gruo.mutation.UpdatedAt(); !ok {
		v := grouprole.UpdateDefaultUpdatedAt()
		gruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gruo *GroupRoleUpdateOne) check() error {
	if v, ok := gruo.mutation.Name(); ok {
		if err := grouprole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := gruo.mutation.GroupsID(); gruo.mutation.GroupsCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"groups\"")
	}
	if _, ok := gruo.mutation.AccountsID(); gruo.mutation.AccountsCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"accounts\"")
	}
	return nil
}

func (gruo *GroupRoleUpdateOne) sqlSave(ctx context.Context) (_node *GroupRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grouprole.Table,
			Columns: grouprole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: grouprole.FieldID,
			},
		},
	}
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GroupRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := gruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: grouprole.FieldName,
		})
	}
	if value, ok := gruo.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: grouprole.FieldGroupID,
		})
	}
	if value, ok := gruo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: grouprole.FieldAccountID,
		})
	}
	if value, ok := gruo.mutation.Attributes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: grouprole.FieldAttributes,
		})
	}
	if gruo.mutation.AttributesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: grouprole.FieldAttributes,
		})
	}
	if value, ok := gruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grouprole.FieldUpdatedAt,
		})
	}
	if gruo.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gruo.mutation.AccountsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.AccountsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupRole{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grouprole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
