// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adnaan/authn/models/account"
	"github.com/google/uuid"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetProvider sets the "provider" field.
func (ac *AccountCreate) SetProvider(s string) *AccountCreate {
	ac.mutation.SetProvider(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AccountCreate) SetEmail(s string) *AccountCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AccountCreate) SetPassword(s string) *AccountCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetLocked sets the "locked" field.
func (ac *AccountCreate) SetLocked(b bool) *AccountCreate {
	ac.mutation.SetLocked(b)
	return ac
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ac *AccountCreate) SetNillableLocked(b *bool) *AccountCreate {
	if b != nil {
		ac.SetLocked(*b)
	}
	return ac
}

// SetConfirmed sets the "confirmed" field.
func (ac *AccountCreate) SetConfirmed(b bool) *AccountCreate {
	ac.mutation.SetConfirmed(b)
	return ac
}

// SetNillableConfirmed sets the "confirmed" field if the given value is not nil.
func (ac *AccountCreate) SetNillableConfirmed(b *bool) *AccountCreate {
	if b != nil {
		ac.SetConfirmed(*b)
	}
	return ac
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (ac *AccountCreate) SetConfirmationSentAt(t time.Time) *AccountCreate {
	ac.mutation.SetConfirmationSentAt(t)
	return ac
}

// SetNillableConfirmationSentAt sets the "confirmation_sent_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableConfirmationSentAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetConfirmationSentAt(*t)
	}
	return ac
}

// SetConfirmationToken sets the "confirmation_token" field.
func (ac *AccountCreate) SetConfirmationToken(s string) *AccountCreate {
	ac.mutation.SetConfirmationToken(s)
	return ac
}

// SetNillableConfirmationToken sets the "confirmation_token" field if the given value is not nil.
func (ac *AccountCreate) SetNillableConfirmationToken(s *string) *AccountCreate {
	if s != nil {
		ac.SetConfirmationToken(*s)
	}
	return ac
}

// SetRecoverySentAt sets the "recovery_sent_at" field.
func (ac *AccountCreate) SetRecoverySentAt(t time.Time) *AccountCreate {
	ac.mutation.SetRecoverySentAt(t)
	return ac
}

// SetNillableRecoverySentAt sets the "recovery_sent_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableRecoverySentAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetRecoverySentAt(*t)
	}
	return ac
}

// SetRecoveryToken sets the "recovery_token" field.
func (ac *AccountCreate) SetRecoveryToken(s string) *AccountCreate {
	ac.mutation.SetRecoveryToken(s)
	return ac
}

// SetNillableRecoveryToken sets the "recovery_token" field if the given value is not nil.
func (ac *AccountCreate) SetNillableRecoveryToken(s *string) *AccountCreate {
	if s != nil {
		ac.SetRecoveryToken(*s)
	}
	return ac
}

// SetOtpSentAt sets the "otp_sent_at" field.
func (ac *AccountCreate) SetOtpSentAt(t time.Time) *AccountCreate {
	ac.mutation.SetOtpSentAt(t)
	return ac
}

// SetNillableOtpSentAt sets the "otp_sent_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableOtpSentAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetOtpSentAt(*t)
	}
	return ac
}

// SetOtp sets the "otp" field.
func (ac *AccountCreate) SetOtp(s string) *AccountCreate {
	ac.mutation.SetOtp(s)
	return ac
}

// SetNillableOtp sets the "otp" field if the given value is not nil.
func (ac *AccountCreate) SetNillableOtp(s *string) *AccountCreate {
	if s != nil {
		ac.SetOtp(*s)
	}
	return ac
}

// SetEmailChange sets the "email_change" field.
func (ac *AccountCreate) SetEmailChange(s string) *AccountCreate {
	ac.mutation.SetEmailChange(s)
	return ac
}

// SetNillableEmailChange sets the "email_change" field if the given value is not nil.
func (ac *AccountCreate) SetNillableEmailChange(s *string) *AccountCreate {
	if s != nil {
		ac.SetEmailChange(*s)
	}
	return ac
}

// SetEmailChangeSentAt sets the "email_change_sent_at" field.
func (ac *AccountCreate) SetEmailChangeSentAt(t time.Time) *AccountCreate {
	ac.mutation.SetEmailChangeSentAt(t)
	return ac
}

// SetNillableEmailChangeSentAt sets the "email_change_sent_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableEmailChangeSentAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetEmailChangeSentAt(*t)
	}
	return ac
}

// SetEmailChangeToken sets the "email_change_token" field.
func (ac *AccountCreate) SetEmailChangeToken(s string) *AccountCreate {
	ac.mutation.SetEmailChangeToken(s)
	return ac
}

// SetNillableEmailChangeToken sets the "email_change_token" field if the given value is not nil.
func (ac *AccountCreate) SetNillableEmailChangeToken(s *string) *AccountCreate {
	if s != nil {
		ac.SetEmailChangeToken(*s)
	}
	return ac
}

// SetAttributes sets the "attributes" field.
func (ac *AccountCreate) SetAttributes(m map[string]interface{}) *AccountCreate {
	ac.mutation.SetAttributes(m)
	return ac
}

// SetSensitiveAttributes sets the "sensitive_attributes" field.
func (ac *AccountCreate) SetSensitiveAttributes(m map[string]string) *AccountCreate {
	ac.mutation.SetSensitiveAttributes(m)
	return ac
}

// SetAttributeBytes sets the "attribute_bytes" field.
func (ac *AccountCreate) SetAttributeBytes(b []byte) *AccountCreate {
	ac.mutation.SetAttributeBytes(b)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetLastSigninAt sets the "last_signin_at" field.
func (ac *AccountCreate) SetLastSigninAt(t time.Time) *AccountCreate {
	ac.mutation.SetLastSigninAt(t)
	return ac
}

// SetNillableLastSigninAt sets the "last_signin_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableLastSigninAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetLastSigninAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(u uuid.UUID) *AccountCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AccountCreate) SetNillableID(u *uuid.UUID) *AccountCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.Locked(); !ok {
		v := account.DefaultLocked
		ac.mutation.SetLocked(v)
	}
	if _, ok := ac.mutation.Confirmed(); !ok {
		v := account.DefaultConfirmed
		ac.mutation.SetConfirmed(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := account.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`models: missing required field "Account.provider"`)}
	}
	if v, ok := ac.mutation.Provider(); ok {
		if err := account.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`models: validator failed for field "Account.provider": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`models: missing required field "Account.email"`)}
	}
	if v, ok := ac.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`models: validator failed for field "Account.email": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`models: missing required field "Account.password"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`models: validator failed for field "Account.password": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Locked(); !ok {
		return &ValidationError{Name: "locked", err: errors.New(`models: missing required field "Account.locked"`)}
	}
	if v, ok := ac.mutation.ConfirmationToken(); ok {
		if err := account.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf(`models: validator failed for field "Account.confirmation_token": %w`, err)}
		}
	}
	if v, ok := ac.mutation.RecoveryToken(); ok {
		if err := account.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf(`models: validator failed for field "Account.recovery_token": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Otp(); ok {
		if err := account.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf(`models: validator failed for field "Account.otp": %w`, err)}
		}
	}
	if v, ok := ac.mutation.EmailChange(); ok {
		if err := account.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf(`models: validator failed for field "Account.email_change": %w`, err)}
		}
	}
	if v, ok := ac.mutation.EmailChangeToken(); ok {
		if err := account.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf(`models: validator failed for field "Account.email_change_token": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`models: missing required field "Account.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`models: missing required field "Account.updated_at"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Provider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProvider,
		})
		_node.Provider = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := ac.mutation.Locked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldLocked,
		})
		_node.Locked = value
	}
	if value, ok := ac.mutation.Confirmed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldConfirmed,
		})
		_node.Confirmed = value
	}
	if value, ok := ac.mutation.ConfirmationSentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldConfirmationSentAt,
		})
		_node.ConfirmationSentAt = &value
	}
	if value, ok := ac.mutation.ConfirmationToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldConfirmationToken,
		})
		_node.ConfirmationToken = &value
	}
	if value, ok := ac.mutation.RecoverySentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldRecoverySentAt,
		})
		_node.RecoverySentAt = &value
	}
	if value, ok := ac.mutation.RecoveryToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldRecoveryToken,
		})
		_node.RecoveryToken = &value
	}
	if value, ok := ac.mutation.OtpSentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOtpSentAt,
		})
		_node.OtpSentAt = &value
	}
	if value, ok := ac.mutation.Otp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOtp,
		})
		_node.Otp = &value
	}
	if value, ok := ac.mutation.EmailChange(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChange,
		})
		_node.EmailChange = value
	}
	if value, ok := ac.mutation.EmailChangeSentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldEmailChangeSentAt,
		})
		_node.EmailChangeSentAt = &value
	}
	if value, ok := ac.mutation.EmailChangeToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChangeToken,
		})
		_node.EmailChangeToken = &value
	}
	if value, ok := ac.mutation.Attributes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldAttributes,
		})
		_node.Attributes = value
	}
	if value, ok := ac.mutation.SensitiveAttributes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldSensitiveAttributes,
		})
		_node.SensitiveAttributes = value
	}
	if value, ok := ac.mutation.AttributeBytes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldAttributeBytes,
		})
		_node.AttributeBytes = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.LastSigninAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldLastSigninAt,
		})
		_node.LastSigninAt = &value
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
