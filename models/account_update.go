// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adnaan/authn/models/account"
	"github.com/adnaan/authn/models/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where adds a new predicate for the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetProvider sets the "provider" field.
func (au *AccountUpdate) SetProvider(s string) *AccountUpdate {
	au.mutation.SetProvider(s)
	return au
}

// SetEmail sets the "email" field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetLocked sets the "locked" field.
func (au *AccountUpdate) SetLocked(b bool) *AccountUpdate {
	au.mutation.SetLocked(b)
	return au
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (au *AccountUpdate) SetNillableLocked(b *bool) *AccountUpdate {
	if b != nil {
		au.SetLocked(*b)
	}
	return au
}

// SetConfirmed sets the "confirmed" field.
func (au *AccountUpdate) SetConfirmed(b bool) *AccountUpdate {
	au.mutation.SetConfirmed(b)
	return au
}

// SetNillableConfirmed sets the "confirmed" field if the given value is not nil.
func (au *AccountUpdate) SetNillableConfirmed(b *bool) *AccountUpdate {
	if b != nil {
		au.SetConfirmed(*b)
	}
	return au
}

// ClearConfirmed clears the value of the "confirmed" field.
func (au *AccountUpdate) ClearConfirmed() *AccountUpdate {
	au.mutation.ClearConfirmed()
	return au
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (au *AccountUpdate) SetConfirmationSentAt(t time.Time) *AccountUpdate {
	au.mutation.SetConfirmationSentAt(t)
	return au
}

// SetNillableConfirmationSentAt sets the "confirmation_sent_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableConfirmationSentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetConfirmationSentAt(*t)
	}
	return au
}

// ClearConfirmationSentAt clears the value of the "confirmation_sent_at" field.
func (au *AccountUpdate) ClearConfirmationSentAt() *AccountUpdate {
	au.mutation.ClearConfirmationSentAt()
	return au
}

// SetConfirmationToken sets the "confirmation_token" field.
func (au *AccountUpdate) SetConfirmationToken(s string) *AccountUpdate {
	au.mutation.SetConfirmationToken(s)
	return au
}

// SetNillableConfirmationToken sets the "confirmation_token" field if the given value is not nil.
func (au *AccountUpdate) SetNillableConfirmationToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetConfirmationToken(*s)
	}
	return au
}

// ClearConfirmationToken clears the value of the "confirmation_token" field.
func (au *AccountUpdate) ClearConfirmationToken() *AccountUpdate {
	au.mutation.ClearConfirmationToken()
	return au
}

// SetRecoverySentAt sets the "recovery_sent_at" field.
func (au *AccountUpdate) SetRecoverySentAt(t time.Time) *AccountUpdate {
	au.mutation.SetRecoverySentAt(t)
	return au
}

// SetNillableRecoverySentAt sets the "recovery_sent_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableRecoverySentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetRecoverySentAt(*t)
	}
	return au
}

// ClearRecoverySentAt clears the value of the "recovery_sent_at" field.
func (au *AccountUpdate) ClearRecoverySentAt() *AccountUpdate {
	au.mutation.ClearRecoverySentAt()
	return au
}

// SetRecoveryToken sets the "recovery_token" field.
func (au *AccountUpdate) SetRecoveryToken(s string) *AccountUpdate {
	au.mutation.SetRecoveryToken(s)
	return au
}

// SetNillableRecoveryToken sets the "recovery_token" field if the given value is not nil.
func (au *AccountUpdate) SetNillableRecoveryToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetRecoveryToken(*s)
	}
	return au
}

// ClearRecoveryToken clears the value of the "recovery_token" field.
func (au *AccountUpdate) ClearRecoveryToken() *AccountUpdate {
	au.mutation.ClearRecoveryToken()
	return au
}

// SetOtpSentAt sets the "otp_sent_at" field.
func (au *AccountUpdate) SetOtpSentAt(t time.Time) *AccountUpdate {
	au.mutation.SetOtpSentAt(t)
	return au
}

// SetNillableOtpSentAt sets the "otp_sent_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableOtpSentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetOtpSentAt(*t)
	}
	return au
}

// ClearOtpSentAt clears the value of the "otp_sent_at" field.
func (au *AccountUpdate) ClearOtpSentAt() *AccountUpdate {
	au.mutation.ClearOtpSentAt()
	return au
}

// SetOtp sets the "otp" field.
func (au *AccountUpdate) SetOtp(s string) *AccountUpdate {
	au.mutation.SetOtp(s)
	return au
}

// SetNillableOtp sets the "otp" field if the given value is not nil.
func (au *AccountUpdate) SetNillableOtp(s *string) *AccountUpdate {
	if s != nil {
		au.SetOtp(*s)
	}
	return au
}

// ClearOtp clears the value of the "otp" field.
func (au *AccountUpdate) ClearOtp() *AccountUpdate {
	au.mutation.ClearOtp()
	return au
}

// SetEmailChange sets the "email_change" field.
func (au *AccountUpdate) SetEmailChange(s string) *AccountUpdate {
	au.mutation.SetEmailChange(s)
	return au
}

// SetNillableEmailChange sets the "email_change" field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmailChange(s *string) *AccountUpdate {
	if s != nil {
		au.SetEmailChange(*s)
	}
	return au
}

// ClearEmailChange clears the value of the "email_change" field.
func (au *AccountUpdate) ClearEmailChange() *AccountUpdate {
	au.mutation.ClearEmailChange()
	return au
}

// SetEmailChangeSentAt sets the "email_change_sent_at" field.
func (au *AccountUpdate) SetEmailChangeSentAt(t time.Time) *AccountUpdate {
	au.mutation.SetEmailChangeSentAt(t)
	return au
}

// SetNillableEmailChangeSentAt sets the "email_change_sent_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmailChangeSentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetEmailChangeSentAt(*t)
	}
	return au
}

// ClearEmailChangeSentAt clears the value of the "email_change_sent_at" field.
func (au *AccountUpdate) ClearEmailChangeSentAt() *AccountUpdate {
	au.mutation.ClearEmailChangeSentAt()
	return au
}

// SetEmailChangeToken sets the "email_change_token" field.
func (au *AccountUpdate) SetEmailChangeToken(s string) *AccountUpdate {
	au.mutation.SetEmailChangeToken(s)
	return au
}

// SetNillableEmailChangeToken sets the "email_change_token" field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmailChangeToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetEmailChangeToken(*s)
	}
	return au
}

// ClearEmailChangeToken clears the value of the "email_change_token" field.
func (au *AccountUpdate) ClearEmailChangeToken() *AccountUpdate {
	au.mutation.ClearEmailChangeToken()
	return au
}

// SetAttributes sets the "attributes" field.
func (au *AccountUpdate) SetAttributes(m map[string]interface{}) *AccountUpdate {
	au.mutation.SetAttributes(m)
	return au
}

// ClearAttributes clears the value of the "attributes" field.
func (au *AccountUpdate) ClearAttributes() *AccountUpdate {
	au.mutation.ClearAttributes()
	return au
}

// SetSensitiveAttributes sets the "sensitive_attributes" field.
func (au *AccountUpdate) SetSensitiveAttributes(m map[string]string) *AccountUpdate {
	au.mutation.SetSensitiveAttributes(m)
	return au
}

// ClearSensitiveAttributes clears the value of the "sensitive_attributes" field.
func (au *AccountUpdate) ClearSensitiveAttributes() *AccountUpdate {
	au.mutation.ClearSensitiveAttributes()
	return au
}

// SetAttributeBytes sets the "attribute_bytes" field.
func (au *AccountUpdate) SetAttributeBytes(b []byte) *AccountUpdate {
	au.mutation.SetAttributeBytes(b)
	return au
}

// ClearAttributeBytes clears the value of the "attribute_bytes" field.
func (au *AccountUpdate) ClearAttributeBytes() *AccountUpdate {
	au.mutation.ClearAttributeBytes()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetLastSigninAt sets the "last_signin_at" field.
func (au *AccountUpdate) SetLastSigninAt(t time.Time) *AccountUpdate {
	au.mutation.SetLastSigninAt(t)
	return au
}

// SetNillableLastSigninAt sets the "last_signin_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableLastSigninAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetLastSigninAt(*t)
	}
	return au
}

// ClearLastSigninAt clears the value of the "last_signin_at" field.
func (au *AccountUpdate) ClearLastSigninAt() *AccountUpdate {
	au.mutation.ClearLastSigninAt()
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Provider(); ok {
		if err := account.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if v, ok := au.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := au.mutation.ConfirmationToken(); ok {
		if err := account.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := au.mutation.RecoveryToken(); ok {
		if err := account.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := au.mutation.Otp(); ok {
		if err := account.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := au.mutation.EmailChange(); ok {
		if err := account.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := au.mutation.EmailChangeToken(); ok {
		if err := account.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProvider,
		})
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := au.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldLocked,
		})
	}
	if value, ok := au.mutation.Confirmed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldConfirmed,
		})
	}
	if au.mutation.ConfirmedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldConfirmed,
		})
	}
	if value, ok := au.mutation.ConfirmationSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if au.mutation.ConfirmationSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if value, ok := au.mutation.ConfirmationToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldConfirmationToken,
		})
	}
	if au.mutation.ConfirmationTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldConfirmationToken,
		})
	}
	if value, ok := au.mutation.RecoverySentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldRecoverySentAt,
		})
	}
	if au.mutation.RecoverySentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldRecoverySentAt,
		})
	}
	if value, ok := au.mutation.RecoveryToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldRecoveryToken,
		})
	}
	if au.mutation.RecoveryTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldRecoveryToken,
		})
	}
	if value, ok := au.mutation.OtpSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOtpSentAt,
		})
	}
	if au.mutation.OtpSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldOtpSentAt,
		})
	}
	if value, ok := au.mutation.Otp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOtp,
		})
	}
	if au.mutation.OtpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldOtp,
		})
	}
	if value, ok := au.mutation.EmailChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChange,
		})
	}
	if au.mutation.EmailChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChange,
		})
	}
	if value, ok := au.mutation.EmailChangeSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if au.mutation.EmailChangeSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if value, ok := au.mutation.EmailChangeToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChangeToken,
		})
	}
	if au.mutation.EmailChangeTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChangeToken,
		})
	}
	if value, ok := au.mutation.Attributes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldAttributes,
		})
	}
	if au.mutation.AttributesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldAttributes,
		})
	}
	if value, ok := au.mutation.SensitiveAttributes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldSensitiveAttributes,
		})
	}
	if au.mutation.SensitiveAttributesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldSensitiveAttributes,
		})
	}
	if value, ok := au.mutation.AttributeBytes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldAttributeBytes,
		})
	}
	if au.mutation.AttributeBytesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: account.FieldAttributeBytes,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.LastSigninAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldLastSigninAt,
		})
	}
	if au.mutation.LastSigninAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldLastSigninAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// SetProvider sets the "provider" field.
func (auo *AccountUpdateOne) SetProvider(s string) *AccountUpdateOne {
	auo.mutation.SetProvider(s)
	return auo
}

// SetEmail sets the "email" field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetLocked sets the "locked" field.
func (auo *AccountUpdateOne) SetLocked(b bool) *AccountUpdateOne {
	auo.mutation.SetLocked(b)
	return auo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableLocked(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetLocked(*b)
	}
	return auo
}

// SetConfirmed sets the "confirmed" field.
func (auo *AccountUpdateOne) SetConfirmed(b bool) *AccountUpdateOne {
	auo.mutation.SetConfirmed(b)
	return auo
}

// SetNillableConfirmed sets the "confirmed" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableConfirmed(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetConfirmed(*b)
	}
	return auo
}

// ClearConfirmed clears the value of the "confirmed" field.
func (auo *AccountUpdateOne) ClearConfirmed() *AccountUpdateOne {
	auo.mutation.ClearConfirmed()
	return auo
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (auo *AccountUpdateOne) SetConfirmationSentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetConfirmationSentAt(t)
	return auo
}

// SetNillableConfirmationSentAt sets the "confirmation_sent_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableConfirmationSentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetConfirmationSentAt(*t)
	}
	return auo
}

// ClearConfirmationSentAt clears the value of the "confirmation_sent_at" field.
func (auo *AccountUpdateOne) ClearConfirmationSentAt() *AccountUpdateOne {
	auo.mutation.ClearConfirmationSentAt()
	return auo
}

// SetConfirmationToken sets the "confirmation_token" field.
func (auo *AccountUpdateOne) SetConfirmationToken(s string) *AccountUpdateOne {
	auo.mutation.SetConfirmationToken(s)
	return auo
}

// SetNillableConfirmationToken sets the "confirmation_token" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableConfirmationToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetConfirmationToken(*s)
	}
	return auo
}

// ClearConfirmationToken clears the value of the "confirmation_token" field.
func (auo *AccountUpdateOne) ClearConfirmationToken() *AccountUpdateOne {
	auo.mutation.ClearConfirmationToken()
	return auo
}

// SetRecoverySentAt sets the "recovery_sent_at" field.
func (auo *AccountUpdateOne) SetRecoverySentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetRecoverySentAt(t)
	return auo
}

// SetNillableRecoverySentAt sets the "recovery_sent_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableRecoverySentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetRecoverySentAt(*t)
	}
	return auo
}

// ClearRecoverySentAt clears the value of the "recovery_sent_at" field.
func (auo *AccountUpdateOne) ClearRecoverySentAt() *AccountUpdateOne {
	auo.mutation.ClearRecoverySentAt()
	return auo
}

// SetRecoveryToken sets the "recovery_token" field.
func (auo *AccountUpdateOne) SetRecoveryToken(s string) *AccountUpdateOne {
	auo.mutation.SetRecoveryToken(s)
	return auo
}

// SetNillableRecoveryToken sets the "recovery_token" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableRecoveryToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetRecoveryToken(*s)
	}
	return auo
}

// ClearRecoveryToken clears the value of the "recovery_token" field.
func (auo *AccountUpdateOne) ClearRecoveryToken() *AccountUpdateOne {
	auo.mutation.ClearRecoveryToken()
	return auo
}

// SetOtpSentAt sets the "otp_sent_at" field.
func (auo *AccountUpdateOne) SetOtpSentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetOtpSentAt(t)
	return auo
}

// SetNillableOtpSentAt sets the "otp_sent_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableOtpSentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetOtpSentAt(*t)
	}
	return auo
}

// ClearOtpSentAt clears the value of the "otp_sent_at" field.
func (auo *AccountUpdateOne) ClearOtpSentAt() *AccountUpdateOne {
	auo.mutation.ClearOtpSentAt()
	return auo
}

// SetOtp sets the "otp" field.
func (auo *AccountUpdateOne) SetOtp(s string) *AccountUpdateOne {
	auo.mutation.SetOtp(s)
	return auo
}

// SetNillableOtp sets the "otp" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableOtp(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetOtp(*s)
	}
	return auo
}

// ClearOtp clears the value of the "otp" field.
func (auo *AccountUpdateOne) ClearOtp() *AccountUpdateOne {
	auo.mutation.ClearOtp()
	return auo
}

// SetEmailChange sets the "email_change" field.
func (auo *AccountUpdateOne) SetEmailChange(s string) *AccountUpdateOne {
	auo.mutation.SetEmailChange(s)
	return auo
}

// SetNillableEmailChange sets the "email_change" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmailChange(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetEmailChange(*s)
	}
	return auo
}

// ClearEmailChange clears the value of the "email_change" field.
func (auo *AccountUpdateOne) ClearEmailChange() *AccountUpdateOne {
	auo.mutation.ClearEmailChange()
	return auo
}

// SetEmailChangeSentAt sets the "email_change_sent_at" field.
func (auo *AccountUpdateOne) SetEmailChangeSentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetEmailChangeSentAt(t)
	return auo
}

// SetNillableEmailChangeSentAt sets the "email_change_sent_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmailChangeSentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetEmailChangeSentAt(*t)
	}
	return auo
}

// ClearEmailChangeSentAt clears the value of the "email_change_sent_at" field.
func (auo *AccountUpdateOne) ClearEmailChangeSentAt() *AccountUpdateOne {
	auo.mutation.ClearEmailChangeSentAt()
	return auo
}

// SetEmailChangeToken sets the "email_change_token" field.
func (auo *AccountUpdateOne) SetEmailChangeToken(s string) *AccountUpdateOne {
	auo.mutation.SetEmailChangeToken(s)
	return auo
}

// SetNillableEmailChangeToken sets the "email_change_token" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmailChangeToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetEmailChangeToken(*s)
	}
	return auo
}

// ClearEmailChangeToken clears the value of the "email_change_token" field.
func (auo *AccountUpdateOne) ClearEmailChangeToken() *AccountUpdateOne {
	auo.mutation.ClearEmailChangeToken()
	return auo
}

// SetAttributes sets the "attributes" field.
func (auo *AccountUpdateOne) SetAttributes(m map[string]interface{}) *AccountUpdateOne {
	auo.mutation.SetAttributes(m)
	return auo
}

// ClearAttributes clears the value of the "attributes" field.
func (auo *AccountUpdateOne) ClearAttributes() *AccountUpdateOne {
	auo.mutation.ClearAttributes()
	return auo
}

// SetSensitiveAttributes sets the "sensitive_attributes" field.
func (auo *AccountUpdateOne) SetSensitiveAttributes(m map[string]string) *AccountUpdateOne {
	auo.mutation.SetSensitiveAttributes(m)
	return auo
}

// ClearSensitiveAttributes clears the value of the "sensitive_attributes" field.
func (auo *AccountUpdateOne) ClearSensitiveAttributes() *AccountUpdateOne {
	auo.mutation.ClearSensitiveAttributes()
	return auo
}

// SetAttributeBytes sets the "attribute_bytes" field.
func (auo *AccountUpdateOne) SetAttributeBytes(b []byte) *AccountUpdateOne {
	auo.mutation.SetAttributeBytes(b)
	return auo
}

// ClearAttributeBytes clears the value of the "attribute_bytes" field.
func (auo *AccountUpdateOne) ClearAttributeBytes() *AccountUpdateOne {
	auo.mutation.ClearAttributeBytes()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetLastSigninAt sets the "last_signin_at" field.
func (auo *AccountUpdateOne) SetLastSigninAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetLastSigninAt(t)
	return auo
}

// SetNillableLastSigninAt sets the "last_signin_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableLastSigninAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetLastSigninAt(*t)
	}
	return auo
}

// ClearLastSigninAt clears the value of the "last_signin_at" field.
func (auo *AccountUpdateOne) ClearLastSigninAt() *AccountUpdateOne {
	auo.mutation.ClearLastSigninAt()
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Provider(); ok {
		if err := account.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := auo.mutation.ConfirmationToken(); ok {
		if err := account.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := auo.mutation.RecoveryToken(); ok {
		if err := account.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Otp(); ok {
		if err := account.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := auo.mutation.EmailChange(); ok {
		if err := account.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := auo.mutation.EmailChangeToken(); ok {
		if err := account.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Account.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProvider,
		})
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := auo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldLocked,
		})
	}
	if value, ok := auo.mutation.Confirmed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldConfirmed,
		})
	}
	if auo.mutation.ConfirmedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldConfirmed,
		})
	}
	if value, ok := auo.mutation.ConfirmationSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if auo.mutation.ConfirmationSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if value, ok := auo.mutation.ConfirmationToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldConfirmationToken,
		})
	}
	if auo.mutation.ConfirmationTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldConfirmationToken,
		})
	}
	if value, ok := auo.mutation.RecoverySentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldRecoverySentAt,
		})
	}
	if auo.mutation.RecoverySentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldRecoverySentAt,
		})
	}
	if value, ok := auo.mutation.RecoveryToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldRecoveryToken,
		})
	}
	if auo.mutation.RecoveryTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldRecoveryToken,
		})
	}
	if value, ok := auo.mutation.OtpSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOtpSentAt,
		})
	}
	if auo.mutation.OtpSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldOtpSentAt,
		})
	}
	if value, ok := auo.mutation.Otp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOtp,
		})
	}
	if auo.mutation.OtpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldOtp,
		})
	}
	if value, ok := auo.mutation.EmailChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChange,
		})
	}
	if auo.mutation.EmailChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChange,
		})
	}
	if value, ok := auo.mutation.EmailChangeSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if auo.mutation.EmailChangeSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if value, ok := auo.mutation.EmailChangeToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChangeToken,
		})
	}
	if auo.mutation.EmailChangeTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChangeToken,
		})
	}
	if value, ok := auo.mutation.Attributes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldAttributes,
		})
	}
	if auo.mutation.AttributesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldAttributes,
		})
	}
	if value, ok := auo.mutation.SensitiveAttributes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldSensitiveAttributes,
		})
	}
	if auo.mutation.SensitiveAttributesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldSensitiveAttributes,
		})
	}
	if value, ok := auo.mutation.AttributeBytes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: account.FieldAttributeBytes,
		})
	}
	if auo.mutation.AttributeBytesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: account.FieldAttributeBytes,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.LastSigninAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldLastSigninAt,
		})
	}
	if auo.mutation.LastSigninAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldLastSigninAt,
		})
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
