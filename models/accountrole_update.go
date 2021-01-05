// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AccountRoleUpdate is the builder for updating AccountRole entities.
type AccountRoleUpdate struct {
	config
	hooks    []Hook
	mutation *AccountRoleMutation
}

// Where adds a new predicate for the builder.
func (aru *AccountRoleUpdate) Where(ps ...predicate.AccountRole) *AccountRoleUpdate {
	aru.mutation.predicates = append(aru.mutation.predicates, ps...)
	return aru
}

// SetName sets the name field.
func (aru *AccountRoleUpdate) SetName(s string) *AccountRoleUpdate {
	aru.mutation.SetName(s)
	return aru
}

// SetAccountID sets the account_id field.
func (aru *AccountRoleUpdate) SetAccountID(u uuid.UUID) *AccountRoleUpdate {
	aru.mutation.SetAccountID(u)
	return aru
}

// SetMetadata sets the metadata field.
func (aru *AccountRoleUpdate) SetMetadata(m map[string]interface{}) *AccountRoleUpdate {
	aru.mutation.SetMetadata(m)
	return aru
}

// ClearMetadata clears the value of metadata.
func (aru *AccountRoleUpdate) ClearMetadata() *AccountRoleUpdate {
	aru.mutation.ClearMetadata()
	return aru
}

// SetUpdatedAt sets the updated_at field.
func (aru *AccountRoleUpdate) SetUpdatedAt(t time.Time) *AccountRoleUpdate {
	aru.mutation.SetUpdatedAt(t)
	return aru
}

// SetAccountsID sets the accounts edge to Account by id.
func (aru *AccountRoleUpdate) SetAccountsID(id uuid.UUID) *AccountRoleUpdate {
	aru.mutation.SetAccountsID(id)
	return aru
}

// SetAccounts sets the accounts edge to Account.
func (aru *AccountRoleUpdate) SetAccounts(a *Account) *AccountRoleUpdate {
	return aru.SetAccountsID(a.ID)
}

// Mutation returns the AccountRoleMutation object of the builder.
func (aru *AccountRoleUpdate) Mutation() *AccountRoleMutation {
	return aru.mutation
}

// ClearAccounts clears the "accounts" edge to type Account.
func (aru *AccountRoleUpdate) ClearAccounts() *AccountRoleUpdate {
	aru.mutation.ClearAccounts()
	return aru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AccountRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aru.defaults()
	if len(aru.hooks) == 0 {
		if err = aru.check(); err != nil {
			return 0, err
		}
		affected, err = aru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aru.check(); err != nil {
				return 0, err
			}
			aru.mutation = mutation
			affected, err = aru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aru.hooks) - 1; i >= 0; i-- {
			mut = aru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AccountRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AccountRoleUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AccountRoleUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aru *AccountRoleUpdate) defaults() {
	if _, ok := aru.mutation.UpdatedAt(); !ok {
		v := accountrole.UpdateDefaultUpdatedAt()
		aru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aru *AccountRoleUpdate) check() error {
	if v, ok := aru.mutation.Name(); ok {
		if err := accountrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := aru.mutation.AccountsID(); aru.mutation.AccountsCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"accounts\"")
	}
	return nil
}

func (aru *AccountRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountrole.Table,
			Columns: accountrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountrole.FieldID,
			},
		},
	}
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountrole.FieldName,
		})
	}
	if value, ok := aru.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: accountrole.FieldAccountID,
		})
	}
	if value, ok := aru.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: accountrole.FieldMetadata,
		})
	}
	if aru.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: accountrole.FieldMetadata,
		})
	}
	if value, ok := aru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accountrole.FieldUpdatedAt,
		})
	}
	if aru.mutation.AccountsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.AccountsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AccountRoleUpdateOne is the builder for updating a single AccountRole entity.
type AccountRoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *AccountRoleMutation
}

// SetName sets the name field.
func (aruo *AccountRoleUpdateOne) SetName(s string) *AccountRoleUpdateOne {
	aruo.mutation.SetName(s)
	return aruo
}

// SetAccountID sets the account_id field.
func (aruo *AccountRoleUpdateOne) SetAccountID(u uuid.UUID) *AccountRoleUpdateOne {
	aruo.mutation.SetAccountID(u)
	return aruo
}

// SetMetadata sets the metadata field.
func (aruo *AccountRoleUpdateOne) SetMetadata(m map[string]interface{}) *AccountRoleUpdateOne {
	aruo.mutation.SetMetadata(m)
	return aruo
}

// ClearMetadata clears the value of metadata.
func (aruo *AccountRoleUpdateOne) ClearMetadata() *AccountRoleUpdateOne {
	aruo.mutation.ClearMetadata()
	return aruo
}

// SetUpdatedAt sets the updated_at field.
func (aruo *AccountRoleUpdateOne) SetUpdatedAt(t time.Time) *AccountRoleUpdateOne {
	aruo.mutation.SetUpdatedAt(t)
	return aruo
}

// SetAccountsID sets the accounts edge to Account by id.
func (aruo *AccountRoleUpdateOne) SetAccountsID(id uuid.UUID) *AccountRoleUpdateOne {
	aruo.mutation.SetAccountsID(id)
	return aruo
}

// SetAccounts sets the accounts edge to Account.
func (aruo *AccountRoleUpdateOne) SetAccounts(a *Account) *AccountRoleUpdateOne {
	return aruo.SetAccountsID(a.ID)
}

// Mutation returns the AccountRoleMutation object of the builder.
func (aruo *AccountRoleUpdateOne) Mutation() *AccountRoleMutation {
	return aruo.mutation
}

// ClearAccounts clears the "accounts" edge to type Account.
func (aruo *AccountRoleUpdateOne) ClearAccounts() *AccountRoleUpdateOne {
	aruo.mutation.ClearAccounts()
	return aruo
}

// Save executes the query and returns the updated entity.
func (aruo *AccountRoleUpdateOne) Save(ctx context.Context) (*AccountRole, error) {
	var (
		err  error
		node *AccountRole
	)
	aruo.defaults()
	if len(aruo.hooks) == 0 {
		if err = aruo.check(); err != nil {
			return nil, err
		}
		node, err = aruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aruo.check(); err != nil {
				return nil, err
			}
			aruo.mutation = mutation
			node, err = aruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aruo.hooks) - 1; i >= 0; i-- {
			mut = aruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AccountRoleUpdateOne) SaveX(ctx context.Context) *AccountRole {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AccountRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AccountRoleUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruo *AccountRoleUpdateOne) defaults() {
	if _, ok := aruo.mutation.UpdatedAt(); !ok {
		v := accountrole.UpdateDefaultUpdatedAt()
		aruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aruo *AccountRoleUpdateOne) check() error {
	if v, ok := aruo.mutation.Name(); ok {
		if err := accountrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("models: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := aruo.mutation.AccountsID(); aruo.mutation.AccountsCleared() && !ok {
		return errors.New("models: clearing a required unique edge \"accounts\"")
	}
	return nil
}

func (aruo *AccountRoleUpdateOne) sqlSave(ctx context.Context) (_node *AccountRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountrole.Table,
			Columns: accountrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountrole.FieldID,
			},
		},
	}
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AccountRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := aruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accountrole.FieldName,
		})
	}
	if value, ok := aruo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: accountrole.FieldAccountID,
		})
	}
	if value, ok := aruo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: accountrole.FieldMetadata,
		})
	}
	if aruo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: accountrole.FieldMetadata,
		})
	}
	if value, ok := aruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accountrole.FieldUpdatedAt,
		})
	}
	if aruo.mutation.AccountsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.AccountsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AccountRole{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
